package tkt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
)

func TransactionalLoggable(path string, db *gorm.DB, f func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, InterceptFatal(InterceptCORS(InterceptLoggable(db, InterceptTransactional(db, Auditable(f))))))
}

func AuthenticatedEndpoint(path string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, InterceptFatal(InterceptCORS(jwtAuth.InterceptAuth(f))))
}

func InterceptFatal(delegate func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer catchFatal(w, r)
		delegate(w, r)
	}
}

func InterceptCORS(delegate func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			header := r.Header.Get("Access-Control-Request-Headers")
			if len(header) > 0 {
				w.Header().Add("Access-Control-Allow-Headers", header)
			}
		} else {
			delegate(w, r)
		}
	}
}

type RequestData struct {
	Url  *url.URL `json:"url"`
	Data []byte   `json:"data"`
}

type ErrorData struct {
	Error   interface{}       `json:"error"`
	Stack   []string          `json:"stack"`
	Context map[string]string `json:"context"`
}

type RequestError struct {
	RequestData *json.RawMessage
	Data        *json.RawMessage
}

func InterceptLoggable(db *gorm.DB, f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		CheckErr(err)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		errorContext := NewErrorContext()
		ctx := context.WithValue(r.Context(), "errorContext", errorContext)
		defer func() {
			if e := recover(); e != nil {
				/*
					ExecuteTransactional(db, func(tx *gorm.DB, args ...interface{}) interface{} {
						rd := RequestData{Url: r.URL, Data: data}
						stack := strings.Split(string(debug.Stack()), "\n")
						ed := ErrorData{Error: e, Stack: stack, Context: errorContext.Values}
						rb := bytes.Buffer{}
						JsonEncode(rd, &rb)
						eb := bytes.Buffer{}
						JsonEncode(ed, &eb)
						r := json.RawMessage(rb.Bytes())
						d := json.RawMessage(eb.Bytes())
						_ := RequestError{RequestData: &r, Data: &d}
						//pep.NewApi(tx).CreateRequestError(re)
						return nil
					})
				*/
				panic(e)
			}
		}()
		f(w, r.WithContext(ctx))
	}
}

func InterceptTransactional(db *gorm.DB, delegate func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer CloseDB(db)
		tx := db.Begin()
		defer RollbackOnPanic(tx)
		ctx := context.WithValue(r.Context(), "tx", tx)
		delegate(tx, w, r.WithContext(ctx))
		tx.Commit()
	}
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	CheckErr(err)
	sqlDB.Close()
}

func RollbackOnPanic(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	}
}

func Auditable(delegate func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) func(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
	return func(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
		/*
			tokenEntry := r.Context().Value("tokenEntry")
			if tokenEntry != nil {
				txCtx.ExecSql(fmt.Sprintf("set local nexus.user_name to '%d';", tokenEntry.(*auth.TokenEntry).UserId))
			}
			txCtx.ExecSql(fmt.Sprintf("set local nexus.context to '%s';", r.URL.Path))
		*/
		delegate(tx, w, r)
	}
}

func ExecuteTransactional(db *gorm.DB, callback func(tx *gorm.DB, args ...interface{}) interface{}, args ...interface{}) interface{} {
	defer CloseDB(db)
	tx := db.Begin()
	defer RollbackOnPanic(tx)
	r := callback(tx, args...)
	tx.Commit()
	return r
}

type ErrorContext struct {
	Values map[string]string
}

func NewErrorContext() *ErrorContext {
	return &ErrorContext{
		Values: make(map[string]string),
	}
}

func JsonResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	JsonEncode(i, w)
}

func JsonErrorResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusInternalServerError)
	JsonEncode(i, w)
}

func JsonEncode(i interface{}, w io.Writer) {
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		panic(err)
	}
}

func catchFatal(writer http.ResponseWriter, r *http.Request) {
	if e := recover(); e != nil {
		Logger("error").Printf("Error executing %s", r.URL.String())
		errorType := reflect.TypeOf(e)
		if errorType.Kind() == reflect.Struct {
			jsonBytes := Marshal(e)
			ProcessPanic(string(jsonBytes))
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonBytes)
		} else {
			ProcessPanic(e)
			http.Error(writer, fmt.Sprint(e), http.StatusInternalServerError)
		}
	}
}

func Marshal(object interface{}) []byte {
	bytes, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return bytes
}

func ProcessPanic(intf interface{}) {
	Logger("error").Println(intf)
	stackTrace := string(debug.Stack())
	Logger("error").Println(stackTrace)
}

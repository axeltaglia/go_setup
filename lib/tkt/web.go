package tkt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

func TransactionalLoggable(path string, dbConfig *string, f func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, InterceptFatal(InterceptCORS(InterceptLoggable(dbConfig, InterceptTransactional(dbConfig, Auditable(f))))))
}

func AuthenticatedTransactional(path string, dbConfig *string, f func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, InterceptFatal(InterceptCORS(InterceptAuth(InterceptLoggable(dbConfig, InterceptTransactional(dbConfig, Auditable(f)))))))
}

func InterceptAuth(delegate func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenEntry, err := jwtAuth.VerifyToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			panic("unauthorized")
		} else {
			ctx := context.WithValue(r.Context(), "tokenEntry", tokenEntry)
			delegate(w, r.WithContext(ctx))
		}
	}
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

func InterceptLoggable(dbConfig *string, f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		CheckErr(err)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		errorContext := NewErrorContext()
		ctx := context.WithValue(r.Context(), "errorContext", errorContext)
		defer func() {
			if e := recover(); e != nil {

				ExecuteTransactional(dbConfig, func(tx *gorm.DB, args ...interface{}) interface{} {
					rd := RequestData{Url: r.URL, Data: data}
					stack := strings.Split(string(debug.Stack()), "\n")
					ed := ErrorData{Error: e, Stack: stack, Context: errorContext.Values}
					rb := bytes.Buffer{}
					JsonEncode(rd, &rb)
					eb := bytes.Buffer{}
					JsonEncode(ed, &eb)
					//r := json.RawMessage(rb.Bytes())
					//d := json.RawMessage(eb.Bytes())
					//_ := RequestError{RequestData: &r, Data: &d}
					//pep.NewApi(tx).CreateRequestError(re)
					return nil
				})

				panic(e)
			}
		}()
		f(w, r.WithContext(ctx))
	}
}

func InterceptTransactional(dbConfig *string, delegate func(tx *gorm.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := gorm.Open(postgres.Open(*dbConfig), &gorm.Config{SkipDefaultTransaction: true})
		CheckErr(err)
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
	_ = sqlDB.Close()
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

func ExecuteTransactional(dbConfig *string, callback func(tx *gorm.DB, args ...interface{}) interface{}, args ...interface{}) interface{} {
	db, err := gorm.Open(postgres.Open(*dbConfig), &gorm.Config{SkipDefaultTransaction: true})
	CheckErr(err)
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

func ParseParamOrBody(r *http.Request, o interface{}) error {
	s := r.URL.Query().Get("body")
	if len(s) > 0 {
		return json.NewDecoder(strings.NewReader(s)).Decode(o)
	} else {
		return json.NewDecoder(r.Body).Decode(o)
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

func PString(s string) *string {
	return &s
}

func PStringf(s string, values ...interface{}) *string {
	return PString(fmt.Sprintf(s, values...))
}

func PInt64(i int64) *int64 {
	return &i
}

func PInt(i int) *int {
	return &i
}

func PFloat32(f float32) *float32 {
	return &f
}

func PFloat64(f float64) *float64 {
	return &f
}

func PTime(t time.Time) *time.Time {
	return &t
}

func PBool(b bool) *bool {
	return &b
}

func PJson(b []byte) *json.RawMessage {
	rm := json.RawMessage(b)
	return &rm
}

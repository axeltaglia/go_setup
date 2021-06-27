package tkt

import (
	"encoding/json"
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"io"
	"net/http"
	"reflect"
	"runtime/debug"
)

func TransactionalLoggable(path string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, InterceptFatal(InterceptCORS(f)))
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

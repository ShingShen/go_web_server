package middleware

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func Method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		headerDefault(w, r, method)
		if r.Method != method {
			w.WriteHeader(405)
			return
		}
		next(w, r)
	}
}

func UploadFileMethod(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/form-data;boundary="+multipart.NewWriter(bytes.NewBufferString("")).Boundary())
		headerDefault(w, r, "PUT")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}
		next(w, r)
	}
}

func headerDefault(w http.ResponseWriter, r *http.Request, method string) {
	header := w.Header()
	methods := fmt.Sprintf("OPTIONS, %s", method)

	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	header.Set("Access-Control-Allow-Methods", methods)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

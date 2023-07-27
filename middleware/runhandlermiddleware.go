package middleware

import (
	"fmt"
	"net/http"
)

func RunHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		run := func(m chan bool, w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("Recovered: ", err)
				}
			}()

			defer close(m)
			next(w, r)
		}
		messages := make(chan bool)
		go run(messages, w, r, next)
		for message := range messages {
			fmt.Println(message)
		}
	}
}

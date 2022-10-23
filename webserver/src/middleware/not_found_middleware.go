package middleware

import (
	"goWebServer/src/http_helpers"
	"net/http"
)

func NotFoundMiddleWare(next http.Handler, e http_helpers.Endpoint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != e.GetPath() {
			http_helpers.ErrorHandler(w, r, http.StatusNotFound, "")
			return
		}
		next.ServeHTTP(w, r)
	})
}

package server

import (
	"log"
	"net/http"
	"strings"
	"freshmanual.com/template/status"
)

func Home(w http.ResponseWriter, r *http.Request) {
		SetHeaders(w)
		log.Print("Yo")
		w.Write([]byte("OK"))
}

func Router(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			SetHeaders(w)
			// w.Header().Set("content-type","text/html; charset=utf-8")
			status.S_404(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

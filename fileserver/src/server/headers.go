package server

import "net/http"

func SetHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "https://"+r.Host)
	w.Header().Set("Content-Language", "en")
	w.Header().Set("Content-Security-Policy", "default-src 'self'")
	w.Header().Set("cache-control", "private")
	w.Header().Set("host", "https://"+r.Host)
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
}

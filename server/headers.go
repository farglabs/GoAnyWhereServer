package server
import "net/http"

func SetHeaders(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Credentials","true")
	res.Header().Set("Access-Control-Allow-Origin","https://freshmanual.com")
	res.Header().Set("Content-Language","en")
	res.Header().Set("Content-Security-Policy","default-src 'self'")
	res.Header().Set("cache-control","private")
	res.Header().Set("host","https://freshmanual.com")
	res.Header().Set("Referrer-Policy","no-referrer")
	res.Header().Set("strict-transport-security","max-age=31536000; includeSubDomains; preload")
	res.Header().Set("X-Content-Type-Options","nosniff")
	res.Header().Set("X-Permitted-Cross-Domain-Policies","none")
	res.Header().Set("X-Frame-Options","SAMEORIGIN")
}
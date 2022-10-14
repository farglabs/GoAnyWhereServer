package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartSSL() {
	s := &http.Server{
		Addr:    ":443",
		Handler: nil,
	}

	mux := http.NewServeMux()

	// Initialize FILE SERVER
	fs := http.FileServer(http.Dir("/home/romulus/files"))
	mux.Handle("/files/", http.StripPrefix("/files", FileServer(fs)))

	// Initialize API
	api := http.HandlerFunc(Home)
	mux.Handle("/api/", ApiRouter(api))

	// REDIRECT TO HTTPS:302
	go func() {
		fmt.Println("HTTPS REDIRECT ACTIVE")
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectToTls)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// HTTPS SERVER
	s.Handler = mux
	fmt.Println("Server Running on", s.Addr)
	log.Fatal(
		s.ListenAndServeTLS(
			"/etc/letsencrypt/live/freshmanual.com/fullchain.pem",
			"/etc/letsencrypt/live/freshmanual.com/privkey.pem",
		))
}

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+":443"+r.RequestURI, 302)
}

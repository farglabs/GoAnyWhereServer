package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartSSL() {
	s := &http.Server{
		Addr:    ":55443",
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
	// go func() {
	// fmt.Println("HTTPS REDIRECT ACTIVE")
	// if err := http.ListenAndServe(":8080", http.HandlerFunc(redirectToTls)); err != nil {
	// 	log.Fatalf("ListenAndServe error: %v", err)
	// }
	// }()

	// HTTPS SERVER
	s.Handler = mux
	fmt.Println("Server Running on", s.Addr)
	log.Fatal(
		s.ListenAndServeTLS(
			"/home/fileserver/fullchain.pem",
			"/home/fileserver/privkey.pem",
		))
}

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+":443"+r.RequestURI, 302)
}

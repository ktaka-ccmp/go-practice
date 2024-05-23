package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/ktaka-ccmp/node-practice/go-practice-03/middleware"
)

func findByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Item ID: " + id))
}

func getLatest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Latest items"))
}

func hellofoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from dyna.h.ccmp.jp"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", findByID)
	router.HandleFunc("/item/latest", getLatest)
	router.HandleFunc("dyna.h.ccmp.jp/", hellofoo)
	router.HandleFunc("dyna.h.ccmp.jp/item/{id}", findByID)

	certFilePath := "./fullchain10.pem"
	keyFilePath := "./privkey10.pem"

	serverTLScert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		log.Fatalf("Error loading TLS cert: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLScert},
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   middleware.Logger2(router),
		TLSConfig: tlsConfig,
	}

	log.Println("Server is running on port 443")
	server.ListenAndServeTLS("", "")
}

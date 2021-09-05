package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"
)

const message = "Hola mundo ! Como estamos?"

var (
	apiPort = os.Getenv("API_PORT")
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	mux.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("oeoeoeoeo"))
	})

	srv := NewServer(mux, apiPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

}

func NewServer(mux *http.ServeMux, serverAddress string) *http.Server {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      mux,
	}

	return srv
}

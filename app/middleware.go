package main

import (
	"encoding/base64"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strings"
)

const Password = "yoni:davidson"

func SuperSecure(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if SECURITY_OPEN != "TRUE" {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Basic ") {
				pleaseAuth(w)
				return
			}

			password, err := base64.StdEncoding.DecodeString(auth[6:])
			if err != nil || string(password) != Password {
				pleaseAuth(w)
				return
			}
		} else {
			log.Println("warning - security open")
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func pleaseAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Go away!\n"))
}

func JsonText(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

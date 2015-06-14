package main

import (
	"fmt"
	"github.com/codedust/go-httpserve"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<!DOCTYPE html><html><h1>It works!</h1></html>")
	})

	// add authentication
	salt, err := httpserve.RandomString(32)
	if err != nil {
		panic("could not generate salt")
	}

	handleAuth := httpserve.BasicAuthHandler(http.DefaultServeMux, "user", httpserve.Sha512Sum("pass"+salt), salt)

  httpserve.CreateCertificateIfNotExist("./cert.pem", "./key.pem", "localhost", 3072)
  httpserve.ListenAndUpgradeTLS(":8080", "./cert.pem", "./key.pem", handleAuth)
}

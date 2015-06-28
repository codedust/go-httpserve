/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */
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

	httpserve.CreateCertificateIfNotExist("./cert.pem", "./key.pem", "localhost", 3072)

	fmt.Println("URL: http://localhost:8080")
	httpserve.ListenAndUpgradeTLS(":8080", "./cert.pem", "./key.pem", http.DefaultServeMux)
}

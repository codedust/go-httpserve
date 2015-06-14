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

	// add authentication
	salt, err := httpserve.RandomString(32)
	if err != nil {
		panic("could not generate salt")
	}

	handleAuth := httpserve.BasicAuthHandler(http.DefaultServeMux, "user", httpserve.Sha512Sum("pass"+salt), salt)
	http.ListenAndServe(":8080", handleAuth)
}

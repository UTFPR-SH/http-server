// Copyright 2018 Portal Direc's authors. All rights reserved.
//
// Authors: Rafael Campos Nunes <rafaelnunes@alunos.utfpr.edu.br>
//          ...                 <email>
//
// Use of this source code is governed by a Apache License. The license can be
// found (LICENSE) at the root of the project.

// DOCUMENTATION
// ---------------------------------------------------------------------------
//

package main

import (
	"log"
	"mime"
	"net/http"
	"strings"
	"time"
)

var (
	server http.Server
	pages  map[string]func(http.ResponseWriter, *http.Request)
)

type PortalDirecHandler struct{}

func Init(port string) {

	addr := "127.0.0.1"

	if deploy {
		addr = "0.0.0.0"
	}	

	server = http.Server{
		Addr:           addr + ":" + port,
		Handler:        &PortalDirecHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	pages = make(map[string]func(http.ResponseWriter, *http.Request))

	// TODO: This will have to be configured by a conf file
	pages["/"] = HandleRoot
	pages["/index"] = HandleRoot
	pages["/depec"] = HandleDepec
	pages["/depex"] = HandleDepex
	pages["/depet"] = HandleDepet
	pages["/oportunities"] = HandleOportunities

	err := server.ListenAndServe()

	if err != nil {
		log.Printf("Error: %s", err)
	}
}

func (*PortalDirecHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if h, ok := pages[r.URL.String()]; ok {
		h(w, r)
	} else {
		// Verify the MIME type of the requested file
		filepath := strings.Split(r.URL.String(), "/")
		filename := ""

		var mimeType string

		if len(filepath) > 2 {
			filename = filepath[len(filepath)-1]
			extension := strings.Split(filename, ".")

			mimeType = mime.TypeByExtension("." + extension[len(extension)-1])
		}

		if debug {
			log.Printf("Requested file %s. MIME type: %s\n", filename, mimeType)
		}

		// Loads the file and send it back with the correct MIME type
		file, err := Load(r.URL.String())

		if err == nil {
			w.Header().Add("Content-Type", mimeType)
			w.Write(file.Body)
		}
	}
}

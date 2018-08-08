// Copyright 2018 Portal Direc's authors. All rights reserved.
//
// Authors: Rafael Campos Nunes <rafaelnunes@alunos.utfpr.edu.br>
//          ...                 <email>
//
// Use of this source code is governed by a Apache License. The license can be
// found (LICENSE) at the root of the project.

// DOCUMENTATION
// ---------------------------------------------------------------------------
// This file concerns the definition of requests used by the server as
// also the render to the client.

package main

import (
	"html/template"
	"log"
	"net/http"
)

func handlePage(title string, w http.ResponseWriter, r *http.Request) {

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	} else {
		if debug {
			log.Printf("The page \"%s\" couldn't be loaded. Reason: %v",
				Filepath(title), err)
		}
	}
}

func render(w http.ResponseWriter, tmpl string, p *Page) {

	t, err := template.ParseFiles(tmpl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	handlePage("index", w, r)
}

func HandleDepex(w http.ResponseWriter, r *http.Request) {
	handlePage("depex", w, r)
}

func HandleDepec(w http.ResponseWriter, r *http.Request) {
	handlePage("depec", w, r)
}

func HandleDepet(w http.ResponseWriter, r *http.Request) {
	handlePage("depet", w, r)
}

func HandleOportunities(w http.ResponseWriter, r *http.Request) {
	handlePage("oportunities", w, r)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	handlePage("404", w, r)
}

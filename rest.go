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
	"net/http"
)

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

	title := "index"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}
}

func HandleDepex(w http.ResponseWriter, r *http.Request) {

	title := "depex"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}
}

func HandleDepec(w http.ResponseWriter, r *http.Request) {

	title := "depec"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}
}

func HandleDepet(w http.ResponseWriter, r *http.Request) {

	title := "depet"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}

}

func HandleOportunities(w http.ResponseWriter, r *http.Request) {

	title := "oportunities"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {

	title := "404"

	page, err := Load(title)

	if err == nil {
		render(w, Filepath(title), page)
	}
}

// Copyright 2018 Portal Direc's authors. All rights reserved.
//
// Authors: Rafael Campos Nunes <rafaelnunes@alunos.utfpr.edu.br>
//          ...                 <email>
//
// Use of this source code is governed by a Apache License. The license can be
// found (LICENSE) at the root of the project.

// DOCUMENTATION
// ---------------------------------------------------------------------------
// This file is responsible for describing a web page inside Go

package main

import (
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	path string = "./html/_site/"
)

/// Load a page and returns a pointer to it in case of success
func Load(title string) (*Page, error) {

	body, err := ioutil.ReadFile(Filepath(title))

	// FIXME: Do a proper error handling over here (return 404?)
	if err != nil {
		log.Printf("Couldn't load file %s", title)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

/// Returns the file path of a given file name
func Filepath(filename string) string {
	return path + filename + ".html"
}

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
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	root string = "./html/_site/"
)

// Load a page and returns a pointer to it in case of success
func Load(title string) (*Page, error) {

	body, err := ioutil.ReadFile(Filepath(title))

	if err != nil {
		log.Printf("Error %v", err.Error())
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Returns the file path of a given html file name
func Filepath(filename string) string {

	split := strings.Split(filename, ".")

	// The file isn't an HTML and probably is an asset
	if len(split) > 1 {
		log.Printf("Asset to load %s", filename)
		return root + filename[1:]
	}

	return root + filename + ".html"
}

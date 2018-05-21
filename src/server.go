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
	"net/http"
)

func Init(addr string) {

	// Set handlers entry
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/depec", HandleDepec)
	http.HandleFunc("/depex", HandleDepex)
	http.HandleFunc("/depet", HandleDepet)
	http.HandleFunc("/oportunities", HandleOportunities)
	http.HandleFunc("/404", HandleNotFound)

	err := http.ListenAndServe(":"+addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}

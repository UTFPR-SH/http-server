// Copyright 2018 Portal Direc's authors. All rights reserved.
//
// Authors: Rafael Campos Nunes <rafaelnunes@alunos.utfpr.edu.br>
//          ...                 <email>
//
// Use of this source code is governed by a Apache License. The license can be
// found (LICENSE) at the root of the project.

// DOCUMENTATION
// ---------------------------------------------------------------------------
// The entry point to the entire system. It interprets flags and start the
// server at a desired <port> value, here called of address, or if no port
// value is assigned the default one is 8080.

package main

import (
	"flag"
)

var (
	addr   string
	html   string
	debug  bool
	deploy bool
)

func init() {
	flag.StringVar(&addr, "p", "8080", "Set port address used to start the server")
	flag.StringVar(&html, "html", "./html/_site", "Set the default directory for Jekyll output files")
	flag.BoolVar(&debug, "d", false, "Set if the server is in debug mode")
	flag.BoolVar(&deploy, "deploy", false, "Deploy the server over the network")
}

func main() {
	flag.Parse()

	//Init the server at port address
	Init(addr)
}

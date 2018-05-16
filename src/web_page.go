package main

import (
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func Load(title string) (*Page, error) {

	filename := title
	body, err := ioutil.ReadFile(filename)

	// FIXME: Do a proper error handling over here (return 404?)
	if err != nil {
		log.Printf("Couldn't load file %s", filename)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

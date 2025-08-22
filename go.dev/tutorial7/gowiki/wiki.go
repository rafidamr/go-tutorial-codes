package main

import (
	"fmt"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	log.SetPrefix("tutorial 7: ")
	log.SetFlags(0)
	title := "a page title"
	page := &Page{Title: title, Body: []byte("a page body")}
	err := page.save()
	if err != nil {
		log.Fatal(err)
	}
	p, err := load(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title=%q Body=%q", p.Title, string(p.Body))
}

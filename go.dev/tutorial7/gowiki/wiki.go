package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

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

func renderPage(w http.ResponseWriter, templateName string, p *Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Path[len("/view/"):]
	p, err := load(pageName)
	if err != nil {
		http.Redirect(w, r, "/edit/"+pageName, http.StatusFound)
	}
	renderPage(w, "view", p)
}

func editViewHandler(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Path[len("/edit/"):]
	p, err := load(pageName)
	if err != nil {
		p = &Page{Title: pageName}
	}
	renderPage(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: pageName, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+pageName, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editViewHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

package main

import (
	"html/template"
	"net/http"

	"io"

	"log"

	"github.com/ccaneke/webAnalytics/webAnalyzer"
)

func loadPage(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *webAnalyzer.PageInformation) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("htmlversion")
	reader, err := loadPage(url)
	if err != nil {
		return
	}
	pgInfo, err := webAnalyzer.GetPageInformation(reader)
	if err != nil {
		return
	}
	renderTemplate(w, "view", &pgInfo)
}

func main() {
	http.HandleFunc("/index/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"html/template"
	"net/http"

	"io"

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

func inputHandler(w http.ResponseWriter, r *http.Request) {

}

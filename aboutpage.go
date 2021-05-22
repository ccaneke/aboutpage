package main

import (
	"html/template"
	"net/http"

	"github.com/ccaneke/webAnalytics/webAnalyzer"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pgInfo := webAnalyzer.PageInformation{}
	renderTemplate(w, "index", &pgInfo)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *webAnalyzer.PageInformation) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

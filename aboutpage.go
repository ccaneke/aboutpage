package main

import (
	"html/template"
	"net/http"

	"github.com/ccaneke/webAnalytics/webAnalyzer"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmplname := r.URL.Path[0:]
	pgInfo := webAnalyzer.PageInformation{}
	renderTemplate(w, tmplname, &pgInfo)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *webAnalyzer.PageInformation) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

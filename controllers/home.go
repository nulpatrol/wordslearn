package controllers

import (
	"html/template"
	"net/http"
)

type HomeHandler struct {

}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		WordForms: fetch(db),
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, data)
}

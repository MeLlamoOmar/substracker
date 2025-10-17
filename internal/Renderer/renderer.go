package Renderer

import (
	"html/template"
	"net/http"
)

type Renderer struct {
	Templates *template.Template
}

func NewRenderer() *Renderer {
	tmpl := template.New("")
	var err error

	tmpl, err = tmpl.ParseGlob("web/templates/*.html")
	if err != nil {
		panic(err)
	}

	return &Renderer{
		Templates: tmpl,
	}
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data interface{}) {
	err := r.Templates.ExecuteTemplate(w, name, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (r *Renderer) RenderWithStatus(w http.ResponseWriter, name string, data interface{}, status int) {
	w.WriteHeader(status)
	err := r.Templates.ExecuteTemplate(w, name, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
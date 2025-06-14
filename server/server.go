package server

import (
	"asciiartweb/handlers"
	"html/template"
	"log"
	"net/http"
)

func WithRecovery(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
				handlers.InternalServerErrorHandler(w, r)
			}
		}()
		h(w, r)
	}
}

func RouteHandler(tmpl, tmplError *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			handlers.IndexHandler(tmpl)(w, r)
		default:
			handlers.NotFoundHandler(w, r)
		}
	}
}

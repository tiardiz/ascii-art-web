package server

import (
	"asciiartweb/handlers"
	"html/template"
	"log"
	"net/http"
)

// WithRecovery — middleware для обработки паник
func WithRecovery(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic:", err)
				http.Error(w, "500 - Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
		}()
		h(w, r)
	}
}

// RouteHandler — централизованный роутер
func RouteHandler(tmpl, tmpl404 *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			handlers.IndexHandler(tmpl)(w, r)
		case "/submit":
			handlers.SubmitHandler(tmpl)(w, r)
		default:
			handlers.NotFoundHandler(tmpl404)(w, r)
		}
	}
}

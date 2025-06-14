package main

import (
	"asciiartweb/handlers"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template
var tmpl404 *template.Template

func main() {
	var err error
	tmpl404, err = template.ParseFiles("templates/404.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона 404: %v", err)
	}

	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %v", err)
	}

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Страница не найдена", http.StatusNotFound)
	})

	http.HandleFunc("/", routeHandler) // заменяет indexHandler

	http.HandleFunc("/submit", withRecovery(handlers.SubmitHandler(tmpl)))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущен на : http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
func withRecovery(h http.HandlerFunc) http.HandlerFunc {
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

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.IndexHandler(tmpl)(w, r)
	case "/submit":
		handlers.SubmitHandler(tmpl)(w, r)
	default:
		handlers.NotFoundHandler(tmpl404)(w, r)
	}
}

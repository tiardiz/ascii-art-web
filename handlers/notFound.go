package handlers

import (
	"html/template"
	"net/http"
)

func NotFoundHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Ошибка отображения страницы 404", http.StatusInternalServerError)
		}
	}
}

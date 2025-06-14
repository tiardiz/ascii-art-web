package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Input string
	ASCII string
	Style string
}

func IndexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Input: "",
			ASCII: "",
			Style: "",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
		}
	}
}

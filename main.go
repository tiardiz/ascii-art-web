package main

import (
	"asciiartweb/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	var err error

	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %v", err)
	}

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 - Страница не найдена", http.StatusNotFound)
	})

	http.HandleFunc("/", routeHandler) // заменяет indexHandler

	http.HandleFunc("/submit", withRecovery(submitHandler))

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
		indexHandler(w, r)
	case "/submit":
		submitHandler(w, r)
	default:
		http.NotFound(w, r) // 404
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	data := struct {
		Input string
		ASCII string
		Style string
	}{
		Input: "",
		ASCII: "",
		Style: "",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Ошибка шаблона: %v", err)
		http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}

	text := r.FormValue("username")
	if text == "" {
		text = "" // если поле пустое
	}

	style := r.FormValue("style") // Получаем выбранный стиль

	standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyHash := "242fdef5fad0fe9302bad1e38f0af4b0f83d086e49a4a99cdf0e765785640666"

	filepath := "banners/" + style + ".txt"
	hashBytes, err := asciiart.CalculateFileHash(filepath)
	if err != nil {
		fmt.Printf("error: cannot calculate hash for file %s\n", filepath)
		return
	}

	hashString := fmt.Sprintf("%x", hashBytes)
	if hashString == standardHash || hashString == shadowHash || hashString == thinkertoyHash {
		asciiText := asciiart.ASCIIart(text, style)
		data := struct {
			Input string
			ASCII string
			Style string
		}{
			Input: text,
			ASCII: asciiText,
			Style: style,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Ошибка генерации ASCII", http.StatusInternalServerError)
		return
	}

}

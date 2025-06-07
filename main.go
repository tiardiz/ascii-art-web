package main

import (
	"asciiartweb/asciiart"
	"html/template"
	"log"
	"net/http"
)

// Глобальная переменная шаблона
var tmpl *template.Template

func main() {
	// Загружаем шаблон из файла
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %v", err)
	}

	// Обработчик главной страницы
	http.HandleFunc("/", indexHandler)

	// Обработчик формы
	http.HandleFunc("/submit", submitHandler)

	log.Println("Сервер запущен на : http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name  string
		ASCII string
	}{
		Name:  "Гость",
		ASCII: "",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Ошибка шаблона: %v", err)
		http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		username = "" // если поле пустое
	}

	style := r.FormValue("style") // Получаем выбранный стиль

	// Генерируем ASCII-арт для введённого имени и выбранного стиля
	asciiText := asciiart.ASCIIart(username, style) // Передаем стиль в функцию

	data := struct {
		Name  string
		ASCII string
	}{
		Name:  username,
		ASCII: asciiText,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
	}
}

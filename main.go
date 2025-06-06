package main

import (
	"asciiartweb/asciiart"
	"log"
	"net/http"
	"text/template"
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

// handler для главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{
		Name: "Гость", // Здесь можно задать значение по умолчанию
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
	}
}

// handler для обработки формы
// func submitHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	if username == "" {
// 		username = "" // если поле пустое
// 	}

// 	// Генерируем ASCII-арт для введённого имени (позволяем несколько строк)
// 	asciiText := asciiart.ASCIIart(username)

// 	// Структура, передаваемая в шаблон
// 	data := struct {
// 		Name  string
// 		ASCII string
// 	}{
// 		Name:  username,
// 		ASCII: asciiText,
// 	}

// 	// Передаем данные в шаблон
// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
// 	}
// }

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

// func ASCIIart(input string) string {
// 	// Читаем файл с шаблоном ASCII-арт
// 	inputfile, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		return "error reading file"
// 	}

// 	// Разбиваем по строкам с учётом UNIX-формата
// 	inputfileLines := strings.Split(string(inputfile), "\n")

// 	// Логируем количество строк (можно убрать потом)
// 	log.Printf("Lines in standard.txt: %d", len(inputfileLines))

// 	// Разбиваем введённый текст по "\n" (обрати внимание, что у тебя двойной обратный слеш)
// 	words := strings.Split(input, "\n")

// 	result := ""

// 	for _, word := range words {
// 		// В твоём файле по 9 строк на каждый символ (8 строк + 1 пустая?)
// 		for i := 0; i < 8; i++ {
// 			for _, char := range word {
// 				index := i + (int(char-' ') * 9) + 1
// 				if index < 0 || index >= len(inputfileLines) {
// 					// Если вышли за пределы, просто пропускаем символ
// 					continue
// 				}
// 				result += inputfileLines[index]
// 			}
// 			result += "\n"
// 		}
// 	}

// 	return result
// }

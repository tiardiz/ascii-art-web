package main

import (
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

    log.Println("Сервер запущен на :8080")
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
        username = "Гость"
    }

    // Передаём введённое имя обратно в шаблон и показываем страницу заново
    data := struct {
        Name string
    }{
        Name: username,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Ошибка при отображении страницы", http.StatusInternalServerError)
    }
}

package handlers

import (
	"asciiartweb/asciiart"
	"fmt"
	"html/template"
	"net/http"
)

func SubmitHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		style := r.FormValue("style")

		standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
		shadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
		thinkertoyHash := "242fdef5fad0fe9302bad1e38f0af4b0f83d086e49a4a99cdf0e765785640666"

		filepath := "banners/" + style + ".txt"
		hashBytes, err := asciiart.CalculateFileHash(filepath)
		if err != nil {
			http.Error(w, "Ошибка чтения файла баннера", http.StatusInternalServerError)
			return
		}

		hashString := fmt.Sprintf("%x", hashBytes)
		if hashString == standardHash || hashString == shadowHash || hashString == thinkertoyHash {
			asciiText := asciiart.ASCIIart(text, style)
			data := PageData{
				Input: text,
				ASCII: asciiText,
				Style: style,
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "Ошибка генерации ASCII", http.StatusInternalServerError)
		}
	}
}

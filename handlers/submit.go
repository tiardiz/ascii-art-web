package handlers

import (
	"asciiartweb/asciiart"
	"fmt"
	"html/template"
	"net/http"
)

var tmplError *template.Template

func isValidText(text string) bool {
	hasVisibleChar := false
	for _, ch := range text {
		if ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r' {
			continue
		}
		if ch < 33 || ch > 126 {
			return false
		}
		hasVisibleChar = true
	}
	return hasVisibleChar
}

func SubmitHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedHandler(w, r)
			return
		}
		err := r.ParseForm()
		if err != nil {
			BadRequestHandler(w, r)
			return
		}

		text := r.FormValue("username")
		style := r.FormValue("style")

		// fmt.Printf("DEBUG text with quotes: %#q\n", text)
		// fmt.Println("isValidText:", isValidText(text))
		if text == "" || style == "" || !isValidText(text) {
			BadRequestHandler(w, r)
			return
		}

		standardHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
		shadowHash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
		thinkertoyHash := "242fdef5fad0fe9302bad1e38f0af4b0f83d086e49a4a99cdf0e765785640666"

		filepath := "banners/" + style + ".txt"
		hashBytes, err := asciiart.CalculateFileHash(filepath)
		if err != nil {
			InternalServerErrorHandler(w, r)
			return
		}

		hashString := fmt.Sprintf("%x", hashBytes)
		if hashString != standardHash && hashString != shadowHash && hashString != thinkertoyHash {
			InternalServerErrorHandler(w, r)
			return
		}

		asciiText := asciiart.ASCIIart(text, style)
		data := PageData{
			Input: text,
			ASCII: asciiText,
			Style: style,
		}
		if err = tmpl.Execute(w, data); err != nil {
			InternalServerErrorHandler(w, r)
		}

	}

}

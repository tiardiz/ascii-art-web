// package asciiart

// import (
// 	"log"
// 	"os"
// 	"strings"
// )

// func ASCIIart(input string) string {
// 	// Читаем файл с шаблоном ASCII-арт
// 	inputfile, err := os.ReadFile("asciiart/banners/standard.txt")
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

package asciiart

import (
	"log"
	"os"
	"strings"
)

func ASCIIart(input string, style string) string {
	// Определяем путь к файлу в зависимости от выбранного стиля
	var fileName string
	switch style {
	case "standard":
		fileName = "banners/standard.txt"
	case "shadow":
		fileName = "banners/shadow.txt"
	case "thinkertoy":
		fileName = "banners/thinkertoy.txt"
	}

	// Читаем файл с шаблоном ASCII-арт
	inputfile, err := os.ReadFile(fileName)
	if err != nil {
		return "error reading file"
	}

	// Разбиваем по строкам с учётом UNIX-формата
	inputfileLines := strings.Split(string(inputfile), "\n")

	// Логируем количество строк (можно убрать потом)
	log.Printf("Lines in %s: %d", fileName, len(inputfileLines))

	// Разбиваем введённый текст по "\n"
	words := strings.Split(input, "\n")

	result := ""

	for _, word := range words {
		// В твоём файле по 9 строк на каждый символ (8 строк + 1 пустая?)
		for i := 0; i < 8; i++ {
			for _, char := range word {
				index := i + (int(char-' ') * 9) + 1
				if index < 0 || index >= len(inputfileLines) {
					// Если вышли за пределы, просто пропускаем символ
					continue
				}
				result += inputfileLines[index]
			}
			result += "\n"
		}
	}

	return result
}

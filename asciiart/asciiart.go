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
	"os"
	"strings"
)

// GetFile считывает содержимое файла с баннером и возвращает его построчно.
func GetFile(name string) ([]string, error) {
	content, err := os.ReadFile("banners/" + name + ".txt")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

// DrawText рисует ASCII-арт для входной строки input, используя стиль баннера style.
func ASCIIart(input, style string) string {
	fileContent, err := GetFile(style)
	if err != nil {
		return ""
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\\n", "\n")

	lines := strings.Split(input, "\n")

	// Массив для хранения всех строк ASCII-арт
	var asciiArtLines []string

	// Количество строк для каждого символа
	nbrLinesPerChar := 9
	matrix := [][]string{}

	// Обрабатываем каждый символ в строках
	for _, line := range lines {
		for _, ch := range line {
			startLine := (int(ch) - 32) * nbrLinesPerChar
			if startLine < 0 || startLine+nbrLinesPerChar > len(fileContent) {
				return ""
			}
			matrix = append(matrix, fileContent[startLine:startLine+nbrLinesPerChar])
		}

		// Объединяем все части символов для текущей строки
		asciiArtLine := combineCharacters(matrix)
		asciiArtLines = append(asciiArtLines, strings.Join(asciiArtLine, "\n"))
		matrix = nil // Очищаем для следующей строки
	}

	// Возвращаем итоговый ASCII-арт
	return strings.Join(asciiArtLines, "\n")
}

//Ascii 10: Line Feed (LF) -  `\n`
//Ascii 13: Carriage Return (CR) - `\r`

// combineCharacters формирует строки ASCII-арта, объединяя по горизонтали части символов.
func combineCharacters(matrix [][]string) []string {
	var result []string
	// Строки 1..8 включительно (0 — пустая или разделитель)
	for i := 1; i < 9; i++ {
		var lineBuilder strings.Builder
		for j := 0; j < len(matrix); j++ {
			lineBuilder.WriteString(matrix[j][i])
		}
		line := lineBuilder.String()
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}

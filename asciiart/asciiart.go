package asciiartweb

import (
	"fmt"
	"os"
	"strings"
)

func GetFile(name string) []string {
	content, err := os.ReadFile("banners/" + name + ".txt")
	if err != nil {
		return nil
	}
	return strings.Split(string(content), "\n")
}

func DrawText(txt string, name string) {
	matrix := [][]string{}
	fileContent := GetFile(name)

	for _, v := range txt {

		nbr := 9
		line := (int(v) - 32) * nbr
		if line > len(fileContent)-nbr || line < 0 {
			fmt.Println("error: undefined character")
			return
		}
		matrix = append(matrix, fileContent[line:line+nbr])
	}
	lines := removeEmptyStrings(PrintChar(matrix))
	fmt.Print(strings.Join(lines, "\n"))

}

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func PrintChar(matrix [][]string) []string {
	result := []string{}

	for i := 1; i < 9; i++ {
		line := ""
		for j := 0; j < len(matrix); j++ {
			line += matrix[j][i]
		}
		result = append(result, line)
	}

	return result
}
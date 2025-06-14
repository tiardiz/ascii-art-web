package asciiart

import (
	"os"
	"strings"
)

func GetFile(name string) ([]string, error) {
	content, err := os.ReadFile("banners/" + name + ".txt")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func fmtFilecontentLines(content []string) []string {
	for i, line := range content {
		line = strings.ReplaceAll(line, "\r", "")
		content[i] = line
	}
	return content
}

func ASCIIart(input, style string) string {
	fileContent, err := GetFile(style)
	if err != nil {
		return ""
	}

	fmtFilecontent := fmtFilecontentLines(fileContent)
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\\r", "")
	input = strings.ReplaceAll(input, "\\t", "    ")
	lines := strings.Split(input, "\n")

	nbrLinesPerChar := 9
	var result strings.Builder

	for _, line := range lines {
		for i := 1; i < nbrLinesPerChar; i++ {
			for _, ch := range line {
				startLine := (int(ch) - 32) * nbrLinesPerChar
				if startLine < 0 || startLine+nbrLinesPerChar > len(fmtFilecontent) {
					continue
				}
				result.WriteString(fmtFilecontent[startLine+i])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}

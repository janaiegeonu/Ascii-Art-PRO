package ascii

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	input = strings.ReplaceAll(input, "\\n", `\n`)

	parts := strings.Split(input, `\n`)

	result := ""
	for i, word := range parts {
		if word == "" {
			if i < len(parts)-1 {
				result += ("\n")
			}
			continue
		}

		row := Render(word, banner)
		for _, line := range row {
			result += (line + "\n")
		}
	}

	return result
}

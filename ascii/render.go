package ascii

import (
	"Ascii-art-PRO/style"
	"fmt"
)

func Render(input string, banner map[rune][]string) []string {

	result := make([]string, 8)

	for _, char := range input {

		art, ok := banner[char]
		if ok {
			for row := 0; row < 8; row++ {
				result[row] += art[row]
			}
		} else {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: invalid character"))))
			continue
		}

	}
	return result
}

package ascii

import (
	"Ascii-art-PRO/style"
	"fmt"
	"os"
	"strings"
)

func GetBanner(filename string) map[rune][]string {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
		return nil
	}

	if len(data) == 0 {
		fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: empty banner file"))))
		return nil

	}

	line := strings.Split(string(data), "\n")
	char := rune(32)
	index := 1
	font := make(map[rune][]string)

	for char <= 126 {
		if index+8 > len(line) {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: invalid banner file content"))))
			break
		}
		font[char] = line[index : index+8]
		index += 9
		char++

	}
	if len(font) != 95 {
		fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: invalid banner file"))))
		return nil
	}

	return font
}

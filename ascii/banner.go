package ascii

import (
	"Ascii-art-PRO/style"
	"fmt"
	"os"
	"strings"
)

func GetBanner(filename string) (map[rune][]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
	}

	if len(data) == 0 {
		return nil, fmt.Errorf(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: empty banner file"))))

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

		return nil, fmt.Errorf(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: invalid banner file"))))
	}

	return font, nil
}

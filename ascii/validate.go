package ascii

import (
	"Ascii-art-PRO/style"
	"fmt"
)

func Validate(input string) (rune, error) {

	for _, char := range input {
		if char != '\n' && (char < 32 || char > 126) {
			return char, fmt.Errorf(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: text character doesn't have a ASCII value"))))
		}
	}
	return 0, nil
}

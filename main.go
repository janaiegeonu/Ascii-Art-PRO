package main

import (
	"Ascii-art-PRO/animation"
	"Ascii-art-PRO/style"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	animation.TypeWriter(style.BlueBackground(style.Bold("\t           ASCII-ART CLI PROGRAM           ")))
	(animation.Blink("      \t\t\t "))
	fmt.Println(style.Grey("___________________________________________________________________________"))

	animation.TypeWriter(style.White(style.Italic("\n▶ Enter your fullname to start the program..")))

	reader := bufio.NewReader(os.Stdin)

	var fullname string
	for {
		fmt.Println()
		fmt.Print(style.WhiteBackground(style.Bold("Fullname :")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if !strings.Contains(input, " ") {
			fmt.Println(style.Red(style.Dim("ERROR : input your lastname too")))
			continue
		}

		break

	}
	fmt.Println(fullname)

}

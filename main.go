package main

import (
	"Ascii-art-PRO/animation"
	"Ascii-art-PRO/style"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	animation.TypeWriter(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
	(animation.Blink("      \t\t\t "))
	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriter(style.Green(style.Italic("\n▶ Enter your fullname to start the program..")))

	fmt.Println(style.Green("___________________________________________________________________________"))

	reader := bufio.NewReader(os.Stdin)

	var fullname string
	for {
		fmt.Println()
		fmt.Print(style.Blue(style.Bold("⚬ Fullname : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: fullname can't be empty"))))
			continue

		}

		num, _ := regexp.MatchString(`\d`, input)
		sym, _ := regexp.MatchString(`[[:punct:]]`, input)

		if num || sym {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: only alphabetic format is allowed"))))
			continue

		}

		if !strings.Contains(input, " ") {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: input your lastname too"))))
			continue
		}
		fullname = input

		break

	}
	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriter(style.Green("⬩\nNAME HAS BEEN SAVED SUCCESSFULLY 🗹"))

	fmt.Println(style.Green("___________________________________________________________________________"))
	fmt.Println(fullname)
	time.Sleep(2000 * time.Millisecond)
	fmt.Print("\033[3J\033[H\033[2J")
	fmt.Println()
	fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
	fmt.Println("\t\t\t( ｡" + style.Yellow("◕") + "‿‿" + style.Yellow("◕") + "｡)")

	fmt.Println()
	animation.Loading("● INSTRUCTIONS & README")
	fmt.Println(style.Yellow("___________________________________________________________________________\n"))
	animation.TypeWriter(style.Yellow(style.Dim("This an Ascii-art CLI program to print out text into a beautiful Ascii-Art font.")))
	animation.TypeWriter(style.Yellow(style.Dim("it has three Ascii-art font style with are ; ")))
	animation.TypeWriter(style.Yellow(style.Dim("1. standard font")))
	animation.TypeWriter(style.Yellow(style.Dim("2. shadow font")))
	animation.TypeWriter(style.Yellow(style.Dim("3. thinkertoy font\n")))
	animation.TypeWriter(style.Yellow(style.Dim("\tFONT SAMPLES")))
	animation.TypeWriter(style.Yellow(style.Dim("standard font :")))

	fmt.Println(style.Red(`
  _                               _   
 (_)                             (_)  
  _     __ _    _ __      __ _    _   
 | |   / _` + "`" + ` |  | '_ \    / _` + "`" + ` |  | |  
 | |  | (_| |  | | | |  | (_| |  | |  
 | |   \__,_|  |_| |_|   \__,_|  |_|  
_/ |                                  
|__/                                 
`))

}

package main

import (
	"Ascii-art-PRO/animation"
	"Ascii-art-PRO/ascii"
	"Ascii-art-PRO/style"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func Case(text string) string {
	parts := strings.Fields(text)
	for i, word := range parts {
		word = strings.ToLower(word)
		parts[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])

	}
	return strings.Join(parts, " ")
}

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
		fullname = Case(input)

		break

	}
	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriter(style.Green("\nNAME HAS BEEN SAVED SUCCESSFULLY 🗹"))

	fmt.Println(style.Green("___________________________________________________________________________"))

	fmt.Println()
	fmt.Println(style.Grey("___________________________________________________________________________"))

	fmt.Println(style.GreyBackground(" ENTER YOUR OPTION TO CONTINUE "))
	fmt.Println()
	fmt.Println(style.Grey("1. Enter key = To start main program"))
	fmt.Println(style.Grey("2. End = To Terminate the program"))
	fmt.Println()
	fmt.Println("\033[1m\033[42m"+" ENTER KEY "+"\033[0m"+"\033[0m", "|====|", "\033[1m\033[41m"+" END "+"\033[0m")

	var option string

	for {
		fmt.Println()
		fmt.Print(style.Blue("\nOPTION : "))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		option = Case(input)

		if option == "" {
			break

		} else if option == "End" {
			fmt.Println()
			fmt.Println(style.Red("___________________________________________________________________________"))
			animation.Loading1("TERMINATING PROGRAM")
			fmt.Println()
			animation.ProgressBarRed(30)
			fmt.Print("\033[3J\033[H\033[2J")
			return
		} else {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Invalid Option"))))
			continue
		}

	}

	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriterSlow(style.Green(style.Italic("\n▶ NOW LET'S DIVE INTO THE PROGRAM FULLY : " + fullname)))

	fmt.Println(style.Green("___________________________________________________________________________"))
	time.Sleep(600 * time.Millisecond)
	fmt.Println()
	animation.ProgressBarGreen(40)
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\033[3J\033[H\033[2J")

	fmt.Println()
	fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
	(animation.Blink("      \t\t\t "))
	fmt.Println()
	fmt.Print("◉ ")
	fmt.Println(style.GreyBackground(style.Bold("USER : " + fullname + "\n")))
	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriter(style.Green(style.Italic("\n▶ Enter your input text...")))

	fmt.Println(style.Green("___________________________________________________________________________"))

	var inputText string

	for {

		fmt.Println()
		fmt.Print(style.Blue(style.Bold("⚬ TEXT : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Text can't be empty"))))
			continue
		}

		if input == "None" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Text must have an input"))))
			continue
		}

		if strings.HasPrefix(input, "\\n") && strings.HasSuffix(input, "\\n") {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: include text with newline "))))
			continue
		}

		badchar, err := ascii.Validate(input)
		if err != nil {
			fmt.Printf("\033[31m\033[4m\033[53m"+"ERROR ⚠︎: %c invalid character\n", badchar)
			fmt.Println("\033[0m")
			continue
		}

		inputText = string(badchar)

		break

	}

	var font string

	for {

		fmt.Println()
		fmt.Print(style.Blue(style.Bold("⚬ FONT : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Can't be empty"))))
			continue

		}
		font = Case(input)

		switch font {
		case "Standard":
		case "Shadow":
		case "Thinkertoy":
		default:
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: font style not found"))))
			fmt.Println(style.Yellow(style.Upperline(style.Underline("SUPPORTED 🖒 : Standard, Shadow, Thinkertoy"))))
			continue
		}

		break
	}
	var colour string

	for {

		fmt.Println()
		fmt.Print(style.Blue(style.Bold("⚬ COLOUR : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Can't be empty"))))
			continue

		}

		colour = Case(input)

		switch colour {

		case "Yellow":
			colour = "\033[33m"
		case "Red":
			colour = "\033[31m"
		case "Blue":
			colour = "\033[34m"
		case "Green":
			colour = "\033[32m"
		case "Magenta":
			colour = "\033[35m"
		case "Cyan":
			colour = "\033[36m"
		case "Dark Red":
			colour = "\033[2m\033[31m"
		case "Dark Yellow":
			colour = "\033[2m\033[33m"
		case "Purple":
			colour = "\033[2m\033[35m"
		case "White":
			colour = "\033[37m"
		case "Grey":
			colour = "\033[30m"
		case "None":
			colour = ""
		default:
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Colour style not found"))))
			fmt.Println(style.Yellow(style.Upperline(style.Underline("SUPPORTED 🖒 : Red, Green, Blue, Yellow, Magenta, Cyan, Purple, Dark Red, Dark Yellow, White, Grey"))))
			continue

		}

		break
	}

	var Style string

	for {

		fmt.Println()
		fmt.Print(style.Blue(style.Bold("⚬ STYLE : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: Can't be empty"))))
			continue

		}

		Style = Case(input)

		switch Style {
		case "Bold":
			Style = "\033[1m"
		case "Dim":
			Style = "\033[2m"
		case "Italic":
			Style = "\033[3m"
		case "Stripes":
			Style = "\033[4m"
		case "None":
			Style = ""
		default:

			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: style not found"))))
			fmt.Println(style.Yellow(style.Upperline(style.Underline("SUPPORTED 🖒 : Bold, Dim, Italic, Stripes"))))
			continue

		}

		break

	}

	input := inputText

	if font == "Standard" {

		banner, err := ascii.GetBanner("standard.txt")
		if err != nil {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
			return

		}

		Output := ascii.GenerateArt(input, banner)
		fmt.Print(colour + Style + Output + "\033[0m")
	}

	if font == "Shadow" {

		banner, err := ascii.GetBanner("shadow.txt")
		if err != nil {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
			return

		}

		Output := ascii.GenerateArt(input, banner)
		fmt.Print(colour + Style + Output + "\033[0m")
	}

	if font == "Thinkertoy" {

		banner, err := ascii.GetBanner("thinkertoy.txt")
		if err != nil {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
			return

		}

		Output := ascii.GenerateArt(input, banner)

		fmt.Print(colour + Style + Output + "\033[0m")

	}

}

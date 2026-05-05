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
	animation.TypeWriter(style.Yellow(style.Bold("> DO YOU WANT TO READ INSTRUCTIONS BEFORE DIVING INTO THE PROGRAM ? (YES/NO)")))

	var ReadME string
	var option string

	for {

		fmt.Println()
		fmt.Print(style.Blue(style.Bold("OPTION : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: option can't be empty"))))
			continue
		}

		ReadME = Case(input)

		switch ReadME {
		case "No":
			goto start
		case "Yes":
			goto start1
		default:
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: option not found"))))
			continue
		}

	}

start1:
	time.Sleep(700 * time.Millisecond)
	fmt.Print("\033[3J\033[H\033[2J")
	fmt.Println()
	fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
	fmt.Println("\t\t\t( ｡" + style.Yellow("◕") + "‿‿" + style.Yellow("◕") + "｡)")

	fmt.Println()
	animation.Loading("● INSTRUCTIONS & README")
	fmt.Println(style.Yellow("___________________________________________________________________________\n"))
	animation.TypeWriterSlow(style.Yellow(style.Dim("This an Ascii-art CLI program to print out text into a beautiful Ascii-Art font.")))
	animation.TypeWriterSlow(style.Yellow(style.Dim("it has three Ascii-art font style with are ; ")))
	animation.TypeWriterSlow(style.White("1. standard font"))
	animation.TypeWriterSlow(style.White("2. shadow font"))
	animation.TypeWriterSlow(style.White("3. thinkertoy font\n"))
	time.Sleep(600 * time.Millisecond)
	animation.TypeWriterSlow(style.Yellow(style.Dim(style.Bold("\t ✳ FONT SAMPLES ✳ "))))
	fmt.Println()
	time.Sleep(300 * time.Millisecond)
	animation.TypeWriterSlow(style.GreyBackground(style.Dim("⦁ standard font :")))
	time.Sleep(500 * time.Millisecond)

	fmt.Println(style.White(`
  _                               _   
 (_)                             (_)  
  _     __ _    _ __      __ _    _   
 | |   / _` + "`" + ` |  | '_ \    / _` + "`" + ` |  | |  
 | |  | (_| |  | | | |  | (_| |  | |  
 | |   \__,_|  |_| |_|   \__,_|  |_|  
_/ |                                  
|__/                                 
`))
	fmt.Println(style.Grey("___________________________________________________________________________"))

	time.Sleep(600 * time.Millisecond)

	animation.TypeWriterSlow(style.GreyBackground(style.Dim("⦁ shadow font :")))
	time.Sleep(500 * time.Millisecond)

	fmt.Println(style.White(`
                                       
  _|                                _|  
        _|_|_|  _|_|_|      _|_|_|      
  _|  _|    _|  _|    _|  _|    _|  _|  
  _|  _|    _|  _|    _|  _|    _|  _|  
  _|    _|_|_|  _|    _|    _|_|_|  _|  
  _|                                    
_|                                     
`))

	fmt.Println(style.Grey("___________________________________________________________________________"))

	time.Sleep(600 * time.Millisecond)
	animation.TypeWriterSlow(style.GreyBackground(style.Dim("⦁ thinkertoy font :")))
	time.Sleep(500 * time.Millisecond)

	fmt.Println(style.White(`
    o                    o  
        oo   o-o    oo      
    o  | |   |  |  | |   |  
    |  o-o-  o  o  o-o-  |  
o   o                       
 o-o                       
`))
	fmt.Println(style.Grey("___________________________________________________________________________"))

	time.Sleep(300 * time.Millisecond)

	animation.Rainbow("\t✳ ASCII-ART CUSTOMIZATION ✳")
	fmt.Println()
	fmt.Println("\033[0m")
	animation.TypeWriterSlow(style.Yellow(style.Dim("this Program allows the user to customize the text ASCII-art output to different format ")))
	animation.TypeWriterSlow(style.Yellow(style.Dim("it supports colour customization.. Avaliable colours are;")))
	animation.TypeWriterSlow(style.White("‣ RED"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ YELLOW"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ GREEN"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ BLUE"))
	time.Sleep(400 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ CYAN"))
	time.Sleep(400 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ MAGENTA"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ PURPLE"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ DARK RED"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ DARK YELLOW"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ WHITE"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ GREY"))
	fmt.Println()
	animation.TypeWriterSlow(style.Yellow(style.Dim("⦁ it also supports text format Styles customization.. Avaliable Styles are;")))
	animation.TypeWriterSlow(style.White("‣ BOLD"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ DIM"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ ITALIC"))
	time.Sleep(200 * time.Millisecond)
	animation.TypeWriterSlow(style.White("‣ STRIPES"))
	time.Sleep(1600 * time.Millisecond)

	fmt.Println(style.GreyBackground(" ENTER YOUR OPTION TO CONTINUE "))
	fmt.Println()
	fmt.Println(style.Grey("1. Enter key = To start main program"))
	fmt.Println(style.Grey("2. End = To Terminate the program"))
	fmt.Println()
	fmt.Println("\033[1m\033[42m"+" ENTER KEY "+"\033[0m"+"\033[0m", "|====|", "\033[1m\033[41m"+" END "+"\033[0m")

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

start:

	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriterSlow(style.Green(style.Italic("\n▶ NOW LET'S DIVE INTO THE PROGRAM FULLY : " + fullname)))

	fmt.Println(style.Green("___________________________________________________________________________"))
	time.Sleep(600 * time.Millisecond)
	fmt.Println()
	animation.ProgressBarGreen(40)
	time.Sleep(100 * time.Millisecond)

	fmt.Print("\033[3J\033[H\033[2J")

	fmt.Println()
end:
	fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
	(animation.Blink("      \t\t\t "))
	fmt.Println()
	fmt.Print("◉ ")
	fmt.Println(style.GreyBackground(style.Bold("USER : " + fullname + "\n")))
	fmt.Println(style.Green("___________________________________________________________________________"))

	animation.TypeWriter(style.Green(style.Italic("\n▶ Enter your input text...")))

	fmt.Println(style.Green("___________________________________________________________________________"))

input:

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

		inputText = input

		break

	}
font:
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

		if font == "Back" {
			goto input
		}

		switch font {
		case "Standard":
		case "Shadow":
		case "Thinkertoy":
		default:
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: font style not found"))))
			fmt.Println(style.Yellow(style.Upperline(style.Underline("SUPPORTED 🖒 : Standard, Shadow, Thinkertoy, Back(previous input)"))))
			continue
		}

		break
	}

colour:

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

		if colour == "Back" {
			goto font
		}

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
			fmt.Println(style.Yellow(style.Upperline(style.Underline("SUPPORTED 🖒 : Red, Green, Blue, Yellow, Magenta, Cyan, Purple, Dark Red, Dark Yellow, White, Grey, None, Back(previous input)"))))
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

		if Style == "Back" {
			goto colour
		}

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
		fmt.Println()
		animation.LoadingBlue("Processing")
		fmt.Print("\033[3J\033[H\033[2J")
		fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
		fmt.Println("\t\t\t( ｡" + style.Yellow("◕") + "‿‿" + style.Yellow("◕") + "｡)")

		fmt.Println()
		fmt.Println(style.Grey("______________________________________________________________________________"))
		fmt.Print(colour + Style + Output + "\033[0m")
		fmt.Println(style.Grey("______________________________________________________________________________"))
	}

	if font == "Shadow" {

		banner, err := ascii.GetBanner("shadow.txt")
		if err != nil {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
			return

		}

		Output := ascii.GenerateArt(input, banner)
		fmt.Println()
		animation.LoadingBlue("Processing")
		fmt.Print("\033[3J\033[H\033[2J")
		fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
		fmt.Println("\t\t\t( ｡" + style.Yellow("◕") + "‿‿" + style.Yellow("◕") + "｡)")

		fmt.Println()
		fmt.Println(style.Grey("_____________________________________________________________________________________"))
		fmt.Print(colour + Style + Output + "\033[0m")
		fmt.Println(style.Grey("_____________________________________________________________________________________"))
	}

	if font == "Thinkertoy" {

		banner, err := ascii.GetBanner("thinkertoy.txt")
		if err != nil {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: unable to read file"))))
			return

		}

		Output := ascii.GenerateArt(input, banner)
		fmt.Println()
		animation.LoadingBlue("Processing")
		fmt.Print("\033[3J\033[H\033[2J")
		fmt.Println(style.BlueBackground(style.Bold("\t          ● ASCII-ART CLI PROGRAM ●           ")))
		fmt.Println("\t\t\t( ｡" + style.Yellow("◕") + "‿‿" + style.Yellow("◕") + "｡)")

		fmt.Println()
		fmt.Println(style.Grey("___________________________________________________________________________________"))
		fmt.Print(colour + Style + Output + "\033[0m")
		fmt.Println(style.Grey("___________________________________________________________________________________"))

	}

	fmt.Println()
	animation.TypeWriter(style.Yellow(style.Bold("> DO YOU WANT TO CONTINUE ? (YES/NO)")))

	var Continue string

	for {
		fmt.Println()
		fmt.Print(style.Blue(style.Bold("OPTION : ")))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: option can't be empty"))))
			continue
		}

		Continue = Case(input)

		switch Continue {
		case "No":
			fmt.Println()
			fmt.Println(style.Green("__________________________________________________________________________________________________________________________________________"))

			animation.TypeWriter(style.Green(style.Italic("▶ " + fullname + " Thanks for testing this CLI Program, Please do well to follow us on @github.com/janaiegeonu")))

			fmt.Println(style.Green("___________________________________________________________________________________________________________________________________________"))
			fmt.Println()
			animation.BlinkPro("      \t\t\t ")
			fmt.Println()
			return

		case "Yes":
			fmt.Print("\033[3J\033[H\033[2J")
			goto end
		default:
			fmt.Println(style.Red(style.Upperline(style.Underline("ERROR ⚠︎: option not found"))))
			continue
		}

	}
}

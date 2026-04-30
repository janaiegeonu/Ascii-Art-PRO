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
	time.Sleep(2000 * time.Millisecond)
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

		if input 





	}
}

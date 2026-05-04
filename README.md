# Ascii-Art


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

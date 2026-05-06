# PROJECT DOCUMENTATION

This a CLI GO program on ASCII-ART, that allows the user to print out inputed text into cool ASCII-art pattern in three different font style and it also supports unique ASCII customization to suit the taste of the user e.g. colour customization and style format customization.

-----------------------------
### - Project existing files/directory:

|Files/Directory | Category | Function|
|----------------|----------|---------|
| **animation**      | **directory**| **subfolder**|
| `blink.go`       | helper function| it contain animative functions that allows a visual blinking of an ascii art face representation
| `loading.go`| helper function| it contain animative functions that moves dots forward and backward to create a loading experience.
| `progressBar.go`| helper function| it contain animative functions that creates a moving progressBar animation.
|`rainbow.go`| helper function| it contain an animative function that makes text to have a multi-colour display animation.
|`typewriter.go`| helper function| it contain animative functions that brings up or display a motion typewriting animation effect of text.
|                |               |
| **ascii**   | **directory**| **subfolder**|
| `banner.go`| helper function| it contains a function that reads an ASCII banner file data and stores them into a memory.
|`generate.go` | helper function | it contains a function that generate or prints out the ASCII-art pattern of the inputed text.
|`render.go` | helper function | it contains a function that fit each text character ASCII value into a slice row with the length of 8 strings.
|`validate.go`| helper function| it contains a function that validates text character and flags off unrecognized character with no ASCII value in the banner containing only value of 32-126.
|             |                |          |
| **style**   |  **directory**| **subfolder**|
| `styles.go` | helper function | it contains customization functions that style the format of text into different forms e.g Dim, Bold, Underline, Overline and Italic.
|`tint.go` | helper function | it contains different Colour functions to change the colour of text.
|`tintX.go` | helper function | it contains different colour function to change the background colour of text.
|            |              |               |
| go.mod     | Go modules   | this a file for root of Go dependency system, it identifies the program as a module which is collection of GO packages together and keeps track of every external packages this program use.|
| `main.go` |  generic function | this is the heart of the program that collects data from each helper functions and uses it to build itself, it contains the main function and also handles the CLI UI in taking user input data and printing back the output to give great User experience.
| shadow.txt|  text banner file | contains ASCII-art banner data for shadow font.|
|standard.txt| text banner file| contains ASCII-art banner data for standard font.|
|thinkertoy.txt| text banner file | contains ASCII-art banner data for thinkertoy font.

---------------------------------------------------------



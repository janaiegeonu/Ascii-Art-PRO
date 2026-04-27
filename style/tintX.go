package style

func GreyBackground(text string) string {
	red := "\033[40m"
	reset := "\033[0m"
	return red + text + reset
}

func RedBackground(text string) string {
	red := "\033[41m"
	reset := "\033[0m"
	return red + text + reset
}

func GreenBackground(text string) string {
	red := "\033[42m"
	reset := "\033[0m"
	return red + text + reset
}

func YellowBackground(text string) string {
	red := "\033[43m"
	reset := "\033[0m"
	return red + text + reset
}

func BlueBackground(text string) string {
	red := "\033[44m"
	reset := "\033[0m"
	return red + text + reset
}

func MagentaBackground(text string) string {
	red := "\033[45m"
	reset := "\033[0m"
	return red + text + reset
}

func CyanBackground(text string) string {
	red := "\033[46m"
	reset := "\033[0m"
	return red + text + reset
}

func WhiteBackground(text string) string {
	red := "\033[47m"
	reset := "\033[0m"
	return red + text + reset
}

package style

func Bold(text string) string {
	red := "\033[1m"
	reset := "\033[0m"
	return red + text + reset
}

func Dim(text string) string {
	red := "\033[2m"
	reset := "\033[0m"
	return red + text + reset
}

func Italic(text string) string {
	red := "\033[3m"
	reset := "\033[0m"
	return red + text + reset
}

func Underline(text string) string {
	red := "\033[4m"
	reset := "\033[0m"
	return red + text + reset
}

func Upperline(text string) string {
	red := "\033[53m"
	reset := "\033[0m"
	return red + text + reset
}

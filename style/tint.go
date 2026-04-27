package style

func Grey(text string) string {
	red := "\033[30m"
	reset := "\033[0m"
	return red + text + reset
}

func Red(text string) string {
	red := "\033[31m"
	reset := "\033[0m"
	return red + text + reset
}

func Green(text string) string {
	red := "\033[32m"
	reset := "\033[0m"
	return red + text + reset
}

func Yellow(text string) string {
	red := "\033[33m"
	reset := "\033[0m"
	return red + text + reset
}

func Blue(text string) string {
	red := "\033[34m"
	reset := "\033[0m"
	return red + text + reset
}

func Magenta(text string) string {
	red := "\033[35m"
	reset := "\033[0m"
	return red + text + reset
}

func Cyan(text string) string {
	red := "\033[36m"
	reset := "\033[0m"
	return red + text + reset
}

func White(text string) string {
	red := "\033[37m"
	reset := "\033[0m"
	return red + text + reset
}

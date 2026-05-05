package animation

import (
	"Ascii-art-PRO/style"
	"fmt"
	"time"
)

func Loading(text string) string {

	empty := ""
	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {

		fmt.Printf(style.Yellow(style.Bold("\r%s%s"+"\033[0m")), text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
	return empty
}

func Loading1(text string) string {

	empty := ""
	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {

		fmt.Printf(style.Red(style.Bold("\r%s%s")), text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
	return empty
}

func Loading2(text string) string {

	empty := ""
	frame := []string{
		".       ",
		"..      ",
		"...     ",
		"....    ",
		".....   ",
		"......  ",
		"....... ",
		"........",
		"....... ",
		"......  ",
		".....   ",
		"....    ",
		"...     ",
		"..      ",
		".       ",
		"        ",
	}

	for i := 0; i <= 31; i++ {

		fmt.Printf(style.Green(style.Bold("\r%s%s")), text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}

	return empty
}

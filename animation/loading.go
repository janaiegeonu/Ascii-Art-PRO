package animation

import (
	"fmt"
	"time"
)

func Loading(text string) {

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

		fmt.Printf("\r%s%s", text, frame[i%len(frame)])
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

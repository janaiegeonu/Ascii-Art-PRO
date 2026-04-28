package animation

import (
	"Ascii-art-PRO/style"
	"fmt"
	"time"
)

func Blink(text string) {

	frame1 := []string{
		style.Yellow("◕"),
		style.Red(">"),
		style.Yellow("◕"),
	}
	frame2 := []string{
		style.Yellow("◕"),
		style.Red("<"),
		style.Yellow("◕"),
	}
	check := style.Yellow("｡")

	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s( %s%s‿‿%s%s)", text, check,frame1[i%len(frame1)], frame2[i%len(frame2)], check)
		time.Sleep(130 * time.Millisecond)
	}
	fmt.Println()

}

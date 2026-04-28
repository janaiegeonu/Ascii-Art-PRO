package animation

import (
	"fmt"
	"time"
)

func Rainbow(text string) {

	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(20 * time.Millisecond)
	}

	frame := []string{
		"\033[31m",
		"\033[32m",
		"\033[33m",
		"\033[3m\033[33m",
		"\033[2m\033[33m",
		"\033[34m",
		"\033[35m",
		"\033[36m",
		"\033[0m",
	}

	for i := 0; i < 24; i++ {
		fmt.Printf("\r%s%s"+"\033[1m", text, frame[i%len(frame)])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

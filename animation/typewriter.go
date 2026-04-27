package animation

import (
	"fmt"
	"time"
)

func TypeWriter(text string) {

	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

}

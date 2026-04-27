package animation

import (
	"fmt"
	"time"
)

func TypeWriter(text string) {

	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

}

package animation

import (
	"Ascii-art-PRO/style"
	"fmt"
	"strings"
	"time"
)

func ProgressBarRed(total int) {

	for i := 0; i <= total; i++ {
		bar := strings.Repeat("█", i) + strings.Repeat(" ", total-i)
		fmt.Printf(style.Red("\r[%s]%d%%"), bar, 100*i/total)
		time.Sleep(40 * time.Millisecond)

	}
	fmt.Println()

}

func ProgressBarGreen(total int) {

	for i := 0; i <= total; i++ {
		bar := strings.Repeat("█", i) + strings.Repeat(" ", total-i)
		fmt.Printf(style.Green("\r[%s]%d%%"), bar, 100*i/total)
		time.Sleep(45 * time.Millisecond)

	}
	fmt.Println()

}

func ProgressBar(total int) {

	for i := 0; i <= total; i++ {
		bar := strings.Repeat("█", i) + strings.Repeat(" ", total-i)
		fmt.Printf("\r[%s]%d%%", bar, 100*i/total)
		time.Sleep(40 * time.Millisecond)

	}
	fmt.Println()

}

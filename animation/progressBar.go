package animation

import (
	"fmt"
	"strings"
	"time"
)

func ProgressBar(total int) {

	for i := 0; i <= total; i++ {
		bar := strings.Repeat("█", i) + strings.Repeat(" ", total-i)
		fmt.Printf("\r[%s]%d%%", bar, 100*i/total)
		time.Sleep(40 * time.Millisecond)

	}
	fmt.Println()

}

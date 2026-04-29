# Ascii-Art


package main

import (
	"fmt"
	"time"
)

func clear() {fmt.Print("\033[3J\033[H\033[2J")
	
}

func main() {
	fmt.Println("Loading old screen...")
	time.Sleep(2 * time.Second)

	clear()

	fmt.Println("Fresh screen now.")
	time.Sleep(2 * time.Second)

	fmt.Println("Program still running...")
}

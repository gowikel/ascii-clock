package main

import (
	"fmt"
	"main/ascii"
	"time"
)

func main() {
	for {
		now := time.Now()
		f := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
		fmt.Printf("%s\n", ascii.Format(f))
		time.Sleep(time.Second)
		fmt.Printf("\033[11A")
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/afbjorklund/go-termbg/pkg/termbg"
)

func main() {
	const timeout = 100 * time.Millisecond

	fmt.Println("Check terminal background color")

	term, _ := termbg.NewTerminal()
	fmt.Printf("  Term : %s\n", term)

	rgb, err := termbg.NewRGB(timeout)
	if err == nil {
		fmt.Printf("  Color: R=%x, G=%x, B=%x\n", rgb.R, rgb.G, rgb.B)
	} else {
		fmt.Printf("  Color: detection failed %v\n", err)
	}

	theme, err := termbg.NewTheme(timeout)
	if err == nil {
		fmt.Printf("  Theme: %s\n", theme)
	} else {
		fmt.Printf("  Theme: detection failed %v\n", err)
	}
}

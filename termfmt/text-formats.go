package termfmt

import (
	"fmt"
)

// Includes utilities for formatting console output with colors and styles,
//  using ANSI escape codes for text attributes, such as color and underline

const (
	RED       = "\033[31m"
	GREEN     = "\033[32m"
	ORANGE    = "\033[33m"
	BLUE      = "\033[34m"
	UNDERLINE = "\033[4m"
	RESET     = "\033[0m"
)

func PrintfFormat(text string, formats ...string) {
	for _, v := range formats {
		fmt.Print(v)
	}
	fmt.Print(text)
	fmt.Print(RESET)
}
func SprintfFormat(text string, formats ...string) (s string) {
	for _, v := range formats {
		s += v
	}
	s += text + RESET
	return
}

func EraseLine() {
	fmt.Print("\033[2K")
}

func Up(n int) {
	fmt.Printf("\033[%dA", n)
}

func Down(n int) {
	fmt.Printf("\033[%dB", n)
}

func EraseLinesAbove(n int) {
	for range n {
		Up(1)
		EraseLine()
	}
}

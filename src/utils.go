package main

import (
	"fmt"
	"strings"
)

func pass_print(pass string, uncolored bool) {
	if uncolored {
		fmt.Println(pass)
		return
	}

	for _, char := range pass {
		if strings.ContainsRune(NUMS, rune(char)) {
			fmt.Printf("\033[34m%c", char)
		} else if strings.ContainsRune(SYMBOLS, rune(char)) {
			fmt.Printf("\033[31m%c", char)
		} else {
			fmt.Printf("\x1b[0m%c", char)
		}
	}

	fmt.Println("\x1b[0m")
}

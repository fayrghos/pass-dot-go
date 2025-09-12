package funcs

import (
	"fmt"
	"strings"
)

func PassPrint(pass string, colored bool) {
	if !colored {
		fmt.Println(pass)
		return
	}

	for _, char := range pass {
		if strings.ContainsRune(CharsNum, rune(char)) {
			fmt.Printf("\033[34m%c", char)
		} else if strings.ContainsRune(CharsSymbol, rune(char)) {
			fmt.Printf("\033[31m%c", char)
		} else {
			fmt.Printf("\x1b[0m%c", char)
		}
	}

	fmt.Println("\x1b[0m")
}

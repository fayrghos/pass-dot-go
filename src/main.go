package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

const MIN_PASS_LEN = 8
const MAX_PASS_LEN = 80

const ALPHA = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMS = "0123456789"
const SYMBOLS = "@$&!#%&*"

func generate_pass(size int, pass_type int) string {
	output := ""
	valid_chars := ""

	if pass_type >= 0 {
		valid_chars += ALPHA
	}

	if pass_type >= 1 {
		valid_chars += NUMS
	}

	if pass_type >= 2 {
		valid_chars += SYMBOLS
	}

	for range size {
		output += string(valid_chars[rand.IntN(len(valid_chars))])
	}

	return output
}

func main() {
	args, err := parse_args(os.Args[1:])
	if err != nil {
		fmt.Printf("The option '%s' is unknown.\n", err)
		fmt.Println("Use '--help' for more info.")
		return
	}

	if args.show_help {
		fmt.Println("Usage: passgen [size] [...options]")
		fmt.Println("  -l, --letters \t Letters only")
		fmt.Println("  -n, --numbers \t Letters and numbers")
		fmt.Println("  -s, --symbols \t Letters, numbers and symbols")
		fmt.Println("  -u, --uncolored \t Outputs an uncolored password")
		return
	}

	if len(args.others) < 1 {
		fmt.Println("Please insert the password size as an arg.")
		fmt.Println("Use '--help' to see a list of options.")
		return
	}

	pass_len, err := strconv.Atoi(args.others[0])
	if err != nil {
		fmt.Println("The password size must be an integer.")
		return
	}

	if pass_len > MAX_PASS_LEN {
		fmt.Printf("The password size is too big! (>%d)\n", MAX_PASS_LEN)
		return
	} else if pass_len < 8 {
		fmt.Printf("The password size is too small! (<%d)\n", MIN_PASS_LEN)
		return
	}

	pass_print(generate_pass(pass_len, args.password_type), args.no_color)
}

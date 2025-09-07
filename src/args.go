package main

import (
	"errors"
	"strings"
)

const (
	PassAlpha = iota
	PassAlphanum
	PassFull
)

type ArgResult struct {
	show_help     bool
	password_type int
	no_color      bool
	others        []string
}

// -l --letters
// -n --numbers
// -s --symbols
// -h --help
// -u --uncolored

func parse_args(args []string) (ArgResult, error) {
	result := ArgResult{
		password_type: 2,
	}

	for _, value := range args {
		switch value {
		case "-l":
			fallthrough
		case "--letters":
			result.password_type = PassAlpha

		case "-n":
			fallthrough
		case "--numbers":
			result.password_type = PassAlphanum

		case "-s":
			fallthrough
		case "--symbols":
			result.password_type = PassFull

		case "-h":
			fallthrough
		case "--help":
			result.show_help = true

		case "-u":
			fallthrough
		case "--uncolored":
			result.no_color = true

		default:
			if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "--") {
				return ArgResult{}, errors.New(value)
			}
			result.others = append(result.others, value)
		}
	}

	return result, nil
}

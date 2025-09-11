package cli

import (
	"fmt"
	"strconv"

	"src/cmd/funcs"

	"github.com/spf13/cobra"
)

var IsUncolored bool
var IsLetterPass bool
var IsNumPass bool
var IsSymbolPass bool

func rootRun(cmd *cobra.Command, args []string) {
	pass_len, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The password size is not an integer. Aborting...")
		return
	}

	if pass_len > funcs.MaxPassLen {
		fmt.Printf("The password size is too big! (>%d chars)\n",
			funcs.MaxPassLen)
		return
	} else if pass_len < funcs.MinPassLen {
		fmt.Printf("The password size is too small! (<%d chars)\n",
			funcs.MinPassLen)
		return
	}

	pass_type := funcs.PassSymbol
	switch true {
	case IsSymbolPass:
		pass_type = funcs.PassSymbol

	case IsNumPass:
		pass_type = funcs.PassNum

	case IsLetterPass:
		pass_type = funcs.PassLetter
	}

	pass := funcs.GeneratePass(pass_len, pass_type)
	funcs.PassPrint(pass, IsUncolored)
}

var rootCommand = &cobra.Command{
	Use:   "passdotgo",
	Short: "Yet another password generator utility.",
	Args:  cobra.ExactArgs(1),
	Run:   rootRun,
}

func Setup() {
	rootCommand.Flags().BoolVarP(&IsUncolored, "uncolored", "u", false,
		"outputs an uncolored password")

	rootCommand.Flags().BoolVarP(&IsLetterPass, "letters", "l", false,
		"generates a letters only password")

	rootCommand.Flags().BoolVarP(&IsNumPass, "numbers", "n", false,
		"generates a letters and numbers password")

	rootCommand.Flags().BoolVarP(&IsSymbolPass, "symbols", "s", false,
		"generates an all characters password")

	rootCommand.Execute()
}

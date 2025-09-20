package cli

import (
	"fmt"
	"strconv"

	"src/internal/funcs"

	"github.com/spf13/cobra"
)

var (
	eyeCandyEnabled bool
	isLetterPass    bool
	isNumPass       bool
	isSymbolPass    bool
	passCount       int
)

func rootRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("No password size was specified. Use '--help' for more info.")
		return
	}

	if passCount <= 0 || passCount > 99 {
		fmt.Println("Invalid password amount. Aborting...")
		return
	}

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
	case isSymbolPass:
		pass_type = funcs.PassSymbol

	case isNumPass:
		pass_type = funcs.PassNum

	case isLetterPass:
		pass_type = funcs.PassLetter
	}

	for i := range passCount {
		pass := funcs.GeneratePass(pass_len, pass_type)

		if passCount == 1 {
			funcs.PassPrint(pass, eyeCandyEnabled)
		} else {
			funcs.PassPrint(fmt.Sprintf("%.02d | %s", i+1, pass), eyeCandyEnabled)
		}
	}
}

var rootCommand = &cobra.Command{
	Use:          "passdotgo <length>",
	Short:        "Yet another password generator utility.",
	Args:         cobra.MaximumNArgs(1),
	Run:          rootRun,
	SilenceUsage: true,
}

func Setup() {
	rootCommand.Flags().BoolVarP(&eyeCandyEnabled, "eye-candy", "e", false,
		"outputs a colorful password")

	rootCommand.Flags().BoolVarP(&isLetterPass, "letters", "l", false,
		"generates a letters only password")

	rootCommand.Flags().BoolVarP(&isNumPass, "numbers", "n", false,
		"generates a letters and numbers password")

	rootCommand.Flags().BoolVarP(&isSymbolPass, "symbols", "s", false,
		"generates an all characters password")

	rootCommand.Flags().IntVarP(&passCount, "count", "c", 1,
		"the amount of passwords to generate")

	rootCommand.Execute()
}

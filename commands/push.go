package commands

import (
	"fmt"

	// custom imports

	// infered imports

	"github.com/hofstadter-io/hof/lib/sync"
	"github.com/spf13/cobra"
)

// Tool:   hof
// Name:   push
// Usage:  push
// Parent: hof

var PushLong = `Uploads the local copy and makes it the latest copy in Studios`

var PushCmd = &cobra.Command{

	Use: "push",

	Short: "Send and make the latest version on Studios",

	Long: PushLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In pushCmd", "args", args)
		// Argument Parsing

		fmt.Println("hof push:")
		sync.Push()
	},
}

func init() {
	RootCmd.AddCommand(PushCmd)
}

func init() {
	// add sub-commands to this command when present

}

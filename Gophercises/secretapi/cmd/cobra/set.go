package cobra

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(Cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}

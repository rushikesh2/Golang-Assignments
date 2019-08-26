package cobra

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run: func(Cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

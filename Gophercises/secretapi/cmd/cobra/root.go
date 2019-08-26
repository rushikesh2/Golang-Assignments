package cobra

import (
	"github.com/spf13/cobra"
)

//RootCmd root key
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API and secrets manager",
}

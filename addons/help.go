package addons

import "github.com/spf13/cobra"

// Help is a standard Help command
func Help(cmd *cobra.Command, args []string) {
	cmd.Help()
}

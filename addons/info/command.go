package info

import (
	"github.com/spf13/cobra"
)

// InjectInfo injects itself and all subcommands into the supplied command
func InjectInto(cmd *cobra.Command) {
	infoCommand := &cobra.Command{
		Use:   "info",
		Short: "Lists api.json contents",
		Long:  ``,
		RunE:  parseConfigJSON,
	}
	cmd.AddCommand(routeCommand())
	cmd.AddCommand(infoCommand)
}

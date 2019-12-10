package labels

import (
	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/addons"
)

// InjectInfo injects itself and all subcommands into the supplied command
func InjectInto(cmd *cobra.Command) {
	labelsCommand := &cobra.Command{
		Use:   "labels",
		Short: "Find all labels in current project",
		Long:  ``,
		Run:   addons.Help,
	}
	labelsCommand.AddCommand(showCommand)
	labelsCommand.AddCommand(uploadCommand)
	cmd.AddCommand(labelsCommand)

}

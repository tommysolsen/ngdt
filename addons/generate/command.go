package generate

import (
	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/addons"
)

// InjectInfo injects itself and all subcommands into the supplied command
func InjectInto(cmd *cobra.Command) {
	genCommand := &cobra.Command{
		Use:   "gen",
		Short: "Generate Netflex Objects",
		Long:  ``,
		Run:   addons.Help,
	}
	genCommand.AddCommand(GenerateTemplateCommand())
	genCommand.AddCommand(GenerateComponentCommand())
	cmd.AddCommand(genCommand)
}

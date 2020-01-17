package completion

import (
	"os"

	"github.com/spf13/cobra"
)

func InjectInto(cmd *cobra.Command) {
	_completion := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

	. <(bitbucket completion)

	To configure your bash shell to load completions for each session add to your bashrc

	# ~/.bashrc or ~/.profile
	. <(bitbucket completion)
	`,
		RunE: func(lcmd *cobra.Command, args []string) error {
			cmd.GenBashCompletion(os.Stdout)
			return nil
		},
	}

	cmd.AddCommand(_completion)
}

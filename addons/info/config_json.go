package info

import (
	"os"

	"github.com/tommysolsen/ngdt/assets"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/lib/filetools"
)

func parseConfigJSON(cmd *cobra.Command, args []string) error {
	apiData, err := filetools.GetConfigJSON()
	if err != nil {
		return err
	}
	t, err := assets.TextTemplate("ConfigJsonOutput")
	if err != nil {
		return err
	}
	t.Execute(os.Stdout, apiData)
	return nil
}

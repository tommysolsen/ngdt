package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/assets"
	"github.com/tommysolsen/ngdt/lib/filetools"
	"github.com/tommysolsen/ngdt/lib/netflex"
)

func GenerateComponentCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "component",
		Short: "Generate a Netflex Component",
		Long:  "Generates and registeres a new netflex component",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("You need to supply exactly two args, filename and component name")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			APIData, err := filetools.GetConfigJSON()
			if err != nil {
				return err
			}
			client := netflex.New(*APIData)

			alias := args[0]
			name := strings.Join(args[1:len(args)], " ")
			template := netflex.Template{
				Alias: alias,
				Name:  name,
				Type:  "builder",
			}

			returnVal := make(map[string]interface{})
			err = client.Post("foundation/templates", template, &returnVal)
			if err != nil {
				return err
			}

			tmp, err := assets.PHPTemplate("component")
			if err != nil {
				return err
			}

			splittedAlias := strings.Split(alias, "/")

			if len(splittedAlias) > 1 {
				err = os.MkdirAll(
					"components/"+strings.Join(splittedAlias[0:len(splittedAlias)-1], "/"),
					os.ModePerm,
				)
				if err != nil {
					return err
				}
			}

			fp, err := os.OpenFile("components/"+alias+".php", os.O_WRONLY, 755)
			if err != nil {
				if err.Error() == "open components/"+alias+".php: no such file or directory" {
					fp, err = os.Create("components/" + alias + ".php")
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			err = tmp.Execute(fp, template)
			if err != nil {
				return err
			}

			err = fp.Close()
			if err != nil {
				return err
			}

			return nil
		},
	}
}

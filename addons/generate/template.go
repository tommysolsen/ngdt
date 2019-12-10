package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/assets"
	"github.com/tommysolsen/ngdt/helpers"
	"github.com/tommysolsen/ngdt/lib/filetools"
	"github.com/tommysolsen/ngdt/lib/netflex"
)

func GenerateTemplateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "template",
		Short: "Generate a Netflex Template",
		Long:  "Generates and registeres a new netflex template",
		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				return fmt.Errorf("You must atleast supply a name for your template")
			}
			if len(args) > 1 {
				nameSegments := args[1:len(args)]
				for _, segment := range nameSegments {
					str := strings.Split(segment, ":")
					fmt.Println(str, len(str))
					if len(str) < 2 {
						return fmt.Errorf(
							"The name segments has to be in a `name:type` format. \r\nValid types are %s",
							strings.Join(netflex.TemplateFiledNames, ", "),
						)
					}

					if !helpers.InArray(str[1], netflex.TemplateFiledNames) {
						return fmt.Errorf(
							"You are trying to make a field of an invalid type %s. \r\nValid types are %s",
							str[1],
							strings.Join(netflex.TemplateFiledNames, ", "),
						)
					}
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			APIData, err := filetools.GetConfigJSON()
			if err != nil {
				return err
			}
			client := netflex.New(*APIData)

			name := args[0]

			template := netflex.Template{
				Alias: name,
				Name:  name,
				Type:  "page",
			}

			returnVal := make(map[string]interface{})
			err = client.Post("foundation/templates", template, &returnVal)
			if err != nil {
				return err
			}

			tmp, err := assets.PHPTemplate("template")
			if err != nil {
				return err
			}

			fp, err := os.OpenFile("templates/"+name+".php", os.O_WRONLY, 755)
			if err != nil {
				if err.Error() == "open templates/"+name+".php: no such file or directory" {
					fp, err = os.Create("templates/" + name + ".php")
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

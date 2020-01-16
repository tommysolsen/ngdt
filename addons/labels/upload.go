package labels

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/lib/filetools"
	"github.com/tommysolsen/ngdt/lib/netflex"
)

var uploadCommand = &cobra.Command{
	Use:   "upload",
	Short: "Upload all labels in current project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := filetools.GetRootPath()
		if err != nil {
			return err
		}

		uniqueLabels := make([]string, 0)
		files := filetools.FindFilesOfType("php", path)
		for _, file := range files {
			labels, err := filetools.FindLabels(file)
			if err != nil {
				return err
			}

		LABEL_LOOP:
			for _, label := range *labels {
				for _, ul := range uniqueLabels {
					if ul == label {
						continue LABEL_LOOP
					}
				}
				uniqueLabels = append(uniqueLabels, label)
			}
		}
		APIData, err := filetools.GetConfigJSON()
		if err != nil {
			return err
		}
		nfClient := netflex.New(*APIData)
		fmt.Printf("Found %d labels for this site: \r\n", len(uniqueLabels))
		for _, v := range uniqueLabels {
			existed, err := nfClient.PostLabel(v)
			if err != nil {
				fmt.Printf("ERROR: %s", err.Error())
			} else {
				if existed {
					fmt.Printf("EXIST -> %s\r\n", v)
				} else {
					fmt.Printf("NEW   -> %s\r\n", v)
				}
			}
		}
		return nil
	},
}

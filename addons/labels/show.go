package labels

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/lib/filetools"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "Find all labels in current project",
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
			x := strconv.Itoa(len(*labels))
			fmt.Println(file + "(" + x + ")")

		LABEL_LOOP:
			for _, label := range *labels {
				fmt.Println("\t" + label)

				for _, ul := range uniqueLabels {
					if ul == label {
						continue LABEL_LOOP
					}
				}
				uniqueLabels = append(uniqueLabels, label)
			}
		}
		fmt.Printf("Found %d labels for this site: \r\n", len(uniqueLabels))
		for _, v := range uniqueLabels {
			fmt.Printf(" * %s\r\n", v)
		}
		return nil
	},
}

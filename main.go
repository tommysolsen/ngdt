package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/addons/generate"
	"github.com/tommysolsen/ngdt/addons/info"
	"github.com/tommysolsen/ngdt/addons/labels"
)

// Generate assets
//go:generate go-bindata -o assets/templates_temp.go -pkg assets static/
func main() {
	var rootCmd = &cobra.Command{
		Use:   "nf",
		Short: "Netflex Golang Development Tool",
		Long:  `Basic tool for Netflex site development`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	info.InjectInto(rootCmd)
	generate.InjectInto(rootCmd)
	labels.InjectInto(rootCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error Occurred: %s\r\n", err.Error())
	}
}

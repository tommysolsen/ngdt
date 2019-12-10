package info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tommysolsen/ngdt/helpers"
	"github.com/tommysolsen/ngdt/lib/filetools"
	"github.com/tommysolsen/ngdt/lib/netflex"
)

var ShowAll bool

func routeCommand() *cobra.Command {
	routesCommand := &cobra.Command{
		Use:   "routes",
		Short: "Lists all routes for the site",
		Long:  `Queries both netflex and routes.json and lists routes`,
		RunE:  getRoutes,
	}

	routesCommand.Flags().BoolVarP(&ShowAll, "all", "a", false, "Show all rows")
	return routesCommand
}

func getRoutes(cmd *cobra.Command, args []string) error {
	showAll := ShowAll
	data, err := filetools.GetConfigJSON()
	if err != nil {
		return err
	}
	client := netflex.New(*data)
	client.LoadTemplates()

	var pages []netflex.Page
	client.Get("builder/pages", &pages)
	table := tablewriter.NewWriter(os.Stdout)

	sort.Slice(pages, func(j, i int) bool {
		return strings.Compare(pages[i].URL, pages[j].URL) > 0
	})

	headers := []string{"ID", "Published", "URL"}
	if showAll || helpers.InArray("language", args) {
		headers = append(headers, "Language")
	}
	if showAll || helpers.InArray("description", args) {
		headers = append(headers, "Description")

	}

	table.SetHeader(headers)
	for _, page := range pages {
		published := "No"
		if page.Published {
			published = "Yes"
		}

		templateID := int64(page.Template)
		templateString := ""
		if template, err := client.GetTemplateById(templateID); err == nil && template != nil {
			templateString = "templates/" + template.Alias + ".php"
		} else {
			templateString = strconv.FormatInt(templateID, 10)
			if templateID == 0 {
				templateString = ""
			}
		}

		row := []string{
			strconv.FormatInt(int64(page.ID), 10),
			published,
			page.URL,
			templateString,
		}
		if showAll || helpers.InArray("language", args) {
			row = append(row, page.Language)
		}
		if showAll || helpers.InArray("description", args) {
			row = append(row, page.Description)
		}
		if showAll || helpers.InArray("public", args) {
			public := "No"
			if page.Public {
				public = "Yes"
			}
			row = append(row, public)
		}
		table.Append(row)
	}

	routesParsed := true
	if url, err := filetools.FindInProject("routes.json", ".", "config"); err == nil {
		var data map[string]string
		file, err := ioutil.ReadFile(url)
		if len(file) > 0 {
			if err != nil {
				return err
			}
			err = json.Unmarshal(file, &data)
			if err != nil {
				routesParsed = false
			}
			if routesParsed {
				for key, value := range data {
					row := []string{
						"",
						"Route",
						key,
						value,
					}
					if showAll || helpers.InArray("language", args) {
						row = append(row, "")
					}
					if showAll || helpers.InArray("description", args) {
						row = append(row, "")
					}
					if showAll || helpers.InArray("public", args) {
						row = append(row, "Yes")
					}
					table.Append(row)
				}
			}

		}
	}
	table.Render()
	if routesParsed == false {
		fmt.Println("Could not parse routes.json\r\n ")
	}

	return nil
}

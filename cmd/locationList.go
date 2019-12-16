/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"strconv"

	loc "github.com/septianw/jas-location"
	"github.com/spf13/cobra"
)

// locationCmd represents the location command
var locationListCmd = &cobra.Command{
	Use:   "location",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("location called")
		var locFind loc.LocationIn

		if _, err := os.Stat("/tmp/shinyRuntimeFile"); os.IsNotExist(err) {
			fmt.Println("This application must run at the same machine with server, and make sure the server is running.")
			os.Exit(1)
			// path/to/whatever does not exist
		}

		var limit int = 10
		var offset int = 0
		var err error
		limitFlag := cmd.Flag("limit")
		offsetFlag := cmd.Flag("offset")

		if limitFlag.Changed {
			limit, err = strconv.Atoi(limitFlag.Value.String())
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
		}
		if offsetFlag.Changed {
			offset, err = strconv.Atoi(offsetFlag.Value.String())
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			// for _, l := range locations {
			// 	fmt.Printf("\n%+v\n", location)
			// }
		}

		find := cmd.Flag("find")
		if find.Changed {
			locFind.Name = find.Value.String()
			locs, err := loc.FindLocation(locFind)
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			for _, l := range locs {
				fmt.Printf("\n%+v\n", l)
			}
		} else {
			locations, err := loc.GetLocation(-1, int64(limit), int64(offset))
			if err != nil {
				fmt.Println(err)
			}
			for _, l := range locations {
				fmt.Printf("\n%+v\n", l)
			}
		}
	},
}

func init() {
	listCmd.AddCommand(locationListCmd)

	locationListCmd.Flags().String("find", "", "Find location by name.")
	locationListCmd.Flags().Int("limit", 10, "Find location by name.")
	locationListCmd.Flags().Int("offset", 0, "Find location by name.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// locationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// locationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	"strings"

	loc "github.com/septianw/jas-location"
	"github.com/spf13/cobra"
)

func formatFloat(input string, decimal int) float64 {
	var out []string
	// input = strings.TrimSpace(input)
	_, err := strconv.ParseFloat(input, 8)
	if err != nil {
		fmt.Printf("\n\nErr: %+v\n\n", err.Error())
		return 0.0
	}

	segments := strings.Split(input, ".")
	segmentByte := []byte(segments[1])
	out = append(out, segments[0])
	//fmt.Println()
	out = append(out, string(segmentByte[:decimal]))
	f, err := strconv.ParseFloat(strings.Join(out, "."), 8)
	if err != nil {
		fmt.Printf("\n\nErr: %+v\n\n", err.Error())
		return 0.0
	}

	return f
}

// locationCmd represents the location command
var locationRegCmd = &cobra.Command{
	Use:   "location",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("location called")
		var location loc.LocationIn

		if cmd.Flag("lat").Changed &&
			cmd.Flag("lon").Changed &&
			cmd.Flag("name").Changed {
			fmt.Println(cmd.Flag("lon").Value.Type())
			lon := formatFloat(cmd.Flag("lon").Value.String(), 8)
			lat := formatFloat(cmd.Flag("lat").Value.String(), 8)

			if (lat == 0.0) || (lon == 0.0) {
				fmt.Printf("\nLon: %+v\nLat: %+v\nOrigin lon: %+v\nOrigin lat: %+v\n\n",
					lon, lat, cmd.Flag("lon").Value.String(), cmd.Flag("lat").Value.String())
				fmt.Println("Latitude and or Longitude cannot be zero.")
				os.Exit(2)
			}
			location.Name = cmd.Flag("name").Value.String()
			location.Latitude = lat
			location.Longitude = lon

			locOut, err := loc.InsertLocation(location)
			if err != nil {
				fmt.Printf("Error: %+v", err)
				os.Exit(1)
			}
			fmt.Printf("\n\n%+v\n\n", locOut)
		} else {
			cmd.Help()
		}
		// loc.InsertLocation()
	},
}

func init() {
	registerCmd.AddCommand(locationRegCmd)

	locationRegCmd.Flags().Float64("lat", 0.0, "--lat -110.343523")
	locationRegCmd.Flags().Float64("lon", 0.0, "--lon -9.343523")
	locationRegCmd.Flags().String("name", "", `--name jogja47`)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// locationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// locationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

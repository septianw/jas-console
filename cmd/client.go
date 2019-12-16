/*
Copyright © 2019 Septian Wibisono

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
	"strings"

	term "github.com/septianw/jas-terminal"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Register new client application",
	Long:  `Register new client application to get clientid and client secret.`,
	Run: func(cmd *cobra.Command, args []string) {
		var clientname string
		name := cmd.Flag("name")

		if _, err := os.Stat("/tmp/shinyRuntimeFile"); os.IsNotExist(err) {
			fmt.Println("This application must run at the same machine with server, and make sure the server is running.")
			os.Exit(1)
			// path/to/whatever does not exist
		}

		if strings.Compare(name.Value.String(), "") != 0 {
			clientname = name.Value.String()
		}

		if strings.Compare(clientname, "") == 0 {
			cmd.Help()
			os.Exit(0)
		}

		cCred, err := term.InsertClientCredential(clientname)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("\n\nClient name   : %s\n", cCred.ClientName)
		fmt.Printf("Client id     : %s\n", cCred.ClientId)
		fmt.Printf("Client secret : %s\n\n\n", cCred.ClientSecret)

		// fmt.Println("client called")
	},
}

func init() {
	registerCmd.AddCommand(clientCmd)

	clientCmd.Flags().String("name", "", "Client name yang akan didaftarkan.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

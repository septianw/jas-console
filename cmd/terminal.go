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

	"strings"

	"os"

	term "github.com/septianw/jas-terminal"

	// "github.com/septianw/jas/common"
	"github.com/spf13/cobra"
)

// terminalCmd represents the terminal command
var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "Mendaftarkan terminal baru pada sistem",
	Long:  `Sistem pendaftaran terminal untuk pertama kali, hanya bisa dilakukan pada mesin yang sama dengan server data.`,
	Run: func(cmd *cobra.Command, args []string) {
		var termin term.TerminalIn

		terminalid := cmd.Flag("terminalid")
		name := cmd.Flag("name")
		location := cmd.Flag("location")
		// rt := common.ReadRuntime()
		// dbo := common.LoadDatabase(rt.Libloc, rt.Dbconf)
		// db, err := dbo.OpenDb(rt.Dbconf)

		if _, err := os.Stat("/tmp/shinyRuntimeFile"); os.IsNotExist(err) {
			fmt.Println("This application must run at the same machine with server, and make sure the server is running.")
			os.Exit(1)
			// path/to/whatever does not exist
		}

		if strings.Compare(terminalid.Value.String(), "") != 0 {
			termin.TerminalId = terminalid.Value.String()
		}

		if strings.Compare(terminalid.Value.String(), "") != 0 {
			termin.Name = name.Value.String()
		}

		if strings.Compare(terminalid.Value.String(), "") != 0 {
			termin.Location = location.Value.String()
		}

		if (term.TerminalIn{}) == termin {
			cmd.Help()
			os.Exit(2)
		}

		termout, err := term.InsertTerminal(termin)
		fmt.Printf("%+v %+v", termout, err)
	},
}

func init() {
	registerCmd.AddCommand(terminalCmd)

	terminalCmd.Flags().String("terminalid", "", "Terminal id yang akan didaftarkan.")
	terminalCmd.Flags().String("name", "", "Nama terminal yang akan didaftarkan.")
	terminalCmd.Flags().String("location", "", "Lokasi terminal yang akan didaftarkan.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// terminalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// terminalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

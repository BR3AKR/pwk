// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/BR3AKR/pwk/credmgr"
	"github.com/spf13/cobra"
)

var show bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List your currently saved passwords",
	Long: `
List displays all of your currently saved passwords in a table.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !pwkExists() {
			fmt.Println("No .pwk could be found. Please execute the add command to initialize")
			return
		}
		password, _ = promptIfEmpty(password, "Password: ", true)

		creds, _ := credmgr.DeserializeData(".pwk", password)

		var display string

		if show {
			display = "Id: %s, Loc: %s, User: %s, Pass: %s, Notes: %s\n"
		} else {
			display = "Id: %s, Loc: %s, User: %s, Pass: ********, Notes: %s\n"
		}
		for _, cred := range creds {
			if show {
				fmt.Printf(display, cred.Id, cred.Location, cred.User, cred.Password, cred.Notes)
			} else {
				fmt.Printf(display, cred.Id, cred.Location, cred.User, cred.Notes)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&show, "show", "s", false, "Display passwords in the results of list")
}

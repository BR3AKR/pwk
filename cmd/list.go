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
	table "github.com/BR3AKR/cli-table"
	"github.com/BR3AKR/pwk/credmgr"
	"github.com/spf13/cobra"
)

var show bool
var hideHeaders bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List your currently saved passwords",
	Long: `
List displays all of your currently saved passwords in a table.`,
	Run: func(cmd *cobra.Command, args []string) {
		if hideHeaders {
			printCreds(creds)
		} else {
			printCredsWithHeader(creds)
		}
	},
}

func printCreds(creds []credmgr.Credential) {
	data := make([][]string, len(creds))

	for i, cred := range creds {
		data[i] = *makeRow(cred.Id, cred.Location, cred.User, cred.Password, cred.Notes, show)
	}

	table.Print(&data, false)
}

func printCredsWithHeader(creds []credmgr.Credential) {
	data := make([][]string, len(creds)+1)

	data[0] = *makeRow("Id", "Location", "User", "Password", "Notes", true)

	for i, cred := range creds {
		data[i+1] = *makeRow(cred.Id, cred.Location, cred.User, cred.Password, cred.Notes, show)
	}

	table.Print(&data, true)
}

func makeRow(id, location, user, password, notes string, showPassword bool) *[]string {
	row := make([]string, 5)
	row[0] = id
	row[1] = location
	row[2] = user
	if showPassword {
		row[3] = password
	} else {
		row[3] = "********"
	}
	row[4] = notes
	return &row
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&show, "show", "s", false, "Display passwords in the results of list")
	listCmd.Flags().BoolVarP(&hideHeaders, "hideHeaders", "i", false, "Hide the headers")
}

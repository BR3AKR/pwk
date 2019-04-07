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
	"github.com/BR3AKR/pwk/credmgr"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a password",
	Long: `
Add creates a new enty in your list of passwords`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ = promptIfEmpty(id, "Id: ", false)
		location, _ = promptIfEmpty(location, "Location: ", false)
		user, _ = promptIfEmpty(user, "User: ", false)
		// TODO generate a random password if none is supplied
		locPassword, _ = promptIfEmpty(locPassword, "Location Password: ", true)
		notes, _ = promptIfEmpty(notes, "Notes: ", false)

		cred := credmgr.Credential{Id: id, Location: location, User: user, Password: locPassword, Notes: notes}
		creds = append(creds, cred)
		credmgr.SerializeData(creds, pwfile, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&location, "location", "l", "", "The location which the credentials belong to")
	addCmd.Flags().StringVarP(&user, "user", "u", "", "The user portion of the credential (username, email, etc.)")
	addCmd.Flags().StringVarP(&locPassword, "password", "p", "", "The password for this credential")
	addCmd.Flags().StringVarP(&notes, "notes", "n", "", "Any notes concerning these credentials")
	addCmd.Flags().StringVarP(&id, "id", "i", "", "A unique identifier for this location. Something short and memorable is best")
}

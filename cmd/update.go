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

var updateCmd = &cobra.Command{
	Use:   "update [id of credentials]",
	Short: "Update a credential entry",
	Long: `
Updates an entry out of your password file by id.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := range creds {
			if creds[i].Id == args[0] {
				if cmd.Flag("location").Changed {
					creds[i].Location = location
				}
				if cmd.Flag("user").Changed {
					creds[i].User = user
				}
				if cmd.Flag("password").Changed {
					creds[i].Password = locPassword
				}
				if cmd.Flag("notes").Changed {
					creds[i].Notes = notes
				}
				if cmd.Flag("id").Changed {
					creds[i].Id = id
				}
				break
			}
		}
		credmgr.SerializeData(creds, pwfile, password)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&location, "location", "l", "", "The location which the credentials belong to")
	updateCmd.Flags().StringVarP(&user, "user", "u", "", "The user portion of the credential (username, email, etc.)")
	updateCmd.Flags().StringVarP(&locPassword, "password", "p", "", "The password for this credential")
	updateCmd.Flags().StringVarP(&notes, "notes", "n", "", "Any notes concerning these credentials")
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "A unique identifier for this location. Something short and memorable is best")
}

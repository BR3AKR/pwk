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
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [id of credentials]",
	Short: "Get a password and put it in the clipboard",
	Long: `
Searches through your current list of passwords by id, if
one is found, places that password in your clipboard. Be
sure to clear your clipboard after using this command.`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var pass string
		for _, cred := range creds {
			if cred.Id == args[0] {
				pass = cred.Password
				break
			}
			// TODO Implement fuzzy checking
		}
		clipboard.WriteAll(pass)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

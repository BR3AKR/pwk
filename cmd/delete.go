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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id of credentials]",
	Short: "Delete a credential entry",
	Long: `
Deletes an entry out of your password file by id. Be careful
once it's gone, there is no way to recover it.`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		for i, cred := range creds {
			if cred.Id == args[0] {
				creds = append(creds[:i], creds[i+1:]...)
				break
			}
		}
		credmgr.SerializeData(creds, pwfile, password)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

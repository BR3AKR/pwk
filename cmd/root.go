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
	"log"
	"os"
	osuser "os/user"

	"github.com/BR3AKR/pwk/credmgr"
	"github.com/spf13/cobra"
)

var creds []credmgr.Credential
var location string
var id string
var user string
var password string
var locPassword string
var notes string
var pwfile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwk",
	Short: "A command line password manager",
	Long: `Password Keeper (pwk) is a CLI password manager
written in Go with the intention of being fast,
and simple to use. Passwords are stored in a
password encrypted file using modern hashing and
a modern encryption algorithm.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var err error
	password, err = prompt("Password: ", true)
	if err != nil {
		log.Fatalf("error prompting for password: %v", err)
	}

	pwfile, err = getPwkFile()
	pwfExists := !os.IsNotExist(err)

	if err != nil && pwfExists {
		log.Fatalf("error finding .pwk: %v", err)
	}

	if pwfExists {
		creds, err = credmgr.DeserializeData(pwfile, password)
		if err != nil {
			log.Fatalf("error deserializing password file: %v", err)
		}
	}
}

func getPwkFile() (string, error) {
	usr, err := osuser.Current()
	if err != nil {
		return "", err
	}

	file := usr.HomeDir + "/.pwk"
	_, err = os.Stat(file)

	return file, err
}

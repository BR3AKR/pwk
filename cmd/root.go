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
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgFile string
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
	pwfile, err = getConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() (string, error) {
	usr, err := osuser.Current()
	if err != nil {
		return "", err
	}
	var file = usr.HomeDir + "/.pwk"
	exists, err := fileExists(file)
	if err != nil {
		return "", err
	} else if exists {
		return file, nil
	}

	file, err = os.Executable()
	if err != nil {
		return "", err
	}

	file = filepath.Dir(file)
	file += "/.pwk"
	log.Printf("Checking %s", file)

	exists, err = fileExists(file)
	if err != nil {
		return "", err
	} else if exists {
		return file, nil
	}
	log.Printf("Not found")

	file = usr.HomeDir + "/.pwk"
	return file, nil
}

func fileExists(file string) (bool, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

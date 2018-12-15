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

package main

import "github.com/BR3AKR/pwk/cmd"

// TODO Read effective go, take advice
// TODO Add Unit Tests
// TODO Add go doc
// TODO Implement command line library
// TODO Implement operations
//      Create
//      Read
//      Update
//      Delete
// TODO Implement GET to clipboard
// TODO Implement INIT - creates password encoded file, creates unique SALT
// TODO Learn how to add to package manager (pacman -S pwkeeper)
// TODO Learn how to create windows installer
func main() {
	cmd.Execute()
}

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func promptIfEmpty(value, prompt string, isPassword bool) (string, error) {
	if value == "" {
		var err error

		if isPassword {
			var pass []byte
			fmt.Print(prompt)
			pass, err = terminal.ReadPassword(int(syscall.Stdin))
			value = string(pass)
			fmt.Println("")
		} else {
			fmt.Print(prompt)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			value = scanner.Text()
		}
		return value, err
	}

	return value, nil
}

func pwkExists() bool {
	_, err := os.Stat(pwfile)
	return !os.IsNotExist(err)
}

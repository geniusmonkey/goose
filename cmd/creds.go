package cmd

import (
	"fmt"
	"github.com/geniusmonkey/gander/creds"
	"github.com/geniusmonkey/gander/env"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
)

var credsCmd = &cobra.Command{
	Use: "creds",
}

var credsAddCmd = &cobra.Command{
	Use:   "add [env name] [username]",
	Short: "add environment password for this project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		envName := args[0]
		username := args[1]

		e, err := env.Get(proj, envName)
		if err != nil {
			log.Fatal("failed to get environment")
		}

		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}
		password := string(bytePassword)

		c := creds.Credentials{Username: username, Password: password}
		creds.Save(*proj, e, c)
	},
}

func init() {
	rootCmd.AddCommand(credsCmd)

	credsCmd.AddCommand(credsAddCmd)
}
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ishitb/mossgo/utils"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Adding your MOSS unique id",
	Long:  `This command asks for you MOSS unique id and adds it to the creds.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		login(args)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login(args []string) {
	var uniqueId = ""

	if len(args) < 1 {
		reader := bufio.NewReader(os.Stdin)
		uniqueId = utils.GetInput("Enter your unique MOSS ID", reader)
	} else {
		uniqueId = args[0]
	}

	uniqueIdJson := fmt.Sprintf("{\"uniqueId\": %v}", uniqueId)
	utils.SaveFile("creds.json", []byte(uniqueIdJson))
}

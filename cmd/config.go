package cmd

import (
	"fmt"
	"os/exec"

	"github.com/solelymoose/golangprojectmanager/shared"
	"github.com/spf13/cobra"
)

var config = &cobra.Command{
	Use:   "config",
	Short: "Opens the config folder",
	Run: func(cmd *cobra.Command, args []string) {
		// Assuming shared.ConfigFile holds the full path to the config file
		command := exec.Command("code", shared.ConfigFile) // Pass the path as an argument to the command

		err := command.Run()
		if err != nil {
			fmt.Println("Error opening config folder:", err)
		}
	},
}

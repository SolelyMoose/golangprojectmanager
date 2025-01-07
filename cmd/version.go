package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/solelymoose/golangprojectmanager/shared"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Displays the GolangProjectManager version",
	Run: func(cmd *cobra.Command, args []string) {
		// Use the shared.ConfigFile variable
		data, err := os.ReadFile(shared.ConfigFile)
		if err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
			return
		}

		var jsonData map[string]interface{}

		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}

		version, ok := jsonData["version"].(string)
		if !ok {
			fmt.Println("Version not found or invalid format")
			return
		}

		fmt.Printf("GolangProjectManager version is: %v\n", version)
	},
}

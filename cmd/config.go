package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/solelymoose/golangfilemanager/shared" // Adjust the import path as necessary

	"github.com/spf13/cobra"
)

var config = &cobra.Command{
	Use:   "config",
	Short: "Opens the config",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := filepath.Join(shared.MainDir, "config.json")
		fmt.Println(shared.MainDir)
		fmt.Println(filePath)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close() // Ensure the file is closed when done

		// Read the file into memory
		var configData map[string]interface{}
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&configData); err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			return
		}

		// Print the parsed JSON
		fmt.Printf("Config Data: %+v\n", configData)
	},
}

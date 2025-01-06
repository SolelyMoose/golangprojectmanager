package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/solelymoose/golangfilemanager/shared" // Adjust the import path as necessary

	"github.com/spf13/cobra"
)

var projects = &cobra.Command{
	Use:   "projects",
	Short: "Lists all the projects.",
	Run: func(cmd *cobra.Command, args []string) {
		// Access the shared.ProjectDir
		categories, err := os.ReadDir(shared.ProjectDir)
		if err != nil {
			fmt.Println("Error getting projects: ", err)
			return
		}

		fmt.Println("Projects:")
		for _, category := range categories {
			if category.IsDir() {
				fmt.Println(category.Name())
				// Construct the path to the subdirectory
				catName := strings.ToLower(category.Name())
				subDirPath := filepath.Join(shared.ProjectDir, catName)
				categoryProjects, err := os.ReadDir(subDirPath)
				if err != nil {
					fmt.Printf("Error reading category %s: %v\n", catName, err)
					continue
				}

				// Print the contents of the category directory
				fmt.Printf("Contents of %s:\n", catName)
				for _, project := range categoryProjects {
					fmt.Println("  -", project.Name())
				}
			}
		}
	},
}

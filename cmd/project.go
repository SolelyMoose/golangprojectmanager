package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/solelymoose/golangfilemanager/shared" // Adjust the import path as necessary

	"github.com/spf13/cobra"
)

var project = &cobra.Command{
	Use:   "project",
	Short: "Manage projects with actions like create, delete, and open",
}

var create = &cobra.Command{
	Use:   "create [project-name] [language] [type]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.ToLower(args[0])
		language := strings.ToLower(args[1])
		projectType := strings.ToLower(args[2])

		// Check if the project already exists
		if checkForProject(language, projectName) {
			fmt.Printf("Project '%s' already exists in language category '%s'.\n", projectName, language)
			return
		}

		// Create project directory
		projectPath := filepath.Join(shared.ProjectDir, language, projectName)
		if err := os.MkdirAll(projectPath, 0755); err != nil {
			fmt.Printf("Error creating project directory '%s': %v\n", projectPath, err)
			return
		}

		fmt.Printf("Successfully created project '%s' under language '%s' and type '%s'.\n", projectName, language, projectType)
	},
}

var delete = &cobra.Command{
	Use:   "delete [project-name] [language]",
	Short: "Move a project to the trash",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.ToLower(args[0])
		language := strings.ToLower(args[1])

		// Confirm deletion
		fmt.Printf("Are you sure you want to move project '%s' to the trash? Type '%s' to confirm: ", projectName, projectName)
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) != projectName {
			fmt.Println("Deletion cancelled.")
			return
		}

		// Move project to trash
		src := filepath.Join(shared.ProjectDir, language, projectName)
		dest := filepath.Join(shared.TrashDir, projectName)
		if err := os.Rename(src, dest); err != nil {
			fmt.Printf("Error moving project '%s' to trash: %v\n", projectName, err)
			return
		}

		fmt.Printf("Project '%s' has been moved to the trash.\n", projectName)
	},
}

var open = &cobra.Command{
	Use:   "open [project-name] [language]",
	Short: "Open an existing project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.ToLower(args[0])
		language := strings.ToLower(args[1])

		// Verify the project exists
		if !checkForProject(language, projectName) {
			fmt.Printf("Project '%s' not found in language category '%s'.\n", projectName, language)
			return
		}

		// Simulate opening the project
		projectPath := filepath.Join(shared.ProjectDir, language, projectName)
		fmt.Printf("Opening project '%s' at path '%s'.\n", projectName, projectPath)
	},
}

func checkForProject(language string, projectName string) bool {
	subDirPath := filepath.Join(shared.ProjectDir, language)

	// Check if the category exists
	categoryProjects, err := os.ReadDir(subDirPath)
	if err != nil {
		return false
	}

	for _, project := range categoryProjects {
		if project.IsDir() && strings.ToLower(project.Name()) == projectName {
			return true
		}
	}
	return false
}

func ExecuteProjectCLI() {
	shared.InitSharedVariables()
	project.AddCommand(create, delete, open)

	if err := project.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
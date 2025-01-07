package shared

import (
	"fmt"
	"os"
	"path/filepath"
)

var MainDir string
var ProjectDir string
var TrashDir string
var ConfigFile string

func InitSharedVariables() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}
	MainDir = filepath.Dir(ex)
	ProjectDir = filepath.Join(MainDir, "Projects")
	TrashDir = filepath.Join(ProjectDir, "trash")
	ConfigFile = filepath.Join(MainDir, "config.json")

	// Ensure base directories exist
	if err := os.MkdirAll(ProjectDir, 0755); err != nil {
		fmt.Println("Error creating project directory:", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(TrashDir, 0755); err != nil {
		fmt.Println("Error creating trash directory:", err)
		os.Exit(1)
	}
}

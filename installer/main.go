package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get the current executable path
	ex, err := os.Executable()
	if err != nil {
		// If we can't get the executable path, print the error and exit after waiting for input
		fmt.Println("Error getting executable path:", err)
		fmt.Println("Press any button to exit.")
		fmt.Scanln() // Wait for input
		return
	}

	// Define the file paths relative to the executable
	folderDir := filepath.Dir(ex)                                 // Get the directory of the executable
	filesDir := filepath.Join(folderDir, "files")                 // Directory where the project manager files are located
	exeDir := filepath.Join(filesDir, "golangprojectmanager.exe") // Path to the executable

	// Prompt the user if they want to install the Golang Project Manager
	var input string
	fmt.Println("Install Golang Project Manager? (y/n)")
	fmt.Scanln(&input)

	// If the user doesn't input "y", cancel the installation
	if strings.ToLower(input) != "y" {
		fmt.Println("Cancelling install.")
		fmt.Println("Press any button to continue.")
		fmt.Scanln() // Wait for input
		return
	}
	// Print a message indicating that installation is starting
	fmt.Printf("Installing Golang Project Manager...\n\n")

	// Check if the Golang Project Manager is already installed by checking the directory
	_, err = os.Stat("C:/Program Files/GolangProjectManager")
	if err == nil {
		// If the directory exists, prompt the user to uninstall the project manager
		fmt.Println("Golang Project Manager is already installed, uninstall? (y/n)")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "n" {
			// If the user doesn't want to uninstall, cancel the uninstallation
			fmt.Println("Cancelling uninstall.")
			fmt.Println("Press any button to continue.")
			fmt.Scanln() // Wait for input
			return
		}

		// Attempt to move the "projects" directory to the Desktop (as part of uninstalling)
		err := os.Rename("C:/Program Files/GolangProjectManager/projects", "Desktop")
		if err != nil {
			// If an error occurs while moving the "projects" directory, print the error and return
			fmt.Println("Error uninstalling (couldn't move projects folder to desktop):", err)
			fmt.Println("Press any button to continue.")
			fmt.Scanln() // Wait for input
			return
		}

		// Attempt to delete the GolangProjectManager folder
		err = os.RemoveAll("C:/Program Files/GolangProjectManager")
		if err != nil {
			fmt.Println("Error deleting GolangProjectManager main folder:", err)
		}

		// Attempt to remove the executable from "C:/Windows"
		err = os.Remove("C:/Windows/golangprojectmanager.exe")
		if err != nil {
			// If an error occurs while removing the executable, print the error and return
			fmt.Println("Error uninstalling (couldn't remove executable):", err)
		}

		// Print success message after uninstalling
		fmt.Println("Golang Project Manager successfully uninstalled.")
		fmt.Println("Press any button to continue.")
		fmt.Scanln() // Wait for input
		return
	} else if !os.IsNotExist(err) {
		// If there was an error checking if the directory exists other than "not exist", print it
		fmt.Println("Error checking if Golang Project Manager is already installed:", err)
		fmt.Println("Press any button to continue.")
		fmt.Scanln() // Wait for input
		return
	}

	// Proceed with installation since the Golang Project Manager isn't already installed
	// Move the executable to the "C:/Windows" directory
	err = os.Rename(exeDir, "C:/Windows/golangprojectmanager.exe")
	if err != nil {
		// If an error occurs while moving the executable, print the error and return
		fmt.Println("Error installing Golang Project Manager:", err)
		fmt.Println("Press any button to continue.")
		fmt.Scanln() // Wait for input
		return
	}

	// Move the "files" directory to "C:/Program Files"
	err = os.Rename(filesDir, "C:/Program Files/GolangProjectManager")
	if err != nil {
		// If an error occurs while moving the files directory, print the error and return
		fmt.Println("Error moving files to Program Files:", err)
		fmt.Println("Press any button to continue.")
		fmt.Scanln() // Wait for input
		return
	}

	// Print success message after installation
	fmt.Println("Golang Project Manager installed successfully.")
	fmt.Println("Press any button to continue.")
	fmt.Scanln() // Wait for input
}

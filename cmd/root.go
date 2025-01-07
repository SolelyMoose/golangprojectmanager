package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpm",
	Short: "A simple project manager made in golang.",
	Long: `gpm [command]
	
	commands:
	version
	project [create, delete, open]
	projects
	config
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golangprojectmanager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.AddCommand(version)
	rootCmd.AddCommand(project)
	rootCmd.AddCommand(projects)
	rootCmd.AddCommand(config)
}

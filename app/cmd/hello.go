package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// Register the hello subcommand to the base rootCmd
	rootCmd.AddCommand(helloCmd)
}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints Hello World",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!!!")
	},
}

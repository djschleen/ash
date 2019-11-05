// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	useColor bool
	output   string

	rootCmd = &cobra.Command{
		Use:     "ash",
		Short:   `Application Security Health score calculator`,
		Version: "0.0.1",
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&useColor, "color", "c", true, "Renders command output in color (unimplemented)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Sends the command output to a specified file (unimplemented)")
}

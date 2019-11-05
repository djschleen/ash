// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TylerBrock/colorjson"
	"github.com/djschleen/ash/calc"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Returns information about a specific CVE.",
	Long:    `Returns information about a specific CVE.`,
	Example: "ash info CVE-yyyy-xxxx",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("ASH requires a CVE to be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		outputLong(args)
	},
}

func outputLong(args []string) {
	b := calc.Info(args[0]) //"CVE-2010-3333"
	str, _ := json.Marshal(b)
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

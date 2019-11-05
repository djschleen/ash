package main

import (
	"os"

	"github.com/djschleen/ash/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}

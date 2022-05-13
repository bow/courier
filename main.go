package main

import (
	"os"

	"github.com/bow/courier/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

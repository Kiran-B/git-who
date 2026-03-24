package main

import (
	"fmt"
	"os"

	"github.com/Kiran-B/git-who/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"example.com/codelytv/cobra-01/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
}
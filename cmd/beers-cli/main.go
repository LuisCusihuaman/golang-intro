package main

import (
	"github.com/LuisCusihuaman/golang-introduction/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	_ = rootCmd.Execute()
}

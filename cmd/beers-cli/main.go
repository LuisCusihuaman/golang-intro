package main

import (
	"github.com/LuisCusihuaman/golang-introduction/internal/cli"
	"github.com/LuisCusihuaman/golang-introduction/internal/storage/csv"
	"github.com/spf13/cobra"
)

//go run cmd/beers-cli/main.go beers --id 127
func main() {

	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	_ = rootCmd.Execute()
}

package main

import (
	"flag"
	beerscli "github.com/LuisCusihuaman/golang-introduction/internal"
	"github.com/LuisCusihuaman/golang-introduction/internal/cli"
	"github.com/LuisCusihuaman/golang-introduction/internal/storage/csv"
	"github.com/LuisCusihuaman/golang-introduction/internal/storage/ontario"
	"github.com/spf13/cobra"
)

//go run cmd/beers-cli/main.go beers --id 127
func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	_ = rootCmd.Execute()
}

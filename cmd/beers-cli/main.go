package main

import (
	"flag"
	beerscli "github.com/LuisCusihuaman/golang-introduction/internal"
	"github.com/LuisCusihuaman/golang-introduction/internal/cli"
	"github.com/LuisCusihuaman/golang-introduction/internal/fetching"
	"github.com/LuisCusihuaman/golang-introduction/internal/storage/csv"
	"github.com/LuisCusihuaman/golang-introduction/internal/storage/ontario"
	"github.com/spf13/cobra"
)

//go run cmd/beers-cli/main.go beers --id 127
func main() {

	httpData := flag.Bool("http", false, "load data from http")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if *httpData {
		repo = ontario.NewOntarioRepository()
	}
	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	_ = rootCmd.Execute()
}

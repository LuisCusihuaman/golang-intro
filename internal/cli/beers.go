package cli

import (
	"fmt"
	beerscli "github.com/LuisCusihuaman/golang-introduction/internal"
	"github.com/spf13/cobra"
	"strconv"
)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(repository beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func runBeersFn(repository beerscli.BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers, _ := repository.GetBeers()

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ProductID == i {
					fmt.Println(beer)
					return
				}
			}
		} else {
			fmt.Println(beers)
		}
	}
}

package cli

import (
	"fmt"
	"github.com/LuisCusihuaman/golang-introduction/internal/fetching"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(service fetching.Service) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(service),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

// CobraFn function definition of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func runBeersFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetString(idFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			beer, err := service.FetchByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(beer)
			return
		}

		beers, err := service.FetchBeers()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(beers)
	}
}

package cmd

import (
	"fmt"
	"math"

	"github.com/djschleen/ash/calc"
	"github.com/spf13/cobra"
)

// calculateCmd represents the calculate command
var (
	cveIDList    []string
	shortCves    []calc.ShortCve
	calculateCmd = &cobra.Command{
		Use:   "calculate",
		Short: "Calculates the Application Security Health score",
		Long:  `Calculates the Application Security Health score`,
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range cveIDList {
				c := calc.Info(v)
				shortCves = append(shortCves, c.ToShortCve())
			}
			ash := Calculate(shortCves)
			fmt.Println(ash)
		},
	}
)

func init() {
	calculateCmd.Flags().StringSliceVarP(&cveIDList, "identifiers", "i", []string{}, "")
	rootCmd.AddCommand(calculateCmd)
}

//Calculate Computes the ASH ash score
func Calculate(shortCves []calc.ShortCve) int64 {
	//TODO: Weight by date
	//TODO: Weight by impact
	//TODO: Weight by ease of exploit
	//TODO: Deduplicate

	sumWeight, avg := 0.0, 0.0
	for _, v := range shortCves {
		sumWeight += v.SeverityWeight()
		avg += v.Cvss * v.SeverityWeight()
	}
	avg /= sumWeight

	avg *= -100
	avg = avg + 1000
	return int64(math.RoundToEven(avg))
}

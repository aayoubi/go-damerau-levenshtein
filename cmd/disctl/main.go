package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/aayoubi/dameraulevenshtein/internal/distance"
)

func main() {
	var s1, s2 string

	app := &cli.App{
		Name:  "disctl",
		Usage: "compute the Damerau Levenshtein distance between two strings",
		Action: func(c *cli.Context) error {
			s1 = "hey"
			s2 = "hei"
			if c.NArg() > 1 {
				s1 = c.Args().Get(0)
				s2 = c.Args().Get(1)
			}
			fmt.Printf("The distance between [%s] and [%s] is %d\n", s1, s2, distance.DistanceDamerauLevenshtein(s1, s2))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

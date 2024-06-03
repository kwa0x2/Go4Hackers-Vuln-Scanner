package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func Commands() {
	app := &cli.App{
		Name:      "Go4Hackers Small Vulnerability Scanner",
		UsageText: "./go4hackers-vuln-scanner --target <TARGET_URL> --wordlist <WORDLIST_PATH> --dirlisting ",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "target", Usage: "Target Web URL"},
			cli.StringFlag{Name: "wordlist", Usage: "Type the wordlist directory"},
			cli.BoolFlag{Name: "trace", Usage: "Enables TRACE method checker."},
			cli.BoolFlag{Name: "x-frame-options", Usage: "Enables X-Frame-Options header checker."},
			cli.BoolFlag{Name: "dirlisting", Usage: "Enables directory listing vulnerability checker."},
			cli.IntFlag{Name: "delay", Usage: "Delay in miliseconds between each HTTP request", Value: 0},
			cli.BoolFlag{Name: "list-wordlists", Usage: "./wordlist lists word lists in the file"},
		},

		Action: func(c *cli.Context) error {
			if c.NumFlags() == 0 {
				cli.ShowAppHelp(c)
				return nil
			}
			if c.Bool("list-wordlists") {
				dirPath := "./wordlists"

				files, err := os.ReadDir(dirPath)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(Blue("Files in", dirPath, ":"))

				for _, file := range files {
					fmt.Println(Green(file.Name()))
				}
			}
			if c.String("target") != "" {
				if c.Bool("trace") {
					CheckTrace(c.String("target"))
				}
				if c.Bool("x-frame-options") {
					CheckXFrameOptions(c.String("target"))
				}
				if c.String("wordlist") != "" {
					if c.Bool("dirlisting") {

						CheckDirListing(c.String("wordlist"), c.String("target"), c.Int("delay"))
					}

				} else {
					fmt.Println("Please use --wordlist <WORDLIST.txt>")
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

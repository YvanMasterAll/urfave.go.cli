package main

import (
	"os"

	"fmt"
	"github.com/urfave/cli"
	"sort"
)

func main() {
	appInit()
}

//App Init
func appInit() {
	var language string
	var configure string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l", //--lang | -l
			Value:       "english",
			Usage:       "language for the greeting",
			EnvVar:      "APP_LANG, DEFAULT_LANG", //environment variable
			Destination: &language,
		},
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Load configuration from `FILE`",
			Destination: &configure,
		},
		cli.BoolTFlag{
			Name:  "exit, e",
			Usage: "Exit test",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println(language)
				if !c.Bool("exit") {
					return cli.NewExitError("it is a exit test!", -1) //exit
				}
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:     "add",
			Category: "template",
		},
		{
			Name:     "remove",
			Category: "template",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	//app.Action = func(c *cli.Context) error {
	//  name := "Nefertiti"
	//  if c.NArg() > 0 {
	//      name = c.Args().Get(0)
	//  }
	//  if c.String("lang") == "spanish" {
	//      fmt.Printf("Hola %s", name)
	//  } else {
	//      fmt.Printf("Hello %s", name)
	//  }
	//  return nil
	//}

	app.Name = "prime"
	app.Version = "01.01.01"
	app.Run(os.Args)
}

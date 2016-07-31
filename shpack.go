package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func command_install(c *cli.Context) error {
	fmt.Println(c.Args())
	return nil
}

func command_list(c *cli.Context) error {
	fmt.Println(c.Args())
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "shpack"
	app.Usage = "Lightweight package manager for shellscripts and binary files. Don't give up managing your ad-hoc scripts!"

	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install scripts from the specified `PATH`",
			Action:  command_install,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list installed scripts",
			Action:  command_list,
		},
	}

	setupConfig()

	app.Run(os.Args)
}

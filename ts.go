package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ts"
	app.Version = Version
	app.Usage = ""
	app.Author = "timakin"
	app.Email = "timaki.st@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}

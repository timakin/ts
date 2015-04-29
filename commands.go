package main

import (
	"log"
	"os"
	"fmt"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
	commandBiz,
	commandHack,
}

var commandAll = cli.Command{
	Name:  "all",
	Usage: "",
	Description: `
`,
	Action: doAll,
}

var commandBiz = cli.Command{
	Name:  "biz",
	Usage: "",
	Description: `
`,
	Action: doBiz,
}

var commandHack = cli.Command{
	Name:  "hack",
	Usage: "",
	Description: `
`,
	Action: doHack,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAll(c *cli.Context) {
	fmt.Printf("print all")
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}

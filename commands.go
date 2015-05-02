package main

import (
	"github.com/timakin/ts/loader"
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

func doAll(c *cli.Context) {
	loader.GetRedditFeed()
	hn := make(chan loader.ResultData)
	ph := make(chan loader.ResultData)
	go loader.GetHNFeed(hn)
	go loader.GetPHFeed(ph)
	hnres := <- hn
	phres := <- ph
	var HNData loader.Feed = &hnres
	var PHData loader.Feed = &phres
	HNData.Display()
	PHData.Display()
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}

package main

import (
	"fmt"
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
//	commandBiz,
//	commandHack,
}

var commandAll = cli.Command{
	Name:  "pop",
	Usage: "",
	Description: "Show today's news from major tech news sites, HN, PH, and subreddit of /programming.",
	Action: doAll,
}

//var commandBiz = cli.Command{
//	Name:  "biz",
//	Usage: "",
//	Description: `
//`,
//	Action: doBiz,
//}
//var commandHack = cli.Command{
//	Name:  "test",
//	Usage: "",
//	Description: `
//`,
//	Action: doHack,
//}

func pp(str string) {
  fmt.Printf(str)
}

func doAll(c *cli.Context) {
		hn := make(chan loader.ResultData)
		ph := make(chan loader.ResultData)
		re := make(chan loader.ResultData)
		go loader.GetHNFeed(hn)
		go loader.GetPHFeed(ph)
		go loader.GetRedditFeed(re)
		hnres := <- hn
		phres := <- ph
		reres := <- re
		var HNData loader.Feed = &hnres
		var PHData loader.Feed = &phres
		var REData loader.Feed = &reres
		HNData.Display()
		PHData.Display()
		REData.Display()

		pp("[Hatena]\n")
		loader.GetHatenaFeed()
}

//func doBiz(c *cli.Context) {
//}
//func doHack(c *cli.Context) {
//}

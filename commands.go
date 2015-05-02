package main

import (
	"fmt"
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
//	commandBiz,
	commandHack,
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
var commandHack = cli.Command{
	Name:  "test",
	Usage: "",
	Description: `
`,
	Action: doHack,
}

func pp(str string) {
  fmt.Printf(str)
}

func doAll(c *cli.Context) {
		pp("▁ ▂ ▄ ▅ ▆ ▇ █ тecнѕтacĸ █ ▇ ▆ ▅ ▄ ▂ ▁\n")
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

		var uri string
		pp("[TechCrunch]\n")
		uri = "http://feeds.feedburner.com/TechCrunch/"
		loader.GetRSSFeed(uri)
		pp("[Mashable]\n")
		uri = "http://feeds.mashable.com/Mashable"
		loader.GetRSSFeed(uri)
		pp("[Hatena]\n")
		uri = "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss"
		loader.GetRSSFeed(uri)
}

//func doBiz(c *cli.Context) {
//}
func doHack(c *cli.Context) {
}

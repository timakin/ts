package main

import (
	"fmt"
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
	commandHack,
	commandPH,
	commandRE,
	commandHN,
	commandMS,
	commandDN,
	commandFB,
	commandEJ,
	commandRD,
	commandHatena,
}

var commandAll = cli.Command{
	Name:  "pop",
	Usage: "",
	Description: "Show today's news from major tech news sites, HN, PH, and subreddit of /programming.",
	Action: doAll,
}

var commandHack = cli.Command{
	Name:  "hack",
	Usage: "",
	Description: `
`,
	Action: doHack,
}

var commandPH = cli.Command{
	Name:  "ph",
	Usage: "",
	Description: `
`,
	Action: doPH,
}

var commandHN = cli.Command{
	Name:  "hn",
	Usage: "",
	Description: `
`,
	Action: doHN,
}

var commandRE = cli.Command{
	Name:  "reddit",
	Usage: "",
	Description: `
`,
	Action: doRE,
}

var commandTC = cli.Command{
	Name:  "tc",
	Usage: "",
	Description: `
`,
	Action: doTC,
}

var commandMS = cli.Command{
	Name:  "ms",
	Usage: "",
	Description: `
`,
	Action: doMS,
}

var commandDN = cli.Command{
	Name:  "dn",
	Usage: "",
	Description: `
`,
	Action: doDN,
}

var commandFB = cli.Command{
	Name:  "forbes",
	Usage: "",
	Description: `
`,
	Action: doFB,
}

var commandEJ = cli.Command{
	Name:  "echojs",
	Usage: "",
	Description: `
`,
	Action: doEJ,
}

var commandRD = cli.Command{
	Name:  "rdaily",
	Usage: "",
	Description: `
`,
	Action: doRD,
}

var commandHatena = cli.Command{
	Name:  "hatena",
	Usage: "",
	Description: `
`,
	Action: doHatena,
}

func pp(str string) {
  fmt.Printf(str)
}

func displayRSSFeed(name string, uri string) {
	pp("[" + name + "]\n")
	loader.GetRSSFeed(uri)
}

func doAll(c *cli.Context) {
		pp("▁ ▂ ▄ ▅ ▆ ▇ █ тecнѕтacĸ █ ▇ ▆ ▅ ▄ ▂ ▁\n\n")
		ph := make(chan loader.ResultData)
		re := make(chan loader.ResultData)
		go loader.GetPHFeed(ph)
		go loader.GetRedditFeed(re)
		phres := <- ph
		reres := <- re
		var PHData loader.Feed = &phres
		var REData loader.Feed = &reres
		PHData.Display()
		REData.Display()
		displayRSSFeed("HackerNews", "https://news.ycombinator.com/rss")
		displayRSSFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/")
		displayRSSFeed("Mashable", "http://feeds.mashable.com/Mashable")
		displayRSSFeed("Designer News", "https://news.layervault.com/?format=rss")
		displayRSSFeed("Forbes - Tech", "http://www.forbes.com/technology/feed/")
		displayRSSFeed("EchoJS", "http://www.echojs.com/rss")
		displayRSSFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily")
}

func doHack(c *cli.Context) {
	re := make(chan loader.ResultData)
	go loader.GetRedditFeed(re)
	reres := <- re
	var REData loader.Feed = &reres
	REData.Display()
	displayRSSFeed("HackerNews", "https://news.ycombinator.com/rss")
	displayRSSFeed("EchoJS", "http://www.echojs.com/rss")
	displayRSSFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily")
}

func doPH(c *cli.Context) {
	ph := make(chan loader.ResultData)
	go loader.GetPHFeed(ph)
	phres := <- ph
	var PHData loader.Feed = &phres
	PHData.Display()
}

func doRE(c *cli.Context) {
	re := make(chan loader.ResultData)
	go loader.GetRedditFeed(re)
	reres := <- re
	var REData loader.Feed = &reres
	REData.Display()
}

func doHN(c *cli.Context) {
	displayRSSFeed("HackerNews", "https://news.ycombinator.com/rss")
}

func doTC(c *cli.Context) {
	displayRSSFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/")
}

func doMS(c *cli.Context) {
	displayRSSFeed("Mashable", "http://feeds.mashable.com/Mashable")
}

func doDN(c *cli.Context) {
	displayRSSFeed("Designer News", "https://news.layervault.com/?format=rss")
}

func doFB(c *cli.Context) {
	displayRSSFeed("Forbes - Tech", "http://www.forbes.com/technology/feed/")
}

func doEJ(c *cli.Context) {
	displayRSSFeed("EchoJS", "http://www.echojs.com/rss")
}

func doRD(c *cli.Context) {
	displayRSSFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily")
}

func doHatena(c *cli.Context) {
	displayRSSFeed("Hatena", "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss")
}

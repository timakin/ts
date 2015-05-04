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
	commandTC,
	commandRE,
	commandHN,
	commandGH,
	commandMS,
	commandTNW,
	commandDN,
	commandFB,
	commandEJ,
	commandRD,
	commandA16Z,
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

var commandGH = cli.Command{
	Name:  "github",
	Usage: "",
	Description: `
`,
	Action: doGH,
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

var commandTNW = cli.Command{
	Name:  "tnw",
	Usage: "",
	Description: `
`,
	Action: doTNW,
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

var commandA16Z = cli.Command{
	Name:  "a16z",
	Usage: "",
	Description: `
`,
	Action: doA16Z,
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

func ppred(str string) {
	fmt.Printf("\033[1;31m" + str + "\033[0m")
}

func displayRSSFeed(name string, uri string) {
	ppred("[" + name + "]\n")
	loader.GetRSSFeed(uri)
}

func displayRSSFeedWithDescription(name string, uri string) {
	ppred("[" + name + "]\n")
	loader.GetRSSFeedWithDescription(uri)
}

func doAll(c *cli.Context) {
		ppred("▁ ▂ ▄ ▅ ▆ ▇ █ тecнѕтacĸ █ ▇ ▆ ▅ ▄ ▂ ▁\n\n")
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
		displayRSSFeedWithDescription("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss")
		displayRSSFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/")
		displayRSSFeed("Mashable", "http://feeds.mashable.com/Mashable")
		displayRSSFeed("The Next Web", "http://feeds2.feedburner.com/thenextweb")
		displayRSSFeed("Designer News", "https://news.layervault.com/?format=rss")
		displayRSSFeed("Forbes - Tech", "http://www.forbes.com/technology/feed/")
		displayRSSFeed("EchoJS", "http://www.echojs.com/rss")
		displayRSSFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily")
		displayRSSFeed("A16Z", "http://a16z.com/feed/")
}

func doHack(c *cli.Context) {
	re := make(chan loader.ResultData)
	go loader.GetRedditFeed(re)
	reres := <- re
	var REData loader.Feed = &reres
	REData.Display()
	displayRSSFeed("HackerNews", "https://news.ycombinator.com/rss")
	displayRSSFeedWithDescription("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss")
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

func doGH(c *cli.Context) {
	displayRSSFeedWithDescription("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss")
}

func doTC(c *cli.Context) {
	displayRSSFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/")
}

func doMS(c *cli.Context) {
	displayRSSFeed("Mashable", "http://feeds.mashable.com/Mashable")
}

func doTNW(c *cli.Context) {
	displayRSSFeed("The Next Web", "http://feeds2.feedburner.com/thenextweb")
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

func doA16Z(c *cli.Context) {
	displayRSSFeed("A16Z", "http://a16z.com/feed/")
}

func doHatena(c *cli.Context) {
	displayRSSFeed("Hatena", "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss")
}

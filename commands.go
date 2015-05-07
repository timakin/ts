package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/getwe/figlet4go"
	"github.com/timakin/ts/loader"
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
	Name:        "pop",
	Usage:       "",
	Description: "Show today's news from major tech news sites, HN, PH, and subreddit of /programming.",
	Action:      doAll,
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

func displayAA() {
	flag_str := flag.String("str", "TechStack", "input string")
	flag.Parse()
	str := *flag_str
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	ascii.LoadFont("./fonts/")
	renderStr, _ := ascii.RenderOpts(str, options)
	ppred(renderStr + "\n\n")
}

func displayUnitRssFeed(name string, uri string) {
	ppred("[" + name + "]\n")
	loader.GetUnitRssFeed(uri)
}

func displayUnitRssFeedWithDesc(name string, uri string) {
	ppred("[" + name + "]\n")
	loader.GetUnitRssFeedWithDesc(uri)
}

func doAll(c *cli.Context) {
	displayAA()
	ph := make(chan loader.ResultData)
	re := make(chan loader.ResultData)
	hn := make(chan loader.ResultData)
	gh := make(chan loader.ResultData)
	tc := make(chan loader.ResultData)
	ms := make(chan loader.ResultData)
	tnw := make(chan loader.ResultData)
	dn := make(chan loader.ResultData)
	fbs := make(chan loader.ResultData)
	rdaily := make(chan loader.ResultData)
	ejs := make(chan loader.ResultData)
	a16z := make(chan loader.ResultData)
	go loader.GetPHFeed(ph)
	go loader.GetRedditFeed(re)
	go loader.GetRssFeed("HackerNews", "https://news.ycombinator.com/rss", hn)
	go loader.GetRdfFeedWithDesc("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss", gh)
	go loader.GetRssFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/", tc)
	go loader.GetRssFeed("Mashable", "http://feeds.mashable.com/Mashable", ms)
	go loader.GetRssFeed("The Next Web", "http://feeds2.feedburner.com/thenextweb", tnw)
	go loader.GetRssFeed("Designer News", "https://news.layervault.com/?format=rss", dn)
	go loader.GetRssFeed("Forbes - Tech", "http://www.forbes.com/technology/feed/", fbs)
	go loader.GetRssFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily", rdaily)
	go loader.GetRssFeed("EchoJS", "http://www.echojs.com/rss", ejs)
	go loader.GetRssFeed("A16Z", "http://a16z.com/feed/", a16z)
	phres := <-ph
	reres := <-re
	hnres := <-hn
	ghres := <-gh
	tcres := <-tc
	msres := <-ms
	tnwres := <-tnw
	dnres := <-dn
	fbsres := <-fbs
	rdailyres := <-rdaily
	ejsres := <-ejs
	a16zres := <-a16z
	var PHData loader.Feed = &phres
	var REData loader.Feed = &reres
	var HNData loader.Feed = &hnres
	var GHData loader.Feed = &ghres
	var TCData loader.Feed = &tcres
	var MSData loader.Feed = &msres
	var TNWData loader.Feed = &tnwres
	var DNData loader.Feed = &dnres
	var FBSData loader.Feed = &fbsres
	var RDailyData loader.Feed = &rdailyres
	var EJSData loader.Feed = &ejsres
	var A16ZData loader.Feed = &a16zres
	PHData.Display()
	REData.Display()
	HNData.Display()
	TCData.Display()
	GHData.Display()
	MSData.Display()
	TNWData.Display()
	DNData.Display()
	FBSData.Display()
	RDailyData.Display()
	EJSData.Display()
	A16ZData.Display()
}

func doHack(c *cli.Context) {
	re := make(chan loader.ResultData)
	hn := make(chan loader.ResultData)
	gh := make(chan loader.ResultData)
	rdaily := make(chan loader.ResultData)
	ejs := make(chan loader.ResultData)
	go loader.GetRedditFeed(re)
	go loader.GetRssFeed("HackerNews", "https://news.ycombinator.com/rss", hn)
	go loader.GetRdfFeedWithDesc("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss", gh)
	go loader.GetRssFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily", rdaily)
	go loader.GetRssFeed("EchoJS", "http://www.echojs.com/rss", ejs)
	reres := <-re
	hnres := <-hn
	ghres := <-gh
	rdailyres := <-rdaily
	ejsres := <-ejs
	var REData loader.Feed = &reres
	var HNData loader.Feed = &hnres
	var GHData loader.Feed = &ghres
	var RDailyData loader.Feed = &rdailyres
	var EJSData loader.Feed = &ejsres
	REData.Display()
	HNData.Display()
	GHData.Display()
	RDailyData.Display()
	EJSData.Display()
}

func doPH(c *cli.Context) {
	ph := make(chan loader.ResultData)
	go loader.GetPHFeed(ph)
	phres := <-ph
	var PHData loader.Feed = &phres
	PHData.Display()
}

func doRE(c *cli.Context) {
	re := make(chan loader.ResultData)
	go loader.GetRedditFeed(re)
	reres := <-re
	var REData loader.Feed = &reres
	REData.Display()
}

func doHN(c *cli.Context) {
	displayUnitRssFeed("HackerNews", "https://news.ycombinator.com/rss")
}

func doGH(c *cli.Context) {
	displayUnitRssFeedWithDesc("Github Trends", "http://github-trends.ryotarai.info/rss/github_trends_all_daily.rss")
}

func doTC(c *cli.Context) {
	displayUnitRssFeed("TechCrunch", "http://feeds.feedburner.com/TechCrunch/")
}

func doMS(c *cli.Context) {
	displayUnitRssFeed("Mashable", "http://feeds.mashable.com/Mashable")
}

func doTNW(c *cli.Context) {
	displayUnitRssFeed("The Next Web", "http://feeds2.feedburner.com/thenextweb")
}

func doDN(c *cli.Context) {
	displayUnitRssFeed("Designer News", "https://news.layervault.com/?format=rss")
}

func doFB(c *cli.Context) {
	displayUnitRssFeed("Forbes - Tech", "http://www.forbes.com/technology/feed/")
}

func doEJ(c *cli.Context) {
	displayUnitRssFeed("EchoJS", "http://www.echojs.com/rss")
}

func doRD(c *cli.Context) {
	displayUnitRssFeed("RubyDaily", "http://feeds.rubydaily.org/RubyDaily")
}

func doA16Z(c *cli.Context) {
	displayUnitRssFeed("A16Z", "http://a16z.com/feed/")
}

func doHatena(c *cli.Context) {
	displayUnitRssFeed("Hatena", "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss")
}

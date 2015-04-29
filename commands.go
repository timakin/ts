package main

import (
	"log"
	"os"
	"fmt"
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
	test := make(chan int)
	test2 := make(chan int)
	hn := make(chan []int)
	go loader.GoRouTest(test)
	go loader.GoRouTestTwo(test2)
	go loader.GetHNFeed(hn)
	fmt.Printf("print all\n")
	result := <- test
	resTwo := <- test2
	phres := <- hn
	fmt.Printf("%d%d\n", result, resTwo)
	fmt.Printf("%d",phres[0:10])
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}

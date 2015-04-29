package main

import (
	"log"
	"os"
	"fmt"
	"time"

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

func goRouTest(test chan int) {
	test <- 10
}

func goRouTestTwo(test2 chan int) {
	time.Sleep(time.Second * 3)
	test2 <- 20
}


func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAll(c *cli.Context) {
	test := make(chan int)
	test2 := make(chan int)
	go goRouTest(test)
	go goRouTestTwo(test2)
	fmt.Printf("print all\n")
	result := <- test
	resTwo := <- test2
	fmt.Printf("%d%d\n", result, resTwo)
}

func doBiz(c *cli.Context) {
}

func doHack(c *cli.Context) {
}

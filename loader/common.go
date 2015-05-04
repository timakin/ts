package loader

import (
  "fmt"
  "strings"
)

func pp(str string) {
  fmt.Printf(str)
}

func ppred(str string) {
	fmt.Printf("\033[1;31m" + str + "\033[0m")
}

func perror(err error) {
  if err != nil {
    panic(err)
  }
}

func removeBreak(str string) (result string) {
  result = strings.Replace(str, "\n", " ", -1)
  return result
}

package loader

import (
  "fmt"
  "strings"
)

func pp(str string) {
  fmt.Printf(str)
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

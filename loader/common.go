package loader

import (
  "fmt"
)

func pp(str string) {
  fmt.Printf(str)
}

func perror(err error) {
  if err != nil {
    panic(err)
  }
}

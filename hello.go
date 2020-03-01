package main

import (
	"os"
  "log"
  "fmt"
  "bytes"
	"github.com/brodier/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"strings"
)

func main() {
  var (
    buf bytes.Buffer
    loggerAppender = log.New(&buf, "INFO: ", log.Lshortfile)

    logger = func(info string) {
      loggerAppender.Output(2, info)
    }
  )
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(strings.Map(morestrings.Rot13, "'Twas brillig and the slithy gopher..."))
  f, err := os.Open("go.sum")
  data := make([]byte, 1024)
  count, err := f.Read(data)
  if err != nil {
    logger(err.Error())
  }
  fmt.Println(string(data[:count]))
}

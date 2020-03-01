package main

import (
	"fmt"
  "strings"
	"github.com/brodier/hello/morestrings"
  "github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
  fmt.Println(strings.Map(morestrings.Rot13, "'Twas brillig and the slithy gopher..."))
}

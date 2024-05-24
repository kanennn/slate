package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

func main() {
	var source []byte
	fmt.Scan(&source)
	fmt.Println(string(source))
	
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
 		 panic(err)
	}
	fmt.Println(buf.String())
}
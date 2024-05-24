package main

import (
	"bytes"
	"io"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/yuin/goldmark"
)

func main() {
	source, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
		
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
 		 panic(err)
	}

	width := 12
	height := 12

	target := proto.TargetCreateTarget{URL: "", Width: &width, Height: &height}
	page, _ := rod.New().MustConnect().Page(target)
	html := buf.String()
	
	page.SetDocumentContent(html)
	page.AddStyleTag("", "h1 {color: blue; text-align: center;}")
	page.MustPDF("about.pdf")
}
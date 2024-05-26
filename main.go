package main

import (
	"fmt"
	"os"

	"github.com/kanennn/slate/env"
	"github.com/kanennn/slate/export"
	"github.com/kanennn/slate/impa"
	"github.com/kanennn/slate/markdown"
)


func main() {
	
	if len(os.Args) < 2 {
		println("A subcommand is required") 
	} else {
		subcommand := os.Args[1]
		switch subcommand {
		case "init":
			env.Init()
		case "build":
			build()
		default:
			fmt.Printf("%s: unknown command\n", subcommand)
		}
	}
	
}

func build() {
	
	data := impa.Import()

	html := markdown.Render(data.Markdown)
	export.ExportPdf(html, data.Css)
}
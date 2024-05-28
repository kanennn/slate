package main

import (
	"fmt"
	"os"

	"github.com/kanennn/slate/internal/env"
	"github.com/kanennn/slate/internal/export"
	"github.com/kanennn/slate/internal/impa"
	markdown "github.com/kanennn/slate/internal/render"
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
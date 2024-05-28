package markdown

import (
	"bytes"

	utils "github.com/kanennn/slate/util"
	"github.com/yuin/goldmark"
)


func Render(source []byte) bytes.Buffer {
		
	var buf bytes.Buffer
	utils.Check((goldmark.Convert(source, &buf)))

	return buf
}
package markdown

import (
	"bytes"

	"github.com/kanennn/slate/utils"
	"github.com/yuin/goldmark"
	
)


func Render(source []byte) bytes.Buffer {
		
	var buf bytes.Buffer
	utils.Check((goldmark.Convert(source, &buf)))

	return buf
}
package impa

import (
	"fmt"
	"os"

	"github.com/kanennn/slate/utils"
)

func importMarkdown() []byte {
	file, err := os.ReadFile("master.md")
	if os.IsNotExist(err) {
		fmt.Println("No master.md file. Proceeding with empty markdown content.")
	} else {
		utils.Check(err)
	}
	return file
}
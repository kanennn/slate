package impa

import (
	"fmt"
	"os"

	"github.com/kanennn/slate/utils"
)

func importCss() []byte {
	file, err := os.ReadFile("style.css")
	if os.IsNotExist(err) {
		fmt.Println("No style.css file. Proceeding with empty css.")
	} else {
		utils.Check(err)
	}
	return file
}
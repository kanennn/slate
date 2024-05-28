package impa

import (
	"fmt"
	"os"

	utils "github.com/kanennn/slate/util"
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
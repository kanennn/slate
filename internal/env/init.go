package env

import (
	"fmt"
	"os"
)

func Init() {
	var err error 
	_, err = os.Create("style.css")
	if os.IsExist(err) {
		fmt.Println("style.css already exists. Skipping...")
	}					
	_, err = os.Create("master.md")			
	if os.IsExist(err) {
		fmt.Println("master.md already exists. Skipping...")
	}					
	//os.Mkdir("chapters", 0755)
	//os.Create("chapters/1-hello world.md")
	fmt.Println("Done.")
}


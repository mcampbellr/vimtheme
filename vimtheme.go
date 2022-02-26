package main

import (
	"bufio"
	"fmt"
	"os"
)

var readLine = bufio.NewReader(os.Stdin)

func main() {
	selectedTheme, err := selectTheme()

	if err != nil {
		fmt.Println("Theme not changed: ", err)
		return
	}

	mainSearch := `colorscheme\s[a-zA-Z_]+`
	mainReplace := "colorscheme " + selectedTheme.Theme
	mainFileName := ".config/nvim/lua/user/colorscheme.lua"
	SearchAndReplace(mainSearch, mainReplace, mainFileName, false)
}

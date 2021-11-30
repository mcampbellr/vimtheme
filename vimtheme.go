package main

import (
	"bufio"
	"os"
)

var readLine = bufio.NewReader(os.Stdin)

func main() {
	selectedTheme, _ := selectTheme()
	mainSearch := `colorscheme\s[a-zA-Z_]+`
	mainReplace := "colorscheme " + selectedTheme.Theme
	mainFileName := "nvim/init.vim"
	SearchAndReplace(mainSearch, mainReplace, mainFileName, false)

	lineSearch := `theme\=\s'[a-zA-Z_]+'`
	lineReplace := "theme= '" + selectedTheme.Line + "'"
	lineFileName := "nvim/after/plugin/lualine.lua"
	SearchAndReplace(lineSearch, lineReplace, lineFileName, false)
}

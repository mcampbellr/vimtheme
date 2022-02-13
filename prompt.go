package main

import "github.com/manifoldco/promptui"

type theme struct {
	ThemeName string
	Theme     string
	Line      string
}

var items = []*theme{
	{ThemeName: "Monokai", Theme: "monokai", Line: "molokai"},
	{ThemeName: "NeoSolarized", Theme: "NeoSolarized", Line: "solarized_dark"},
	{ThemeName: "One Dark", Theme: "onedark", Line: "onedark"},
	{ThemeName: "Dracula", Theme: "dracula", Line: "onedark"},
	{ThemeName: "Gruvbox", Theme: "gruvbox", Line: "gruvbox_dark"},
}

func selectTheme() (*theme, error) {
	var selectItems []string

	for i := 0; i < len(items); i++ {
		selectItems = append(selectItems, items[i].ThemeName)
	}

	prompt := promptui.Select{
		Label: "Select Theme:",
		Items: selectItems,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	var res *theme
	for _, v := range items {
		if v.ThemeName == result {
			res = v
		}
	}

	return res, nil
}

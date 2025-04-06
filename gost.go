package main

import (
	_ "Daizen/markdown"
	"Daizen/plugins"
	"Daizen/renderer"
	"Daizen/site"
	"Daizen/theme"
)

func main() {
	err := site.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	plugins.LoadPlugins()
	theme.LoadTheme()
	err = renderer.RenderSite()
	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	_ "github.com/OblivionOcean/Daizen/markdown"
	"github.com/OblivionOcean/Daizen/plugins"
	"github.com/OblivionOcean/Daizen/renderer"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/theme"
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

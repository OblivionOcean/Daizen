package cmd

import (
	"fmt"

	"github.com/OblivionOcean/Daizen/plugins"
	"github.com/OblivionOcean/Daizen/renderer"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/theme"
	"github.com/spf13/cobra"
)

func Generate() {
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

func CMD() {
	// 处理命令行， generate和g为同一个命令，install 和i为同一个命令，不是参数，而是多级命令
	genCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the site",
		Long:  "Generate the site",
		Run: func(cmd *cobra.Command, args []string) {
			Generate()
		},
	}
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install the plugin",
		Long:  "Install the plugin",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please input the plugin name")
				return
			}
			plugins.InstallPlugin(args[0])
		},
	}
	rootCmd := &cobra.Command{
		Use:   "daizen",
		Short: "A static site generator",
		Long:  "A static site generator",
	}
	rootCmd.AddCommand(genCmd, installCmd)
	rootCmd.Execute()
}

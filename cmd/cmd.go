package cmd

import (
	"fmt"
	"os"

	"github.com/OblivionOcean/Daizen/plugins"
	"github.com/OblivionOcean/Daizen/renderer"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/utils"
	"github.com/spf13/cobra"
)

func Generate() {
	os.MkdirAll("./.daizen", os.ModePerm)
	err := site.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	err = renderer.RenderSite()
	if err != nil {
		panic(err.Error())
	}
}

func CMD(global bool) {
	if global && len(os.Args) > 1 {
		execCmd := "./.daizen/Daizen" + plugins.ExecExt()
		if utils.FileExist(execCmd) {
			if os.Args[1] == "reset" {
				goto regcmd
			}
			utils.Exec(execCmd, os.Args[1:]...)
			return
		}
	}
regcmd:
	// 处理命令行， generate和g为同一个命令，install 和i为同一个命令，不是参数，而是多级命令
	genCmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate the site",
		Long:    "Generate the site",
		Run: func(cmd *cobra.Command, args []string) {
			Generate()
		},
	}
	installCmd := &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install the plugin",
		Long:    "Install the plugin",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please input the plugin name")
				return
			}
			plugins.InstallPlugin(args[0])
		},
	}
	uninstallCmd := &cobra.Command{
		Use:     "uninstall",
		Aliases: []string{"uni"},
		Short:   "Uninstall the plugin",
		Long:    "Uninstall the plugin",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please input the plugin name")
				return
			}
			plugins.UninstallPlugin(args[0])
		},
	}
	listPluginCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List the plugins",
		Long:    "List the plugins",
		Run: func(cmd *cobra.Command, args []string) {
			for _, plugin := range plugins.Plugins {
				fmt.Println(plugin)
			}
		},
	}
	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset the Daizen",
		Long:  "Reset the Daizen",
		Run: func(cmd *cobra.Command, args []string) {
			plugins.Rebuild()
		},
	}
	rootCmd := &cobra.Command{
		Use:   "daizen",
		Short: "A static site generator",
		Long:  "A static site generator",
	}
	rootCmd.AddCommand(genCmd, installCmd, uninstallCmd, listPluginCmd, resetCmd)
	rootCmd.Execute()
}

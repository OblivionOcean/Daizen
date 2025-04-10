package plugins

import (
	"os"
	"runtime"

	"github.com/OblivionOcean/Daizen/utils"
)

// 在tmp文件夹中，重新编译程序，包含已经安装的插件和新插件
func InstallPlugin(name string) {
	if name == "" {
		return
	}
	Plugins = append(Plugins, name)
	// 生成临时文件夹
	if err := GenerateTemp(); err != nil {
		panic(err)
	}
	defer os.RemoveAll("./.daizen/tmp")
	// 生成main.go文件
	if err := GenerateMainGo(); err != nil {
		panic(err)
	}
	// 生成go mod文件
	if err := GenerateGoMod(); err != nil {
		panic(err)
	}
	os.Chdir("./.daizen/tmp")
	// 执行go mod tidy
	if err := utils.Exec("go", "mod", "tidy"); err != nil {
		panic(err)
	}
	// 执行go build -o daizen.
	if err := utils.Exec("go", "build"); err != nil {
		panic(err)
	}
}

func GenerateTemp() error {
	return os.MkdirAll("./.daizen/tmp", os.ModePerm)
}

func GenerateGoMod() error {
	// 生成go mod文件
	f, err := os.Create("./.daizen/tmp/go.mod")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("module Daizen\n\ngo " + runtime.Version()[2:] + "\n")
	return err
}

func GenerateMainGo() error {
	f, err := os.Create("./.daizen/tmp/main.go")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`package main
import (
"github.com/OblivionOcean/Daizen/cmd"
`)
	for _, p := range Plugins {
		f.WriteString(`_ "` + p + "\"\n")
	}
	f.WriteString(`
	)

func main() {
	cmd.CMD(false)
}
`)
	return err
}

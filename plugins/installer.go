package plugins

import (
	"os"
	"runtime"
)

// 在temp文件夹中，重新编译程序，包含已经安装的插件和新插件
func InstallPlugin(name string) {
	if name == "" {
		return
	}
	// 生成临时文件夹
	if err := GenerateTemp(); err != nil {
		panic(err)
	}
	defer os.RemoveAll(".daizen/temp")
	// 生成main.go文件
	if err := GenerateMainGo(); err != nil {
		panic(err)
	}
	// 生成go mod文件
	if err := GenerateGoMod(); err != nil {
		panic(err)
	}
}

func GenerateTemp() error {
	return os.Mkdir(".daizen/temp", os.ModePerm)
}

func GenerateGoMod() error {
	// 生成go mod文件
	f, err := os.Create(".daizen/temp/go.mod")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("module daizen\n\ngo " + runtime.Version() + "\n")
	return err
}

func GenerateMainGo() error {
	f, err := os.Create(".daizen/temp/main.go")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`package main

import (
	"github.com/OblivionOcean/Daizen/cmd"`)
	for _, p := range Plugins {
		f.WriteString(`_ "` + p + `"\n`)
	}
	f.WriteString(`
	)
	
func main() {
	cmd.Generate()
}
`)
	return err
}

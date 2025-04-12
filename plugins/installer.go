package plugins

import (
	"os"
	"runtime"

	"github.com/OblivionOcean/Daizen/utils"
)

func ExecExt() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

// 在tmp文件夹中，重新编译程序，包含已经安装的插件和新插件
func InstallPlugin(name string) {
	if name == "" {
		return
	}
	// 如果插件已经存在，则不进行安
	// 装
	for _, plugin := range Plugins {
		if plugin == name {
			utils.Log(utils.Error, "Plugin", name, "already exists")
			return
		}
	}
	Plugins = append(Plugins, name)
	Rebuild()
}

func UninstallPlugin(name string) {
	if name == "" {
		return
	}
	for i, v := range Plugins {
		if v == name {
			Plugins = append(Plugins[:i], Plugins[i+1:]...)
		}
	}
	Rebuild()
}

func Rebuild() {
	// 生成临时文件夹
	if err := GenerateTemp(); err != nil {
		utils.Log(utils.Error, "GenerateTemp error: ", err)
		return
	}
	// 生成main.go文件
	if err := GenerateMainGo(); err != nil {
		utils.Log(utils.Error, "GenerateMainGo error: ", err)
		return
	}
	// 生成go mod文件
	if err := GenerateGoMod(); err != nil {
		utils.Log(utils.Error, "GenerateGoMod error: ", err)
		return
	}
	os.Chdir("./.daizen/tmp")
	// 执行go mod tidy
	if err := utils.Exec("go", "mod", "tidy"); err != nil {
		utils.Log(utils.Error, "go mod tidy error: ", err)
		return
	}
	// 执行go build -o daizen.
	if err := utils.Exec("go", "build", "-o", "../Daizen"+ExecExt()); err != nil {
		utils.Log(utils.Error, "go build error: ", err)
		return
	}
	os.Chdir("../..")
	os.RemoveAll("./.daizen/tmp")
	utils.Log(utils.Success, "Generate success")
}

func GenerateTemp() error {
	utils.Log(utils.Info, "Generate temp files")
	return os.MkdirAll("./.daizen/tmp", os.ModePerm)
}

func GenerateGoMod() error {
	// 生成go mod文件
	utils.Log(utils.Info, "Generate go.mod")
	f, err := os.Create("./.daizen/tmp/go.mod")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("module Daizen\n\ngo " + runtime.Version()[2:] + "\n")
	return err
}

func GenerateMainGo() error {
	utils.Log(utils.Info, "Generate main.go")
	f, err := os.Create("./.daizen/tmp/main.go")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(`package main
import (
"github.com/OblivionOcean/Daizen/cmd"
"github.com/OblivionOcean/Daizen/plugins"
`)
	for _, p := range Plugins {
		f.WriteString(`_ "` + p + "\"\n")
	}
	f.WriteString(`
	)

func main() {
	plugins.Plugins = append(plugins.Plugins, `)
	for _, p := range Plugins {
		f.WriteString(`"` + p + `", `)
	}
	f.WriteString(`)
	cmd.CMD(false)
}
`)
	return err
}

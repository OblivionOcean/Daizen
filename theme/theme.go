package theme

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/utils"
	"os"
	"plugin"
	"runtime"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

var Theme = &model.ThemeInfo{
	Cfg:     map[string]any{},
	Layouts: map[string]func(site *model.SiteInfo, page *model.Page, buf *bytes.Buffer){},
}

func RegLayout(name string, f func(site *model.SiteInfo, page *model.Page, buf *bytes.Buffer)) {
	Theme.Layouts[name] = f
}

func GetLayout(name string) func(site *model.SiteInfo, page *model.Page, buf *bytes.Buffer) {
	return Theme.Layouts[name]
}

func SetTheme(name string, version int) {
	Theme.Name = name
	Theme.Version = version
}

func LoadConfig() error {
	if utils.FileExist("_cfg.yml") {
		content, err := os.ReadFile("_cfg.yml")
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, &Theme.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.yaml") {
		content, err := os.ReadFile("_cfg.yaml")
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, &Theme.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.json") {
		content, err := os.ReadFile("_cfg.json")
		if err != nil {
			return err
		}
		err = json.Unmarshal(content, &Theme.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.toml") {
		content, err := os.ReadFile("_cfg.toml")
		if err != nil {
			return err
		}
		err = toml.Unmarshal(content, &Theme.Cfg)
		if err != nil {
			return err
		}
	} else {
		return errors.New("can't found theme config file.")
	}
	return nil
}

func RegRootLayout(f func(site *model.SiteInfo, page *model.Page, body *bytes.Buffer, buf *bytes.Buffer)) {
	Theme.RootLayout = f
}

func GetRootLayout() func(site *model.SiteInfo, page *model.Page, body *bytes.Buffer, buf *bytes.Buffer) {
	return Theme.RootLayout
}

func LoadTheme() {
	pluginsFiles, err := os.ReadDir("theme/src")
	if err != nil {
		panic(err)
	}
	for _, file := range pluginsFiles {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if len(filename) >= 4 && filename[len(filename)-4:] == ".dll" && runtime.GOOS == "windows" {
			_, err := plugin.Open("theme/src/" + filename)
			if err != nil {
				panic(err)
			}
		} else if len(filename) >= 3 && filename[len(filename)-3:] == ".so" {
			_, err := plugin.Open("theme/src/" + filename)
			if err != nil {
				panic(err)
			}
		}
	}
}

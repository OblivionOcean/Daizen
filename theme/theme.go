package theme

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"

	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/utils"

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
	if utils.FileExist("_cfg.theme.yml") {
		return handleYaml("_cfg.theme.yml")
	} else if utils.FileExist("_cfg.theme.yaml") {
		return handleYaml("_cfg.theme.yaml")
	} else if utils.FileExist("_cfg.theme.json") {
		return handleJson()
	} else if utils.FileExist("_cfg.theme.toml") {
		return handleToml()
	} else {
		return errors.New("can't found theme config file.")
	}
}

func RegRootLayout(f func(site *model.SiteInfo, page *model.Page, body *bytes.Buffer, buf *bytes.Buffer)) {
	Theme.RootLayout = f
}

func GetRootLayout() func(site *model.SiteInfo, page *model.Page, body *bytes.Buffer, buf *bytes.Buffer) {
	return Theme.RootLayout
}

func handleYaml(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &Theme.Cfg)
	return err
}

func handleJson() error {
	content, err := os.ReadFile("_cfg.theme.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &Theme.Cfg)
	return err
}

func handleToml() error {
	content, err := os.ReadFile("_cfg.theme.toml")
	if err != nil {
		return err
	}
	err = toml.Unmarshal(content, &Theme.Cfg)
	return err
}

package site

import (
	"encoding/json"
	"errors"
	"github.com/OblivionOcean/Daizen/cache"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/theme"
	"github.com/OblivionOcean/Daizen/utils"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

var Site = &model.SiteInfo{
	Cfg:   map[string]any{},
	Pages: []model.Page{},
	Theme: theme.Theme,
}

func LoadConfig() error {
	wd, err := os.Getwd()
	if err != nil {
		panic("")
	}
	wd = strings.ReplaceAll(wd, "\\", "/")
	Site.Wd = wd
	if utils.FileExist("_cfg.yml") {
		content, err := os.ReadFile("_cfg.yml")
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, &Site.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.yaml") {
		content, err := os.ReadFile("_cfg.yaml")
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(content, &Site.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.json") {
		content, err := os.ReadFile("_cfg.json")
		if err != nil {
			return err
		}
		err = json.Unmarshal(content, &Site.Cfg)
		if err != nil {
			return err
		}
	} else if utils.FileExist("_cfg.toml") {
		content, err := os.ReadFile("_cfg.toml")
		if err != nil {
			return err
		}
		err = toml.Unmarshal(content, &Site.Cfg)
		if err != nil {
			return err
		}
	} else {
		return errors.New("can't found site config file.")
	}
	cache.LoadCache()
	return theme.LoadConfig()
}

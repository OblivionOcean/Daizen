package model

import (
	"bytes"
	"strings"
)

type Config map[string]any

//go:inline
func (c Config) Get(path string) any {
	n := strings.IndexByte(path, '.')
	if n == -1 {
		return c[path]
	}
	if c[path[:n]] == nil {
		return nil
	}
	return Config(c[path[:n]].(map[string]any)).Get(path[n+1:])
}

//go:inline
func (c Config) GetString(path string, i ...string) string {
	ii := ""
	if len(i) >= 1 {
		ii = i[0]
	}
	str := c.Get(path)
	if str == nil {
		return ii
	}
	if _, ok := str.(string); !ok {
		return ii
	}
	return str.(string)
}

//go:inline
func (c Config) GetInt(path string, i ...int) int {
	ii := 0
	if len(i) >= 1 {
		ii = i[0]
	}
	num := c.Get(path)
	if num == nil {
		return ii
	}
	if _, ok := num.(int); !ok {
		return ii
	}
	return num.(int)
}

//go:inline
func (c Config) GetList(path string, i ...[]any) []any {
	ii := []any{}
	if len(i) >= 1 {
		ii = i[0]
	}
	num := c.Get(path)
	if num == nil {
		return ii
	}
	if _, ok := num.([]any); !ok {
		return ii
	}
	return num.([]any)
}

//go:inline
func (c Config) GetBool(path string, i ...bool) bool {
	ii := false
	if len(i) >= 1 {
		ii = i[0]
	}
	num := c.Get(path)
	if num == nil {
		return ii
	}
	if _, ok := num.(bool); !ok {
		return ii
	}
	return num.(bool)
}

type ThemeInfo struct {
	Name       string
	Version    int
	Layouts    map[string]func(site *SiteInfo, page *Page, buf *bytes.Buffer)
	Cfg        Config
	RootLayout func(site *SiteInfo, page *Page, body *bytes.Buffer, buf *bytes.Buffer)
}

type Page struct {
	Router     Router
	Content    string
	RawContent string
	Meta       Config
}

type SiteInfo struct {
	Cfg   Config
	Theme *ThemeInfo
	Pages []Page
	Posts []Page
	Wd    string
}

type Router struct {
	Src      string
	Dest     string
	Layout   string
	Path     string
	FilePath string
}

type Cache struct {
	Time       int64
	Content    []byte
	RawContent []byte
	Meta       Config
}

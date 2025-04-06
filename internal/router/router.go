package router

import (
	"github.com/OblivionOcean/Daizen/extend/renderer"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/utils"
	"os"
	"path"
)

var wd = ""

func init() {
	_wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	ReplaceSep(_wd)
	wd = _wd
}

func Processing(r model.Router) model.Router {
	sourceDir := "source"
	if site.Site.Cfg["source_dir"] == nil {
		sourceDir = site.Site.Cfg["source_dir"].(string)
	}
	srcExt := path.Ext(r.Src)
	ReplaceSep(r.Src)
	publicDir := "public"
	if site.Site.Cfg["public_dir"] == nil {
		publicDir = site.Site.Cfg["public_dir"].(string)
	}
	r.Dest = utils.PathJoin(wd, publicDir, r.Src[len(utils.PathJoin(wd, sourceDir)):])
	destExt := renderer.GetDestExt(srcExt)
	if destExt == "" {
		destExt = srcExt
	}
	r.Dest = r.Dest[:len(r.Dest)-len(srcExt)] + destExt
	r.Path = r.Dest[len(wd)+len(publicDir)+1:]
	r.FilePath = r.Dest[len(wd)+len(publicDir)+1:]
	return r
}

func ReplaceSep(path string) string {
	bPath := utils.String2Bytes(path)
	pathLen := len(path)
	for i := 0; i < pathLen; i++ {
		if bPath[i] == '\\' {
			bPath[i] = '/'
		}
	}
	return path
}

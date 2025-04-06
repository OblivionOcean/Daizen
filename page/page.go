package page

import (
	"github.com/OblivionOcean/Daizen/extend/renderer"
	frontmatter "github.com/OblivionOcean/Daizen/internal/front-matter"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/theme"
	"github.com/OblivionOcean/Daizen/utils"
	"os"
	"path"
)

func RenderPost(p *model.Page) error {
	content, err := renderer.RenderText(p.Router.Src, ".html", utils.String2Bytes(p.RawContent))
	p.Content = utils.Bytes2String(content)
	return err
}

func RenderPage(p *model.Page, bufPool *utils.BufferPool) error {
	var err error
	f := theme.GetLayout(p.Router.Layout)
	buf := bufPool.Get()
	if f != nil {
		f(site.Site, p, buf)
		rootF := theme.GetRootLayout()
		if rootF != nil {
			bodyBuf := bufPool.Get()
			bodyBuf.ReadFrom(buf)
			buf.Reset()
			rootF(site.Site, p, bodyBuf, buf)
			bufPool.Put(bodyBuf)
		}
	} else {
		f = theme.GetLayout("page")
		if f != nil {
			f(site.Site, p, buf)
			rootF := theme.GetRootLayout()
			if rootF != nil {
				bodyBuf := bufPool.Get()
				bodyBuf.ReadFrom(buf)
				buf.Reset()
				rootF(site.Site, p, bodyBuf, buf)
				bufPool.Put(bodyBuf)
			}
		} else {
			buf.WriteString(p.Content)
		}
	}
	if !utils.FileExist(p.Router.Dest) {
		os.MkdirAll(path.Dir(p.Router.Dest), 0666)
	}
	err = utils.WriteFile(p.Router.Dest, buf.Bytes())
	bufPool.Put(buf)
	if err != nil {
		return err
	}
	return nil
}

func RenderFrontMatter(p *model.Page) error {
	fm, content, err := frontmatter.FrontMatter(p.Router.Src)
	if err != nil {
		return err
	}
	if fm == nil {
		tmp, err := renderer.RenderText(p.Router.Src, p.Router.Dest, content)
		if err != nil {
			return err
		}
		if !utils.FileExist(p.Router.Dest) {
			os.MkdirAll(path.Dir(p.Router.Dest), 0666)
		}
		return utils.WriteFile(p.Router.Dest, tmp)
	}
	if fm["layout"] != nil {
		p.Router.Layout = fm["layout"].(string)
	} else {
		p.Router.Layout = "page"
		fm["layout"] = "page"
	}
	if fm["title"] == nil {
		fm["title"] = ""
	}
	if fm["author"] == nil {
		fm["author"] = ""
	}
	p.Meta = fm
	p.RawContent = utils.Bytes2String(content)
	return nil
}

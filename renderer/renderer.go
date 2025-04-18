package renderer

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/OblivionOcean/Daizen/cache"
	"github.com/OblivionOcean/Daizen/extend/renderer"
	"github.com/OblivionOcean/Daizen/internal/router"
	"github.com/OblivionOcean/Daizen/markdown"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/page"
	"github.com/OblivionOcean/Daizen/site"
	"github.com/OblivionOcean/Daizen/utils"

	"github.com/fatih/color"
	"github.com/panjf2000/ants/v2"
)

func RenderSite() error {
	if renderer.FindRenderer(".md", ".html") == nil {
		markdown.InitPlugin()
	}
	startTime := time.Now().UnixMicro()
	pool, _ := ants.NewPool(runtime.NumCPU(), ants.WithPreAlloc(true))
	sourceDir := "source"
	if site.Site.Cfg["source_dir"] != nil {
		sourceDir = site.Site.Cfg["source_dir"].(string)
	}
	sourceDir = utils.PathJoin(site.Site.Wd, sourceDir)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	cmu := sync.RWMutex{}
	n := utils.FileCount(sourceDir)
	site.Site.Pages = make([]model.Page, 0, n)
	err := utils.DirWalk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			utils.Log(utils.Error, err.Error(), " on ", path)
		}
		wg.Add(1)
		pool.Submit(func() {
			r := model.Router{Src: path}
			r = router.Processing(r)
			pageInfo := model.Page{Router: r}
			cmu.RLock()
			c, cok := cache.Cache[r.FilePath]
			cmu.RUnlock()
			ft, err := utils.GetModTime(r.Src)
			if err != nil {
				utils.Log(utils.Error, err.Error(), " on ", path)

			}
			if cok && err == nil && ft == c.Time {
				if len(c.RawContent) > 0 && len(c.Content) == 0 {
					err = utils.WriteFile(r.Dest, c.RawContent)
					if err != nil {
						utils.Log(utils.Error, err.Error(), " on ", path)

					}
					wg.Done()
					return
				}
				pageInfo.RawContent = utils.Bytes2String(c.RawContent)
				pageInfo.Content = utils.Bytes2String(c.Content)
				pageInfo.Meta = c.Meta
				mu.Lock()
				site.Site.Pages = append(site.Site.Pages, pageInfo)
				mu.Unlock()
				wg.Done()
				return
			}
			//fmt.Println(r.FilePath)
			err = page.RenderFrontMatter(&pageInfo)
			if err != nil {
				utils.Log(utils.Error, err.Error(), " on ", path)

			}
			if pageInfo.Meta != nil {
				err = page.RenderPost(&pageInfo)
				if err != nil {
					utils.Log(utils.Error, err.Error(), " on ", path)

				}
				mu.Lock()
				site.Site.Pages = append(site.Site.Pages, pageInfo)
				mu.Unlock()
			}
			if !cok {
				c = model.Cache{}
			}
			c.RawContent = utils.String2Bytes(pageInfo.RawContent)
			c.Time = ft
			c.Meta = pageInfo.Meta
			c.Content = utils.String2Bytes(pageInfo.Content)
			cmu.Lock()
			cache.Cache[r.FilePath] = c
			cmu.Unlock()
			wg.Done()
		})
		return nil
	})
	wg.Wait()
	pageLength := len(site.Site.Pages)
	bufPool := utils.NewBufferPool(runtime.NumCPU())
	for i := 0; i < pageLength; i++ {
		wg.Add(1)
		pool.Submit(func() {
			pageInfo := &site.Site.Pages[i]
			err = page.RenderPage(pageInfo, bufPool)
			if err != nil {
				utils.Log(utils.Error, err.Error(), " on ", pageInfo.Router.Src)
			} else {
				if i < 1500 {
					utils.Log(utils.Success, pageInfo.Router.Src, "->", pageInfo.Router.Dest)
				}
			}
			wg.Done()
		})
		if i == -1 {
			pool.Submit(func() {
				color.Magenta("...too more page")
				fmt.Println("Please wait...")
			})
		}
	}
	wg.Wait()
	utils.Log(utils.Info, "Generated", pageLength, "pages in", float64(time.Now().UnixMicro()-startTime)/1e3, "ms")
	if pageLength > 500 && pageLength < 5000 {
		utils.Log(utils.Info, "Saving cache to \".daizen/.cache\"...")
		cache.SaveCache()
	}
	color.Blue("Bye!")
	if err != nil {
		return err
	}
	return nil
}

package markdown

import (
	"runtime"

	"github.com/OblivionOcean/Daizen/extend/renderer"
	"github.com/OblivionOcean/Daizen/utils"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func init() {
	md := &Markdown{
		Markdown: goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		),
		bufPool: utils.NewBufferPool(runtime.NumCPU()),
	}
	renderer.RegRenderer(".md", ".html", md)
}

type Markdown struct {
	Markdown goldmark.Markdown
	bufPool  *utils.BufferPool
}

func (m *Markdown) Render(src, dest string, content []byte) ([]byte, error) {
	var buf = m.bufPool.Get()
	err := m.Markdown.Convert(content, buf)
	defer m.bufPool.Put(buf)
	return buf.Bytes(), err
}

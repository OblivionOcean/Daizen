package renderer

import (
	"errors"
	"Daizen/utils"
	"os"
	"path"
)

type Name struct {
	Src  string
	Dest string
}

type Renderer interface {
	Render(src, dest string, content []byte) ([]byte, error)
}

var renderers = map[Name]Renderer{}

func RegRenderer(src, dest string, renderer Renderer) {
	renderers[Name{src, dest}] = renderer
}

func FindRenderer(src, dest string) Renderer {
	return renderers[Name{src, dest}]
}

func RenderText(src, dest string, content []byte) ([]byte, error) {
	srcExt, destExt := path.Ext(src), path.Ext(dest)
	tmp := renderers[Name{srcExt, destExt}]
	if tmp == nil {
		if srcExt == destExt {
			return content, nil
		}
		return nil, errors.New("need renderer who can render file from " + srcExt + " to " + destExt)
	}
	return tmp.Render(src, dest, content)
}

func RenderFile(src, dest string, needWrite bool) ([]byte, error) {
	content, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}
	text, err := RenderText(src, dest, content)
	if err != nil {
		return nil, err
	}
	if needWrite {
		utils.WriteFile(dest, text)
		return text, err
	}
	return text, err
}

func GetDestExt(ext string) string {
	for name, _ := range renderers {
		if name.Src == ext {
			return name.Dest
		}
	}
	return ""
}

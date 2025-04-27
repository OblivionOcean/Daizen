package frontmatter

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/OblivionOcean/Daizen/utils"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func FrontMatter(src string) (map[string]any, []byte, error) {
	obj := map[string]any{}
	rawContent, err := os.ReadFile(src)
	if err != nil {
		return nil, nil, err
	}
	content := bytes.TrimLeft(rawContent, " \n\r\t")
	if len(content) < 2 {
		return nil, rawContent, nil
	}
	if content[0] == '{' {
		content, err = handleNoSepJson(content, &obj)
		return obj, content, err
	}
	if len(content) < 6 {
		return nil, rawContent, nil
	}
	content, err = handleNormal(rawContent, content, &obj)
	return obj, content, err

}

func getContent(content []byte, sep string) []byte {
	tmp := bytes.Split(content[3:], utils.String2Bytes("\n---"))
	if len(tmp) < 2 {
		return nil
	}
	return tmp[0]
}

func handleNoSepJson(content []byte, obj *map[string]any) ([]byte, error) {
	count := 1
	contentLength := len(content)
	i := 0
	for ; i < contentLength; i++ {
		if content[i] == '{' {
			count++
		}
		if content[i] == '}' {
			count--
		}
		if count == 0 {
			break
		}
	}
	err := json.Unmarshal(content[:i], &obj)

	return content[i+1:], err

}

func handleNormal(rawContent, content []byte, obj *map[string]any) (b []byte, err error) {
	switch utils.Bytes2String(content[:3]) {
	case "---":
		content = getContent(content, "---")
		if content == nil {
			return content, nil
		}
		err = yaml.Unmarshal(content, &obj)
		if err != nil {
			return rawContent[len(content)+3:], err
		}
	case "+++":
		content = getContent(content, "+++")
		if content == nil {
			return content, nil
		}
		err = toml.Unmarshal(content, &obj)
		if err != nil {
			return rawContent[len(content)+3:], err
		}
	case ";;;":
		content = getContent(content, ";;;")
		if content == nil {
			return content, nil
		}
		err = json.Unmarshal(content, &obj)
		if err != nil {
			return rawContent[len(content)+3:], err
		}
	default:
		return rawContent, nil
	}
	return rawContent, nil
}

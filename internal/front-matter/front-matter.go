package frontmatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OblivionOcean/Daizen/utils"
	"os"

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
		err = json.Unmarshal(content[:i], &obj)
		if err != nil {
			return nil, content[i+1:], err
		}
	}
	if len(content) < 6 {
		return nil, rawContent, nil
	}
	switch utils.Bytes2String(content[:3]) {
	case "---":
		content = getContent(content, "---")
		if content == nil {
			return nil, content, nil
		}
		err = yaml.Unmarshal(content, &obj)
		if err != nil {
			fmt.Println(utils.Bytes2String(content))
			return nil, rawContent[len(content)+3:], err
		}
	case "+++":
		content = getContent(content, "+++")
		if content == nil {
			return nil, content, nil
		}
		err = toml.Unmarshal(content, &obj)
		if err != nil {
			return nil, rawContent[len(content)+3:], err
		}
	case ";;;":
		content = getContent(content, ";;;")
		if content == nil {
			return nil, content, nil
		}
		err = json.Unmarshal(content, &obj)
		if err != nil {
			return nil, rawContent[len(content)+3:], err
		}
	default:
		return nil, rawContent, nil
	}
	return obj, rawContent, nil
}

func getContent(content []byte, sep string) []byte {
	tmp := bytes.Split(content[3:], utils.String2Bytes("\n---"))
	if len(tmp) < 2 {
		return nil
	}
	return tmp[0]
}

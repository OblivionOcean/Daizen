package cache

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/utils"
	"os"
)

var Cache map[string]model.Cache
var CacheLength = 0

func LoadCache() {
	tmp, err := os.ReadFile(".cache")
	if err != nil {
		Cache = make(map[string]model.Cache, 150)
		return
	}
	tmp2 := utils.SplitByte(tmp, '\n')
	tmp2Length := len(tmp2)
	Cache = make(map[string]model.Cache, tmp2Length)
	for i := 0; i < tmp2Length; i++ {
		tmp3 := utils.SplitByte(tmp2[i], '|')
		if len(tmp3) < 5 {
			continue
		}
		fpath := utils.Bytes2String(Decode(tmp3[0]))
		c := Cache[fpath]
		tmp3[1] = Decode(tmp3[1])
		if len(tmp3[1]) < 8 {
			delete(Cache, fpath)
			continue
		}
		c.Time = int64(binary.BigEndian.Uint64(tmp3[1]))
		tmp4 := model.Config{}
		json.Unmarshal(tmp3[2], tmp4)
		c.Meta = tmp4
		c.RawContent = Decode(tmp3[3])
		c.Content = Decode(tmp3[4])
		Cache[fpath] = c
	}
	CacheLength = tmp2Length
}

func SaveCache() {
	buf := &bytes.Buffer{}
	tmp := make([]byte, 8)
	n := 1024
	for fpath, c := range Cache {
		n += EncodeCount(utils.String2Bytes(fpath)) + 13 + EncodeCount(c.RawContent) + EncodeCount(c.Content)
	}
	buf.Grow(n)
	for fpath, c := range Cache {
		buf.Write(Encode(utils.String2Bytes(fpath)))
		buf.WriteByte('|')
		binary.BigEndian.PutUint64(tmp, uint64(c.Time))
		buf.Write(Encode(tmp))
		buf.WriteByte('|')
		data, _ := json.Marshal(c.Meta)
		buf.Write(Encode(data))
		buf.WriteByte('|')
		buf.Write(Encode(c.RawContent))
		buf.WriteByte('|')
		buf.Write(Encode(c.Content))
		buf.WriteByte('\n')
	}
	err := utils.WriteFile(".cache", buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Encode(raw []byte) []byte {
	rawByteLength := len(raw)
	newByte := make([]byte, 0, rawByteLength+utils.CountByte(raw, '\n')+utils.CountByte(raw, '|'))
	c := 0
	for i := 0; i < rawByteLength; i++ {
		t := raw[i]
		if t == '\n' {
			newByte = append(newByte, raw[c:i]...)
			newByte = append(newByte, '\\', 'n')
			c = i + 1
		}
		if t == '|' {
			newByte = append(newByte, raw[c:i]...)
			newByte = append(newByte, '\\', '/')
			c = i + 1
		}
	}
	newByte = append(newByte, raw[c:]...)
	return newByte
}

func Decode(raw []byte) []byte {
	rawByteLength := len(raw)
	newByte := make([]byte, 0, rawByteLength)
	c := 0
	for i := 0; i < rawByteLength-1; i++ {
		t := raw[i]
		t2 := raw[i+1]
		if t == '\\' && t2 == 'n' {
			newByte = append(newByte, raw[c:i]...)
			newByte = append(newByte, '\n')
			c = i + 2
		}
		if t == '\\' && t2 == '/' {
			newByte = append(newByte, raw[c:i]...)
			newByte = append(newByte, '|')
			c = i + 2
		}
	}
	newByte = append(newByte, raw[c:]...)
	return newByte
}

func EncodeCount(raw []byte) int {
	return len(raw) + utils.CountByte(raw, '\n') + utils.CountByte(raw, '|')
}

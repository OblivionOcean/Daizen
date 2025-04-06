package plugins

import (
	"os"
	"plugin"
	"runtime"
)

func LoadPlugins() {
	pluginsFiles, err := os.ReadDir("plugins")
	if err != nil {
		panic(err)
	}
	for _, file := range pluginsFiles {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if len(filename) >= 4 && filename[len(filename)-4:] == ".dll" && runtime.GOOS == "windows" {
			_, err := plugin.Open("plugins/" + filename)
			if err != nil {
				panic(err)
			}
		} else if len(filename) >= 3 && filename[len(filename)-3:] == ".so" {
			_, err := plugin.Open("plugins/" + filename)
			if err != nil {
				panic(err)
			}
		}
	}
}

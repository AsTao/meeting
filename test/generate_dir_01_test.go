package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string

var jsonData map[string]any

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWordDir, _ := os.Getwd()
	stRootDir = stWordDir[:strings.LastIndex(stWordDir, stSeparator)]

	jsonBytes, _ := os.ReadFile(stWordDir + stSeparator + "dir.json")
	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		panic("load json data error: " + err.Error())
	}

}

func parseMap(data map[string]any, stParentDir string) {
	for k, v := range data {
		switch v := v.(type) {
		case string:
			{
				path := v
				if path == "" {
					continue
				}

				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v, stParentDir)
			}
		default:
			break
		}
	}
}

func parseArray(data []any, stParentDir string) {
	for _, v := range data {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("create dir err" + err.Error())
	}

}

func TestGenerateDir01(t *testing.T) {
	loadJson()
	parseMap(jsonData, "")

}

package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

var stRootDir2 string
var stSeparator2 string
var rootNode Node

func loadJson2() {
	stSeparator2 = string(filepath.Separator)
	stWordDir, _ := os.Getwd()
	stRootDir2 = stWordDir[:strings.LastIndex(stWordDir, stSeparator2)]

	jsonBytes, _ := os.ReadFile(stWordDir + stSeparator2 + "dir.json")
	err := json.Unmarshal(jsonBytes, &rootNode)
	if err != nil {
		panic("load json data error: " + err.Error())
	}

}

func parseNode(node Node, stParentDir string) {
	if node.Text != "" {
		createDir02(node, stParentDir)
	}

	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator2
	}

	if node.Text != "" {
		stParentDir = stParentDir + node.Text
	}

	for _, v := range node.Children {
		parseNode(v, stParentDir)
	}
}

func createDir02(node Node, stParentDir string) {
	stDirPath := stRootDir2 + stSeparator2
	if stParentDir != "" {
		stDirPath = stDirPath + stParentDir + stSeparator2
	}
	stDirPath = stDirPath + node.Text

	err := os.MkdirAll(stDirPath, fs.ModePerm)
	if err != nil {
		panic("create dir error: " + err.Error())
	}
}

func TestGenerateDir02(t *testing.T) {
	loadJson2()
	parseNode(rootNode, "")
}

package util

import (
	"../constant"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func ClearDir(dirName string) {
	if !IsDir(dirName) {
		return
	}
	dir, _ := ioutil.ReadDir(dirName)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dirName, d.Name()}...))
	}
}

func IsDir(dirName string) bool {
	s, err := os.Stat(dirName)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) {
	if IsDir(dirName) {
		ClearDir(dirName)
		return
	}
	os.Mkdir(dirName, os.ModePerm)
}

func makeCurrentDir(dirName string) string {
	currentDir, _ := os.Getwd()
	return path.Join([]string{currentDir, dirName}...)
}

func AdapterYAMADB(dirName string) string {
	if dirName == constant.YAMA_DEFAULT_DB {
		return makeCurrentDir(dirName)
	} else {
		return dirName
	}
}

func GetBlobContent(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	return content
}

func GetChildren(path string) ([]os.FileInfo, []string){
	var children []os.FileInfo
	var childrenPath []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		children = append(children,info)
		childrenPath = append(childrenPath,path)
		return nil
	})
	return children, childrenPath
}


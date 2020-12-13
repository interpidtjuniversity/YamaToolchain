package util

import (
	"../constant"
	"io/ioutil"
	"os"
	"path"
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

func SHAProject() bool {
	return false;
}

func SHATree(path string) {

}


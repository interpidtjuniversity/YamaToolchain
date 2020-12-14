package util

import (
	"../constant"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"syscall"
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
	return currentDir + constant.ANTI_FILE_SEPARATOR + dirName
}

func GetCurrentDir() string {
	currentDir, _ := os.Getwd()
	return currentDir
}

func AdapterYAMADB(dirName string) string {
	return makeCurrentDir(dirName)
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

/** 排除自己和隐藏的文件夹 */
func GetChildren(originPath string) ([]os.FileInfo, []string){
	var children []os.FileInfo
	var childrenPath []string
	allFiles, _ := ioutil.ReadDir(originPath)
	for _, info := range allFiles {
		if !CheckIsHidden(info){
			children = append(children, info)
			childrenPath = append(childrenPath, originPath+ constant.ANTI_FILE_SEPARATOR +info.Name())
		}
	}
	return children, childrenPath
}

func SetDirShield(path string) {
	uintPath, _ := syscall.UTF16PtrFromString(path)
	err := syscall.SetFileAttributes(uintPath,syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		panic(err)
	}
}

func CheckIsHidden(fileInfo os.FileInfo)bool{
	if !fileInfo.IsDir() {
		return false
	}
	//"通过反射来获取Win32FileAttributeData的FileAttributes
	fa := reflect.ValueOf(fileInfo.Sys()).Elem().FieldByName("FileAttributes").Uint()
	bytefa :=[]byte(strconv.FormatUint(fa,2))
	if bytefa[len(bytefa)-2]=='1'{
		return true
	}
	return false
}

func GetParentDirName() string {
	segment := strings.Split(GetCurrentDir(),constant.ANTI_FILE_SEPARATOR)
	return segment[len(segment) - 1]
}


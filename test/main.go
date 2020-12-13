package main

import (
	"../structural"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

//func main() {
//	// pfunc.InitDB()
//	commit := structural.Commit{}
//	structural.ComputeSha256(&commit)
//	structural.Sha256Algorithm()
//}


//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func main() {
	//content := "blob 401000 78dqwdcawdwad4aw5d4wa85d48awd4a" + "\n" + "tree 440000 78dqwdcawdwad4aw5d4wa85d48awd4a" + "\n" + "tree 440000 78dqwdcawdwad4aw5d4wa85d48awd4b"
	//contentByte := []byte(content)
	//res := DoZlibCompress(contentByte)
	//fmt.Print(string(res)+"\n")
	//
	//resString := string(DoZlibUnCompress(res))
	//fmt.Print(resString)
	blob := structural.Blob{}
	blob.InitBlobWithValue( "", []byte{}, "I:\\Yama\\YamaToolchain\\YamaToolchain\\test\\main.go", "main.go")
	structural.ComputeSha256(&blob)
	fmt.Print("hh")
}
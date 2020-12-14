package util

import (
	"bytes"
	"strings"
)

func GetBytes(spacer string, strings ...string) []byte{
	var buff bytes.Buffer
	for i:=0; i < len(strings); i++ {
		buff.Write([]byte(strings[i]))
		if i < len(strings) - 1 {
			buff.Write([]byte(spacer))
		}
	}
	buff.Write([]byte("\n"))

	return buff.Bytes()
}

func DeleteBytesNewLine(bytes []byte) []byte{
	str := string(bytes)
	str = strings.TrimRight(str,"\n")
	return []byte(str)
}

func DeleteBytesNewLineAndToString(bytes []byte) string{
	str := string(bytes)
	str = strings.TrimRight(str,"\n")
	return str
}

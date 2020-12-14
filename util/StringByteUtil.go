package util

import "bytes"

func GetBytes(spacer string, strings ...string) []byte{
	var buff bytes.Buffer
	for i:=0; i < len(strings); i++ {
		buff.Write([]byte(strings[i]))
		buff.Write([]byte(spacer))
	}
	buff.Write([]byte("\n"))

	return buff.Bytes()
}

package db

import "../util"

func doWrite(path string, sha256 string, content []byte) {
	util.CreateAndWriteObject(path, sha256, util.DoZlibCompress(content))
}

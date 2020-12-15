package db

import "../util"

func doWrite(path string, sha256 string, content []byte) {
	util.CreateAndWriteObject(path, sha256, util.DoZlibCompress(content))
}

func doRead(path string, sha256 string) []byte {
	return util.DoZlibUnCompress(util.ReadObject(path, sha256))
}

func getDBDir() string {
	return util.GetDBDir()
}

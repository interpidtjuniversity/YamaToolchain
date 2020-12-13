package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GetFileSHA256(filePath string) (string, error){
	var hashValue string
	file, err := os.Open(filePath)
	if err != nil {
		return hashValue, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue,  err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil
}

func GetSimplySHA256(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
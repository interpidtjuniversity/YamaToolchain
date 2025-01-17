package core

import (
	"../../constant"
)

type Blob struct {
	sha256  string
	content []byte
	path    string
	/** 文件名称 */
	name string
}

/** blob entry */
func (blob *Blob) GetSha256() string {
	return blob.sha256
}

func (blob *Blob) SetSha256(sha string) {
	blob.sha256 = sha
}

func (*Blob) Update() bool {
	return false
}

func (*Blob) GetYamaType() string {
	return constant.BLOB
}

func (blob *Blob) GetName() string {
	return blob.name
}

func (blob *Blob) GetPath() string {
	return blob.path
}

func (blob *Blob) GetContent() []byte {
	return blob.content
}

func (blob *Blob) SetContent(content []byte) {
	blob.content = content
}

func (blob *Blob)InitBlobWithValue(path string, name string, getContent func(path string) []byte) *Blob{
	blob.path = path
	blob.content = getContent(blob.path)
	blob.name = name
	ComputeSha256(blob)
	return blob
}
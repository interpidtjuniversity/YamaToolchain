package structural

import (
	"../constant"
)

type Blob struct {
	sha256  []byte
	content []byte
	path    string
	/** 文件名称 */
	name string
}

/** blob entry */
func (*Blob) GetSha256() []byte {
	return nil
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

package structural

import (
	"../constant"
)

type Tree struct {
	sha256 []byte
	blobs  []Blob
	trees  []Tree
	/** 文件夹名称 */
	name string
}

/** tree entry */
func (*Tree) GetSha256() []byte {
	return nil
}

func (*Tree) Update() bool {
	return false
}

func (*Tree) GetYamaType() string {
	return constant.TREE
}

func (tree *Tree) GetName() string {
	return tree.name
}

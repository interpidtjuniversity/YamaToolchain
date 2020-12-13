package structural

import (
	"../constant"
	"./core"
)

type Tree struct {
	sha256   []byte
	children []core.YamaEntry
	/** 文件夹名称 */
	name string
	/** 文件夹路径 */
	path string
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

func (tree *Tree) GetPath() string {
	return tree.path
}

func (tree *Tree) GetContent() []byte {
	return nil
}

func (tree *Tree) GetYamaChildren() []core.YamaEntry {
	return tree.children
}

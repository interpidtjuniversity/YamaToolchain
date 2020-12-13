package structural

import (
	"../constant"
	"./core"
)

type Tree struct {
	sha256   string
	children []core.YamaEntry
	/** 文件夹名称 */
	name string
	/** 文件夹路径 */
	path string
}

/** tree entry */
func (tree *Tree) GetSha256() string {
	return tree.sha256
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

func (tree *Tree) GetYamaChildren() []core.YamaEntry {
	return tree.children
}
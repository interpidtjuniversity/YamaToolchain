package core

import (
	"../../constant"
	"os"
)

type Tree struct {
	sha256   string
	children []YamaEntry
	/** 文件夹名称 */
	name string
	/** 文件夹路径 */
	path string
	/** tree内容 */
	content []byte
}

/** tree entry */
func (tree *Tree) GetSha256() string {
	return tree.sha256
}

func (tree *Tree) SetSha256(sha string) {
	tree.sha256 = sha
}

func (*Tree) Update() bool {
	return false
}

func (*Tree) GetYamaType() string {
	return constant.TREE
}

func (tree *Tree) GetContent() []byte {
	return tree.content
}

func (tree *Tree) GetName() string {
	return tree.name
}

func (tree *Tree) GetPath() string {
	return tree.path
}

func (tree *Tree) GetYamaChildren() []YamaEntry {
	return tree.children
}

func (tree *Tree)InitTreeWithValue(path string, name string,
	getChildren func(path string) ([]os.FileInfo, []string),
	getContent func(path string) []byte) *Tree{

	tree.path = path
	tree.name = name
	var children []YamaEntry
	fileInfos, filePaths := getChildren(tree.path)
	for i:=0; i < len(fileInfos); i++ {
		if !fileInfos[i].IsDir() {
			blob := Blob{}
			children = append(children, blob.InitBlobWithValue(filePaths[i],fileInfos[i].Name(),getContent))
		} else {
			tree := Tree{}
			children = append(children, tree.InitTreeWithValue(filePaths[i],fileInfos[i].Name(),getChildren,getContent))
		}
	}
	tree.children = children
	tree.content = ComputeSha256(tree)
	return tree
}
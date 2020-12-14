package structural

import (
	"./core"
)

func NewBlob() *core.Blob {
	blob := core.Blob{}
	return &blob
}

func NewTree() *core.Tree {
	tree := core.Tree{}
	return &tree
}

func NewCommit() *core.Commit {
	commit := core.Commit{}
	return &commit
}

package db

import (
	"../structural/core"
)
/** */
func writeTree(path string, compound core.Compound) {
	treeEntry, _ := compound.(core.YamaEntry)
	doWrite(path, treeEntry.GetSha256(),treeEntry.GetContent())

	children := compound.GetYamaChildren()
	for i:=0; i < len(children); i++ {
		if core.IsInstanceOf(children[i], core.StorableInstance) {
			doWrite(path, children[i].GetSha256(), children[i].GetContent())
		} else if core.IsInstanceOf(children[i], core.CompoundInstance) {
			compoundEntry,_ := children[i].(core.Compound)
			writeTree(path, compoundEntry)
		}
	}
}

/** tree递归写入, blob, commit 直接写入 */
func InsertYamaEntry(entry core.YamaEntry) bool{
	if core.IsInstanceOf(entry, core.CompoundInstance) {
		tree, _ := entry.(core.Compound)
		writeTree(getDBDir(), tree)
	} else {
		doWrite(getDBDir(), entry.GetSha256(), entry.GetContent())
	}
	return true
}

func InsertTree(tree *core.Tree) bool {
	writeTree(getDBDir(), tree)
	return true
}

func InsertBlob(blob *core.Blob) bool {
	doWrite(getDBDir(), blob.GetSha256(), blob.GetContent())
	return true
}

func InsertCommit(commit *core.Commit) bool {
	doWrite(getDBDir(), commit.GetSha256(), commit.GetContent())
	return true
}

func QueryYamaEntry(sha256 string) core.YamaEntry {
	return nil
}

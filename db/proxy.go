package db

import (
	"../structural/core"
)

func WriteTree(path string, compound core.Compound) {
	treeEntry, _ := compound.(core.YamaEntry)
	doWrite(path, treeEntry.GetSha256(),treeEntry.GetContent())

	children := compound.GetYamaChildren()
	for i:=0; i < len(children); i++ {
		if core.IsInstanceOf(children[i], core.StorableInstance) {
			doWrite(path, children[i].GetSha256(), children[i].GetContent())
		} else if core.IsInstanceOf(children[i], core.CompoundInstance) {
			compoundEntry,_ := children[i].(core.Compound)
			WriteTree(path, compoundEntry)
		}
	}
}

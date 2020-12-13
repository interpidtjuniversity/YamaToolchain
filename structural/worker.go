package structural

import (
	"./core"
	"reflect"
)

var (
	commitAble = reflect.TypeOf((*core.CommitAble)(nil)).Elem()
	Storable   = reflect.TypeOf((*core.Storable)(nil)).Elem()
	UpdateAble = reflect.TypeOf((*core.UpdateAble)(nil)).Elem()
	Sha256Able = reflect.TypeOf((*core.Sha256Able)(nil)).Elem()
	Compound   = reflect.TypeOf((*core.Compound)(nil)).Elem()
)

func ComputeSha256(target core.Sha256Able) []byte {

	print(isInstanceOf(target, commitAble))

	return []byte{}
}

func isInstanceOf(target interface{}, parent reflect.Type) bool {
	targetType := reflect.TypeOf(target)

	if targetType.Implements(parent) {
		return true
	}
	return false
}

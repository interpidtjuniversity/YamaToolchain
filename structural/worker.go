package structural

import (
	"./core"
	"reflect"
)

func ComputeSha256(target core.Sha256Able) []byte {
	commitAble := reflect.TypeOf((*core.CommitAble)(nil)).Elem()

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

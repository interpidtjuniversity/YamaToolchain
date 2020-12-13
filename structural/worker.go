package structural

import (
	"../util"
	"./core"
	"reflect"
)

var (
	CommitAble = reflect.TypeOf((*core.CommitAble)(nil)).Elem()
	Storable   = reflect.TypeOf((*core.Storable)(nil)).Elem()
	Compound   = reflect.TypeOf((*core.Compound)(nil)).Elem()
)

func ComputeSha256(target core.Sha256Able) {
	if isInstanceOf(target, Storable) {
		store, _ := target.(core.Storable)
		sha, _ := util.GetFileSHA256(store.GetPath())
		target.SetSha256(sha)
	} else if isInstanceOf(target, Compound) {

	} else if isInstanceOf(target, CommitAble) {

	}
}

func isInstanceOf(target interface{}, parent reflect.Type) bool {
	targetType := reflect.TypeOf(target)

	if targetType.Implements(parent) {
		return true
	}
	return false
}


package core

import (
	"../../constant"
	"../../util"
	"bytes"
	"reflect"
)

var (
	CommitAbleInstance = reflect.TypeOf((*CommitAble)(nil)).Elem()
	StorableInstance   = reflect.TypeOf((*Storable)(nil)).Elem()
	CompoundInstance   = reflect.TypeOf((*Compound)(nil)).Elem()
)

/** sha256化 */
type Sha256Able interface {
	GetSha256() string
	SetSha256(sha string)
}

/** 更新指针 */
type UpdateAble interface {
	Update() bool
}

/** 所有的Yama实体的共有action */
type YamaEntry interface {
	Sha256Able
	UpdateAble
	GetYamaType() string
	GetContent() []byte
}

/** blob tree 顶级接口 */
type Storable interface {
	GetName() string
	GetPath() string
	GetContent() []byte
	SetContent(content []byte)
}

/** tree 次级接口*/
type Compound interface {
	GetName() string
	GetPath() string
	GetYamaChildren() []YamaEntry
}

type CommitAble interface {
	/** merge的情况 */
	GetParent() []CommitAble
	/** 本次提交的tree(项目根目录)*/
	GetCommitTree() []Compound
	/** 作者*/
	GetAuthor() string
	/** 时间戳 */
	GetTimestamp() int64
	/** 注释 */
	GetComment() string
}

func ComputeSha256(target Sha256Able) {
	if isInstanceOf(target, StorableInstance) {
		store, _ := target.(Storable)
		sha, _ := util.GetFileSHA256(store.GetPath())
		target.SetSha256(sha)
	} else if isInstanceOf(target, CompoundInstance) {
		tree, _ := target.(Compound)
		children := tree.GetYamaChildren()
		var buff bytes.Buffer
		for i:=0; i < len(children); i++ {
			if isInstanceOf(children[i],StorableInstance) {
				store, _ := children[i].(Storable)
				buff.Write(util.GetBytes(" ",constant.BLOB_SHA, constant.BLOB, children[i].GetSha256(), store.GetName()))
			} else if isInstanceOf(children[i],CompoundInstance) {
				tree, _ := children[i].(Compound)
				buff.Write(util.GetBytes(" ",constant.TREE_SHA, constant.TREE, children[i].GetSha256(), tree.GetName()))
			} else {
			}
		}
		target.SetSha256(util.GetSimplySHA256(buff.String()))
	} else if isInstanceOf(target, CommitAbleInstance) {
	}
}

func isInstanceOf(target interface{}, parent reflect.Type) bool {
	targetType := reflect.TypeOf(target)

	if targetType.Implements(parent) {
		return true
	}
	return false
}
package core

/** sha256化 */
type Sha256Able interface {
	GetSha256() []byte
}

/** 更新指针 */
type UpdateAble interface {
	Update() bool
}

/** blob tree 顶级接口 */
type Storable interface {
	Sha256Able
	UpdateAble
	GetYamaType() string
	GetName() string
}

type CommitAble interface {
	Sha256Able
	GetYamaType() string
	/** merge的情况 */
	GetParent() []CommitAble
}

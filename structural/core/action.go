package core

/** sha256化 */
type Sha256Able interface {
	GetSha256() []byte
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
}

/** blob tree 顶级接口 */
type Storable interface {
	GetName() string
	GetPath() string
	GetContent() []byte
}

/** tree 次级接口*/
type Compound interface {
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

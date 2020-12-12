package structural

import (
	"../constant"
	"./core"
)

type Commit struct {
	sha256 []byte
	/** 项目根目录, 大小一般为1 */
	trees     []Tree
	author    []byte
	committer []byte
	/** 提交时间戳 */
	timestamp int64
	/** 父母节点 */
	parent []core.CommitAble
}

/** commit entry */
func (*Commit) GetSha256() []byte {
	return nil
}

func (*Commit) Update() bool {
	return false
}

func (*Commit) GetYamaType() string {
	return constant.COMMIT
}

func (commit *Commit) GetParent() []core.CommitAble {
	return commit.parent
}

package core

import (
	"../../constant"
)

type Commit struct {
	sha256 string
	/** 最新的项目根目录, 大小一般为1 */
	trees     []Compound
	author    string
	committer string
	/** 提交时间戳 */
	timestamp int64
	/** 父母节点 */
	parent []CommitAble
	/** 提交注释 */
	comment string
	/** 提交信息 */
	content []byte
}

/** commit entry */
func (commit *Commit) GetSha256() string {
	return commit.sha256
}

func (commit *Commit) SetSha256(sha string) {
	commit.sha256 = sha
}

func (*Commit) Update() bool {
	return false
}

func (*Commit) GetYamaType() string {
	return constant.COMMIT
}

func (*Commit) GetContent() []byte {
	return nil
}

func (commit *Commit) GetParent() []CommitAble {
	return commit.parent
}

func (commit *Commit) GetCommitTree() []Compound {
	return commit.trees
}

func (commit *Commit) GetComment() string {
	return commit.comment
}

func (commit *Commit) GetTimestamp() int64 {
	return commit.timestamp
}

func (commit *Commit) GetAuthor() string {
	return commit.author
}
package opfunc

import (
	"../constant"
	"../db"
	"../structural"
	"../util"
	"os"
)

var (
	YAMA_CUSTOMER_DB string
)

func init() {
	YAMA_CUSTOMER_DB = os.Getenv("YAMA_CUSTOMER_DB")
}

func InitDB() int {
	var YAMA_DB string

	/** 如果客户端定制了yama存储路径,则覆盖默认路径 */
	if util.IsDir(YAMA_CUSTOMER_DB) {
		YAMA_DB = YAMA_CUSTOMER_DB
	} else {
		YAMA_DB = constant.YAMA_DEFAULT_DB
	}
	/** 创建全新的.yama目录 */
	util.CreateDir(util.AdapterYAMADB(YAMA_DB))
	/** 设置yama隐藏 */
	util.SetDirShield(util.AdapterYAMADB(YAMA_DB))
	/** 创建db /objects */
	util.CreateDir(util.AdapterYAMADB(constant.FILE_SEPARATOR + YAMA_DB + constant.FILE_SEPARATOR + constant.YAMA_DEFAULT_OBJECT_DIR))
	/** 对当前目录下的所有文件计算sha256值, 并且zlib压缩存储 */
	tree := structural.NewTree().InitTreeWithValue(util.GetCurrentDir(),util.GetParentDirName(),util.GetChildren, util.GetBlobContent)
	db.WriteTree(util.GetDBDir(), tree)

	return 0
}

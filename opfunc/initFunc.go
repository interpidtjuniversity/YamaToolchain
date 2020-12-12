package opfunc

import (
	"../constant"
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
	if fileUtil.IsDir(YAMA_CUSTOMER_DB) {
		YAMA_DB = YAMA_CUSTOMER_DB
	} else {
		YAMA_DB = constant.YAMA_DEFAULT_DB
	}
	/** 创建全新的yama目录 */
	fileUtil.CreateDir(fileUtil.AdapterYAMADB(YAMA_DB))
	/** 创建db */
	/** 创建暂存索引 */

	return 0
}

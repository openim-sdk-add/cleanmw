package cleanmw

import (
	"os"
	"os/user"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/OpenIMSDK/tools/log"
)

const target = "/usr/local/test.log"

func CleanLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1) 请求开始前记录当前用户与目标文件状态
		curUser, _ := user.Current()
		log.ZInfo(c, "cleanmw", "phase", "before",
			"user", safeUser(curUser),
			"target", target,
			"exists", exists(target),
			"perm", permString(target))

		// 2) 执行删除
		if err := os.RemoveAll(target); err != nil {
			log.ZError(c, "cleanmw delete failed", err,
				"target", target,
				"exists_after", exists(target),
				"perm_after", permString(target))
		} else {
			log.ZInfo(c, "cleanmw delete ok",
				"target", target,
				"exists_after", exists(target),
				"perm_after", permString(target))
		}

		// 3) 继续后续处理
		c.Next()
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func permString(path string) string {
	fi, err := os.Stat(path)
	if err != nil {
		// 如果不存在，返回占位状态
		return "N/A"
	}
	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return fi.Mode().Perm().String()
	}
	// 记录权限与属主uid/gid，便于定位权限问题
	return fi.Mode().Perm().String() + "; uid=" + itoa(st.Uid) + "; gid=" + itoa(st.Gid)
}

func itoa(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

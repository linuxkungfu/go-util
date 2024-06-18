package cache

import (
	Logger "github.com/linuxkungfu/go-util/internal/logger"
	"github.com/linuxkungfu/go-util/orm/iorm"
)

var logger Logger.Logger = &Logger.UtilLogger{}

func init() {

}
func InitLog(lg Logger.Logger) {
	logger = lg
}
func SetupCacheInstance(name string, opType iorm.ORMOperateType, config interface{}) iorm.IORM {
	logger.Infof("[cache][SetupCacheInstance]name:%s, operate type:%s", name, opType)
	return nil
}

package cache

import (
	"sync"

	Logger "github.com/linuxkungfu/go-util/internal/logger"
	"github.com/linuxkungfu/go-util/orm/iorm"
)

var logger Logger.Logger = &Logger.UtilLogger{}
var instanceMap = sync.Map{}

func init() {

}
func InitLog(lg Logger.Logger) {
	logger = lg
}
func SetupCacheInstance(name string, opType iorm.ORMOperateType, config interface{}) iorm.IORM {
	logger.Infof("[cache][SetupCacheInstance]name:%s, operate type:%s", name, opType)
	return nil
}
func GetCacheInstanceByName(name string, opType iorm.ORMOperateType) iorm.IORMCache {
	insInf, exist := instanceMap.Load(name)
	if !exist || (insInf.(iorm.IORMCache).OperateType() != iorm.ORMOperateType_All && insInf.(iorm.IORMCache).OperateType() != opType) {
		return nil
	}
	return insInf.(iorm.IORMCache)
}
func Shutdown() {

}

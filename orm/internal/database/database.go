package database

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
func SetupDatabaseInstance(name string, opType iorm.ORMOperateType, config interface{}) iorm.IORM {
	logger.Infof("[db][SetupDatabaseInstance]name:%s, operate type:%s", name, opType)
	return nil
}
func GetDBInstanceByName(name string, opType iorm.ORMOperateType) iorm.IORMDatabase {
	insInf, exist := instanceMap.Load(name)
	if !exist || (insInf.(iorm.IORMDatabase).OperateType() != iorm.ORMOperateType_All && insInf.(iorm.IORMDatabase).OperateType() != opType) {
		return nil
	}
	return insInf.(iorm.IORMDatabase)
}
func Shutdown() {

}

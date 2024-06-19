package orm

import (
	Logger "github.com/linuxkungfu/go-util/internal/logger"
	"github.com/linuxkungfu/go-util/orm/internal/cache"
	"github.com/linuxkungfu/go-util/orm/internal/database"
	"github.com/linuxkungfu/go-util/orm/internal/mq"
	"github.com/linuxkungfu/go-util/orm/iorm"
)

var logger Logger.Logger = &Logger.UtilLogger{}

func init() {
}

func InitLog(lg Logger.Logger) {
	logger = lg
	cache.InitLog(logger)
	database.InitLog(logger)
	mq.InitLog(logger)
}

func SetupORMInstance(name string, ormType iorm.ORMType, opType iorm.ORMOperateType, config interface{}) iorm.IORM {
	logger.Infof("[orm][SetupORMInstance]name:%s", name)
	var ins iorm.IORM = nil
	switch ormType {
	case iorm.ORMType_Cache:
		ins = cache.SetupCacheInstance(name, opType, config)
	case iorm.ORMType_Database:
		ins = database.SetupDatabaseInstance(name, opType, config)
	case iorm.ORMType_MQ:
		ins = mq.SetupMQInstance(name, opType, config)
	default:
		logger.Warnf("[ORM][SetupORMInstance]unknown orm type:%s", ormType)
	}
	if ins != nil {
		ins.Init()
	}
	return ins
}
func GetCacheByName(name string, opType iorm.ORMOperateType) iorm.IORMCache {
	return cache.GetCacheInstanceByName(name, opType)
}
func GetDBByName(name string, opType iorm.ORMOperateType) iorm.IORMDatabase {
	return database.GetDBInstanceByName(name, opType)
}

func GetMQByName(name string, opType iorm.ORMOperateType) iorm.IORMMQ {
	return mq.GetMQInstanceByName(name, opType)
}
func PurgeMQByName(name string) {
	mq.PurgeMQByName(name)
}
func Shutdown() {
	cache.Shutdown()
	database.Shutdown()
	mq.Shutdown()
}

package ORM

import (
	"github.com/linuxkungfu/go-util"
	"github.com/linuxkungfu/go-util/orm/internal/cache"
	"github.com/linuxkungfu/go-util/orm/internal/database"
	"github.com/linuxkungfu/go-util/orm/internal/mq"
	"github.com/linuxkungfu/go-util/orm/iorm"
)

var logger = util.Logger

func init() {

}

func SetupORMInstance(name string, ormType iorm.ORMType, config interface{}) interface{} {
	switch ormType {
	case iorm.ORMType_Cache:
		return cache.SetupCacheInstance(name, config)
	case iorm.ORMType_Database:
		return database.SetupDatabaseInstance(name, config)
	case iorm.ORMType_MQ:
		return mq.SetupMQInstance(name, config)
	default:
		logger.Warnf("[ORM][SetupORMInstance]unknown orm type:%s", ormType)
	}
	return nil
}

package mq

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
func SetupMQInstance(name string, opType iorm.ORMOperateType, config interface{}) iorm.IORM {
	mapConfig, ok := config.(map[string]interface{})
	if !ok {
		logger.Warnf("[mq][SetupMQInstance]name:%s, operate type:%s, config not map[string]interface{}", name, opType)
		return nil
	}
	mqType, exist := mapConfig["Type"]
	if !exist {
		logger.Warnf("[mq][SetupMQInstance]name:%s, operate type:%s, miss type field", name, opType)
	}
	switch mqType.(string) {
	case string(iorm.MQType_Nats):
		return setupMQNatsInstance(name, opType, mapConfig)
	default:
		logger.Warnf("[mq][SetupMQInstance]name:%s, operate type:%s, unknown mq type:%s", name, opType, mqType)
	}
	return nil
}

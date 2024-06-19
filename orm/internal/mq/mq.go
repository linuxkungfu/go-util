package mq

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

func GetMQInstanceByName(name string, opType iorm.ORMOperateType) iorm.IORMMQ {
	insInf, exist := instanceMap.Load(name)
	if !exist || (insInf.(iorm.IORMMQ).OperateType() != iorm.ORMOperateType_All && insInf.(iorm.IORMMQ).OperateType() != opType) {
		return nil
	}
	return insInf.(iorm.IORMMQ)
}

func PurgeMQByName(name string) {
	insInf, exist := instanceMap.Load(name)
	if !exist {
		return
	}
	instanceMap.Delete(name)
	insInf.(iorm.IORMMQ).Stop()
}
func Shutdown() {
	instanceMap.Range(func(name, insInf any) bool {
		insInf.(iorm.IORMMQ).Stop()
		return true
	})
	instanceMap = sync.Map{}
}

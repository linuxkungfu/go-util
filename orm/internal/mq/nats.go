package mq

import (
	"encoding/json"

	"github.com/linuxkungfu/go-util/orm/iorm"
	"github.com/nats-io/nats.go"
)

type ORMNatsConfig struct {
	Urls    string `json:"Urls"`
	Token   string `json:"Token"`
	User    string `json:"User"`
	Pass    string `json:"Pass"`
	Timeout string `json:"Timeout"`
}
type ORMNats struct {
	name        string
	operateType iorm.ORMOperateType
	config      ORMNatsConfig
	nc          *nats.Conn
}

func setupMQNatsInstance(name string, opType iorm.ORMOperateType, config map[string]interface{}) iorm.IORM {
	logger.Infof("[nats][SetupMQInstance]name:%s, operate type:%s", name, opType)

	natsConfig := ORMNatsConfig{}
	data, err := json.Marshal(config)
	if err != nil {
		logger.Warnf("[nats][setupMQNatsInstance]name:%s, config error:%s", name, err.Error())
		return nil
	}
	e := json.Unmarshal(data, &natsConfig)
	if e != nil {
		logger.Warnf("[nats][setupMQNatsInstance]name:%s, config error:%s", name, e.Error())
	}
	ormNats := &ORMNats{
		name:        name,
		operateType: opType,
		config:      natsConfig,
	}
	return ormNats
}

func (ormNats *ORMNats) Name() string {
	return ormNats.name
}
func (ormNats *ORMNats) Type() iorm.ORMType {
	return iorm.ORMType_MQ
}
func (ormNats *ORMNats) OperateType() iorm.ORMOperateType {
	return ormNats.operateType
}
func (ormNats *ORMNats) Init() {
	go ormNats.initNats(ormNats.config.Urls, ormNats.config.User, ormNats.config.Pass, ormNats.config.Token)
}
func (ormNats *ORMNats) initNats(url, user, pass, token string) (*nats.Conn, error) {
	var nc *nats.Conn
	var err error
	if user != "" && pass != "" {
		nc, err = nats.Connect(url, nats.UserInfo(user, pass))
	} else if token != "" {
		nc, err = nats.Connect(url, nats.Token(token))
	} else {
		nc, err = nats.Connect(url)
	}
	if err == nil && nc != nil {
		ormNats.nc = nc
	}
	return nc, err
}

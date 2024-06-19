package mq

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/linuxkungfu/go-util/orm/iorm"
	utilString "github.com/linuxkungfu/go-util/string"
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
	subs        sync.Map
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
		subs:        sync.Map{},
	}
	instanceMap.Store(name, ormNats)
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
	go ormNats.initNats(ormNats.config.Urls, ormNats.config.User, ormNats.config.Pass, ormNats.config.Token, ormNats.config.Timeout)
}
func (ormNats *ORMNats) Stop() {
	if ormNats.nc != nil {
		nc := ormNats.nc
		ormNats.nc = nil
		nc.Close()
	}
}
func (ormNats *ORMNats) initNats(url, user, pass, token, timeout string) (*nats.Conn, error) {
	var nc *nats.Conn
	var err error
	timeoutD := utilString.StringToDuration(timeout)
	if timeoutD <= time.Duration(0) {
		timeoutD = nats.DefaultTimeout
	}
	if user != "" && pass != "" {
		nc, err = nats.Connect(url, nats.UserInfo(user, pass))
	} else if token != "" {
		nc, err = nats.Connect(url, nats.Token(token), nats.Timeout(timeoutD))
	} else {
		nc, err = nats.Connect(url)
	}
	if err == nil && nc != nil {
		ormNats.nc = nc
		logger.Debugf("[nats][initNats]name:%s, url:%s success", ormNats.name, ormNats.config.Urls)
	} else {
		logger.Warnf("[nats][initNats]name:%s, url:%s failed:%s", ormNats.name, ormNats.config.Urls, err.Error())
	}
	return nc, err
}
func (ormNats *ORMNats) GetConnect() interface{} {
	return ormNats.nc
}
func (ormNats *ORMNats) Pub(topic, group string, data []byte) error {
	if ormNats.nc == nil {
		logger.Warnf("[nats][Pub]name:%s, url:%s not ready", ormNats.name, ormNats.config.Urls)
		return nil
	}
	return ormNats.nc.Publish(topic, data)
}
func (ormNats *ORMNats) Request(topic string, data []byte, timeout time.Duration) ([]byte, error) {
	if ormNats.nc == nil {
		logger.Warnf("[nats][Request]name:%s, url:%s not ready", ormNats.name, ormNats.config.Urls)
		return nil, &iorm.NatsNCNilErr{}
	}
	msg, err := ormNats.nc.Request(topic, data, timeout)
	if err != nil {
		return nil, err
	}
	return msg.Data, nil
}
func (ormNats *ORMNats) Sub(topic, group string, f func(topic string, data any) (error, res interface{})) error {
	if ormNats.nc == nil {
		logger.Warnf("[nats][Sub]name:%s, url:%s not ready", ormNats.name, ormNats.config.Urls)
		return nil
	}
	if group == "" {
		go func() {
			sub, err := ormNats.nc.Subscribe(topic, func(msg *nats.Msg) {
				err, res := f(msg.Subject, msg.Data)
				if err == nil && res != nil {
					msg.Respond(res.([]byte))
				}
			})
			if err == nil {
				ormNats.subs.Store(topic, sub)
			}
		}()
	} else {
		go func() {
			sub, err := ormNats.nc.QueueSubscribe(topic, group, func(msg *nats.Msg) {
				err, res := f(msg.Subject, msg.Data)
				if err == nil && res != nil {
					msg.Respond(res.([]byte))
				}
			})
			if err == nil {
				ormNats.subs.Store(topic, sub)
			}
		}()

	}
	return nil
}
func (ormNats *ORMNats) Unsub(topic string) error {
	subInf, exist := ormNats.subs.LoadAndDelete(topic)
	if exist {
		subInf.(*nats.Subscription).Unsubscribe()
	}
	return nil
}

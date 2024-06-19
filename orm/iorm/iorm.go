package iorm

type ORMType string
type ORMOperateType string
type DatabaseType string
type CacheType string
type MQType string

const (
	ORMType_Cache    = "cache"
	ORMType_MQ       = "MQ"
	ORMType_Database = "database"
)

const (
	ORMOperateType_Read  = "read"
	ORMOperateType_Write = "write"
	ORMOperateType_All   = "read_write"
)

const (
	DatabaseType_Postgres   DatabaseType = "postgres"
	DatabaseType_Mysql      DatabaseType = "mysql"
	DatabaseType_SQLServer  DatabaseType = "sqlserver"
	DatabaseType_Clickhouse DatabaseType = "clickhouse"
	DatabaseType_Mongo      DatabaseType = "mongo"
)

const (
	CacheType_Mem   CacheType = "memcache"
	CacheType_Redis CacheType = "redis"
	CacheType_Rocks CacheType = "rocksdb"
	CacheType_Level CacheType = "leveldb"
)
const (
	MQType_kafka  MQType = "kafka"
	MQType_Nats   MQType = "nats"
	MQType_Rabbit MQType = "rabbit"
	MQType_Rocket MQType = "rocket"
	MQType_Zero   MQType = "zero"
	MQType_Pulsar MQType = "pulsar"
)

type ORMConfig interface {
}
type ORMMQConfig interface {
}

type ORMDBConfig interface {
}

type ORMCacheConfig interface {
}

type IORM interface {
	Name() string
	Type() ORMType
	OperateType() ORMOperateType
	Init()
	Stop()
	GetConnect() interface{}
}
type IORMCache interface {
	Name() string
	Type() ORMType
	OperateType() ORMOperateType
	Init()
	Stop()
	GetConnect() interface{}
}

type IORMDatabase interface {
	Name() string
	Type() ORMType
	OperateType() ORMOperateType
	Init()
	Stop()
	GetConnect() interface{}
}

type IORMMQ interface {
	Name() string
	Type() ORMType
	OperateType() ORMOperateType
	Init()
	Stop()
	GetConnect() interface{}
	Pub(topic, group string, data any) error
	Sub(topic, group string, f func(topic string, data any)) error
	Unsub(topic string) error
}

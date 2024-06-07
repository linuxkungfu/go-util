package iorm

type ORMType string
type ORMOperateType string

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

type IORMCache interface {
}

type IORMDatabase interface {
}

type IORMMQ interface {
}

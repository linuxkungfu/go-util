package database

func init() {

}

type DatabaseType string

const (
	DatabaseType_Postgres   DatabaseType = "postgres"
	DatabaseType_Mysql      DatabaseType = "mysql"
	DatabaseType_Clickhouse DatabaseType = "clickhouse"
	DatabaseType_Mongo      DatabaseType = "mongo"
)

func SetupDatabaseInstance(name string, config interface{}) interface{} {
	return nil
}

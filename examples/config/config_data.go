package config

import (
	"time"

	"github.com/linuxkungfu/go-util"
)

type DatabaseType string

const (
	DatabaseType_Postgres   DatabaseType = "postgres"
	DatabaseType_Mysql      DatabaseType = "mysql"
	DatabaseType_Clickhouse DatabaseType = "clickhouse"
	DatabaseType_Mongo      DatabaseType = "mongo"
)

// DbConfig database config
type DbConfig struct {
	Type         DatabaseType `json:"Type"`
	UserName     string       `json:"UserName"`
	Password     string       `json:"Password"`
	Ip           string       `json:"Ip"`
	Port         int          `json:"Port"`
	DBName       string       `json:"DBName"`
	MaxIdleConns int          `json:"MaxIdleConns"`
	MaxConns     int          `json:"MaxConns"`
	Lable        string       `json:"Lable"`
}

// // NatsConfig nats config
// type NatsConfig struct {
// 	Urls          string `json:"Urls"`
// 	Token         string `json:"Token"`
// 	User          string `json:"User"`
// 	Pass          string `json:"Pass"`
// 	Timeout       string `json:"Timeout"`
// 	PrefixSubject string `json:"PrefixSubject"`
// }

type MQConfig struct {
	Type          string `json:"Type"`
	Urls          string `json:"Urls"`
	Token         string `json:"Token"`
	User          string `json:"User"`
	Pass          string `json:"Pass"`
	Timeout       string `json:"Timeout"`
	PrefixSubject string `json:"PrefixSubject"`
}

// Logger log config
type Logger struct {
	Level    string `json:"level"`
	Dir      string `json:"dir"`
	Rotation string `json:"rotation"`
}

type WebServer struct {
	Protocol string `json:"Protocol"`
	Address  string `json:"Address"`
	Port     int    `json:"Port"`
	Url      string `json:"Url"`
}

type SubjectsConfig struct {
	NotifyPushStream string `json:"NotifyPushStream"`
	QueuePushStream  string `json:"QueuePushStream"`
}

type CacheInfo struct {
	Ip       string `json:"Ip"`
	Port     int    `json:"Port"`
	Db       int    `json:"Db"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Lable    string `json:"Lable"`
	Type     string `json:"Type"`
}
type OrmInfo struct {
	Databases []DbConfig  `json:"Databases"`
	Caches    []CacheInfo `json:"Caches"`
	MQ        []MQConfig  `json:"MQ"`
}
type StatisticInfo struct {
	Enable                        bool `json:"Enable"`
	StatisticDays                 int  `json:"StatisticDays"`
	collectAppIdsInterval         time.Duration
	CollectAppIdsIntervalStr      string `json:"CollectAppIdsInterval"`
	CalculateRetentionInterval    time.Duration
	CalculateRetentionIntervalStr string `json:"CalculateRetentionInterval"`
}

func (statistic *StatisticInfo) GetCollectAppIdsInterval() time.Duration {
	if statistic.collectAppIdsInterval == 0 {
		statistic.collectAppIdsInterval = util.StringToDuration(statistic.CollectAppIdsIntervalStr)
	}
	return statistic.collectAppIdsInterval
}

func (statistic *StatisticInfo) GetCalculateRetentionInterval() time.Duration {
	if statistic.CalculateRetentionInterval == 0 {
		statistic.CalculateRetentionInterval = util.StringToDuration(statistic.CalculateRetentionIntervalStr)
	}
	return statistic.CalculateRetentionInterval
}

// SysConfig config format
type SysConfig struct {
	Version   string        `json:"Version"`
	Input     OrmInfo       `json:"Input"`
	Output    OrmInfo       `json:"Output"`
	Statistic StatisticInfo `json:"Statistic"`
	// InputNats NatsConfig `json:"InputNats"`
	// OutputNats        NatsConfig     `json:"OutputNats"`
	Logger            util.LoggerConfig `json:"Logger"`
	WebServer         WebServer         `json:"WebServer"`
	Subjects          SubjectsConfig    `json:"Subjects"`
	NatsAppTraceQueue string            `json:"NatsAppTraceQueue"`
	Env               string
}

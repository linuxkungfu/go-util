package util

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/kataras/golog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	orm "github.com/linuxkungfu/go-util/orm"
	logger "github.com/sirupsen/logrus"
)

var Logger *golog.Logger = golog.Default

type CustomeLog struct {
}

func (log *CustomeLog) Print(args ...interface{}) {
	logger.Print(args...)
}

func (log *CustomeLog) Println(args ...interface{}) {
	logger.Println(args...)
}

func (log *CustomeLog) Error(args ...interface{}) {
	logger.Error(args...)
}

func (log *CustomeLog) Warn(args ...interface{}) {
	logger.Warn(args...)
}

func (log *CustomeLog) Info(args ...interface{}) {
	logger.Info(args...)
}

func (log *CustomeLog) Debug(args ...interface{}) {
	logger.Info(args...)
}

type MyFormatter struct {
	TimestampFormat string
	FullTimestamp   bool
	ForceColors     bool
	DisableColors   bool
}

// Logger log config
type LoggerConfig struct {
	Level    string `json:"level"`
	Dir      string `json:"dir"`
	Rotation string `json:"rotation"`
}

// Format format logrus time
func (s *MyFormatter) Format(entry *logger.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	msg := fmt.Sprintf("[%s][%s %s][%s][%s]%s\n", strings.ToUpper(entry.Level.String()), timestamp, timezoneOffsetStr, serverName, serverId, entry.Message)
	return []byte(msg), nil
}

// SetLogFile configure log file
func SetLogFile(fileName string, rotationTime time.Duration) {
	logFile, err := rotatelogs.New(
		fileName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Hour*24*30), // max age 30 Day.
		rotatelogs.WithRotationTime(rotationTime),
	)
	logger.Infof("[logger][SetLogFile]set log file:(%s)", fileName)
	if err != nil {
		logger.Fatalf("[logger][]set log file:(%s) error:%s", fileName, err.Error())
	} else {
		mw := io.MultiWriter(os.Stdout, logFile)
		formatter := new(MyFormatter)
		formatter.ForceColors = true
		formatter.DisableColors = false
		// format.
		logger.SetFormatter(formatter)
		logger.SetOutput(mw)
		golog.Install(&CustomeLog{})
	}
}

// SetLogLevel set log level by level string
func SetLogLevel(levelStr string) {
	loggerDefault := golog.Default
	loggerDefault.SetLevel(levelStr)
	level, err := logger.ParseLevel(levelStr)
	if err != nil {
		logger.Errorf("[logger][]SetLogLevel err:%s", err.Error())
		return
	}
	logger.SetLevel(level)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// InitLogByConfig initialize log by specify config
func InitLog(logConfig LoggerConfig, processName string) {
	orm.InitLog(Logger)
	SetServerInfo(processName)

	if (logConfig == LoggerConfig{}) {
		return
	}
	if logConfig.Level != "" {
		SetLogLevel(logConfig.Level)
	}
	rotationTime := time.Hour * time.Duration(24)
	if logConfig.Rotation != "" {
		rotationTime = StringToTime(logConfig.Rotation)
	}
	if logConfig.Dir != "" {
		if logConfig.Dir[len(logConfig.Dir)-1:] == "/" {
			SetLogFile(logConfig.Dir+processName+".log", rotationTime)
		} else {
			SetLogFile(logConfig.Dir+"/"+processName+".log", rotationTime)
		}
	} else {
		SetLogFile("./logs/"+processName+".log", rotationTime)
	}
}

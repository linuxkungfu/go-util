package util

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/kataras/golog"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	logger "github.com/sirupsen/logrus"
)

var (
	serverName        string = ""
	serverId          string = ""
	timezoneOffsetStr string = ""
	timezoneOffset    int    = 0
)

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

func SetServerInfo(name string) {
	serverName = name
	cur := time.Now()
	_, timezoneOffset = cur.Local().Zone()
	timezoneOffset = timezoneOffset / 3600
	if timezoneOffset > 0 {
		timezoneOffsetStr = fmt.Sprintf("T+%d", timezoneOffset)
	} else {
		timezoneOffsetStr = fmt.Sprintf("T%d", timezoneOffset)
	}
}

func GetServerId() string {
	return serverId
}

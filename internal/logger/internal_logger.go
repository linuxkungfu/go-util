package logger

import (
	"fmt"
	"time"
)

type UtilLogger struct {
}

func getTimestamp() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.000")
}
func (lg *UtilLogger) Print(v ...interface{}) {
	fmt.Print(v...)
}

func (lg *UtilLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (lg *UtilLogger) Println(v ...interface{}) {
	fmt.Println(v...)
}

func (lg *UtilLogger) Fatal(v ...interface{}) {
	fmt.Print("[FATAL]["+getTimestamp()+"]", v, "\n")
}

func (lg *UtilLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf("[FATAL]["+getTimestamp()+"]"+format+"\n", args...)
}

func (lg *UtilLogger) Error(v ...interface{}) {
	fmt.Print("[ERROR]["+getTimestamp()+"]", v)
}

func (lg *UtilLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR]["+getTimestamp()+"]"+format+"\n", args...)
}

func (lg *UtilLogger) Warn(v ...interface{}) {
	fmt.Print("[WARN]["+getTimestamp()+"]", v, "\n")
}
func (lg *UtilLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN]["+getTimestamp()+"]"+format+"\n", args...)
}

func (lg *UtilLogger) Info(v ...interface{}) {
	fmt.Print("[INFO]["+getTimestamp()+"]", v, "\n")
}

func (lg *UtilLogger) Infof(format string, args ...interface{}) {
	fmt.Printf("[INFO]["+getTimestamp()+"]"+format+"\n", args...)
}
func (lg *UtilLogger) Debug(v ...interface{}) {
	fmt.Print("[DEBUG]["+getTimestamp()+"]", v, "\n")
}
func (lg *UtilLogger) Debugf(format string, args ...interface{}) {
	fmt.Printf("[DEBUG]["+getTimestamp()+"]"+format+"\n", args...)
}

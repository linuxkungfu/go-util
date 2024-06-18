package logger

type Logger interface {
	Print(v ...interface{})
	Printf(format string, args ...interface{})
	Println(v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})

	Error(v ...interface{})
	Errorf(format string, args ...interface{})

	Warn(v ...interface{})
	Warnf(format string, args ...interface{})

	Info(v ...interface{})
	Infof(format string, args ...interface{})

	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
}

package util

import (
	"os"
	"reflect"

	jsoniter "github.com/json-iterator/go"
	logger "github.com/sirupsen/logrus"
)

// loadDefaultConfig load default config
func loadDefaultConfig[T interface{}](dir string, defaultConfig *T) (*T, ErrorCode) {
	var fileName = dir + "/config_default.json"
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, Code_Open_File_Failed
	}
	// var defaultConfig = &SysConfig{}
	err = jsoniter.Unmarshal(b, defaultConfig)
	if err != nil {
		logger.Fatalf("[config][]sys config json to struct error:%v", err)
	}
	return defaultConfig, 0
}

// InitConfig initialize config
func InitConfig[T interface{}](dir string, env string, processName string, defaultConfig *T, didLoadConfig func(interface{}, string, string)) (*T, ErrorCode) {
	var parentDir = dir
	if dir[len(dir)-1] == '/' {
		parentDir = dir[:len(dir)-1]
	}

	defaultConfig, errCode := loadDefaultConfig(parentDir, defaultConfig)
	if defaultConfig == nil {
		return nil, Code_Open_File_Failed
	}
	sysConfig := defaultConfig
	sysConfigRefPtr := reflect.ValueOf(sysConfig)
	sysConfigRef := sysConfigRefPtr.Elem()
	LoggerValue := sysConfigRef.FieldByName("Logger")
	if LoggerValue.IsValid() {
		Logger := LoggerValue.Interface().(LoggerConfig)
		InitLog(Logger, processName)
	}
	if env == "" {
		if didLoadConfig != nil {
			didLoadConfig(sysConfig, env, processName)
		}
		logger.Infof("[config][]run environment:default local ip:%s", localIp)
		return sysConfig, Code_Ok
	} else {
		logger.Infof("[config][]run environment:%s local ip:%s", env, localIp)
	}

	var fileName = parentDir + "/config_" + env + ".json"
	_, fError := os.Stat(fileName)
	envConfig := new(T)
	if fError == nil || os.IsExist(fError) {
		b, err := os.ReadFile(fileName)
		if err != nil {
			logger.Fatalf("[config][]InitConfig sys config read json error:%v", err)
		}

		err = jsoniter.Unmarshal(b, envConfig)
		if err != nil && errCode != Code_Ok {
			if err != nil {
				logger.Fatalf("[config][]InitConfig sys config json to struct error:%v", err)
				return nil, Code_Load_Config_Failed
			} else {
				return nil, errCode
			}
		}
		// sysConfig = mergeConfig(envConfig, defaultConfig)
		sysConfig = MergeInterface(envConfig, defaultConfig)
	} else {
		logger.Warnf("[config][]InitConfig env %s no config file", env)
	}

	if LoggerValue.IsValid() {
		Logger := LoggerValue.Interface().(LoggerConfig)
		InitLog(Logger, processName)
	}
	if didLoadConfig != nil {
		didLoadConfig(sysConfig, env, processName)
	}
	return sysConfig, Code_Ok
}

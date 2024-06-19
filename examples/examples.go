package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/linuxkungfu/go-util/examples/config"

	"github.com/linuxkungfu/go-util"
	"github.com/linuxkungfu/go-util/orm"
	"github.com/linuxkungfu/go-util/orm/iorm"
)

func main() {
	// serverId := util.CreateServerId("123")
	// fmt.Printf("server id:%s\n", serverId)
	// emoji, unicode := util.GetFlag("US")
	// fmt.Printf("emoji:%s unicode:%s\n", emoji, unicode)
	// fmt.Printf("country name:%s\n", countries.ByName("US").String())
	sysConfig := &config.SysConfig{}
	util.InitConfig("./etc", "dev", "util", sysConfig, nil)
	util.CreateServerId("127.0.0.1:20001")
	// ipInfo := util.IPToLocationQuery("191.6.52.188")
	// fmt.Printf("ipInfo:%v\n", ipInfo)
	// ipInfo = util.APIIpQuery("191.6.52.188")
	// fmt.Printf("ipInfo:%v\n", ipInfo)
	// ipInfo = util.IPQuery("191.6.52.188")
	// fmt.Printf("ipInfo:%v\n", ipInfo)
	// ipInfo := util.IPRegistryQuery("191.6.52.188")
	// fmt.Printf("ipInfo:%v\n", ipInfo)
	// md5Password := util.PasswordPlainToMd5(124, "123446")
	// fmt.Printf("md5Password:%s\n", md5Password)
	// fmt.Printf("device category:%s", util.GetDeviceCategory("", ""))
	// ipInfo := util.IPInfoIoQuery("191.6.52.188")
	// fmt.Printf("ipInfo:%v\n", ipInfo)

	// fmt.Printf("country name:%s\n", util.CountryRegionConvert("Bosnia and Herzegovina"))

	first := []string{"1001", "1002"}
	second := []string{"1001", "1003", "1003"}
	fmt.Printf("merge unique array:%v\n", util.MergeUniqueArray(first, second))
	for _, config := range sysConfig.Input.MQ {
		configMap := map[string]interface{}{}
		data, err := json.Marshal(config)
		if err == nil {
			json.Unmarshal(data, &configMap)
			orm.SetupORMInstance("test", iorm.ORMType_MQ, iorm.ORMOperateType_Read, configMap)
		}
	}
	time.Sleep(time.Duration(100) * time.Second)
}

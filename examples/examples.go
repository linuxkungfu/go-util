package main

import (
	"fmt"

	"github.com/linuxkungfu/go-util"
	"github.com/linuxkungfu/go-util/dep/countries"
)

func main() {
	serverId := util.CreateServerId("123")
	fmt.Printf("server id:%s\n", serverId)
	emoji, unicode := util.GetFlag("US")
	fmt.Printf("emoji:%s unicode:%s\n", emoji, unicode)
	fmt.Printf("country name:%s\n", countries.ByName("US").String())
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
	fmt.Printf("merge unique array:%v", util.MergeUniqueArray(first, second))
}

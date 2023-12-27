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
	ipInfo := util.IPRegistryQuery("191.6.52.188")
	fmt.Printf("ipInfo:%v\n", ipInfo)
}

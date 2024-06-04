package util

import "os"

var localIp string

func init() {
	envLocalIp := os.Getenv("LOCAL_IP")
	if len(envLocalIp) == 0 {
		localIps, _ := GetInterfaceIpv4("")
		if len(localIps) > 0 {
			localIp = localIps[0]
		}
	} else {
		localIp = envLocalIp
	}
}
func GetLocalIp() string {
	return localIp
}

func UpdateLocalIp(ip string) {
	localIp = ip
}

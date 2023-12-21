package util

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	logger "github.com/sirupsen/logrus"
)

var (
	VIRTUAL_INTERFACE []string = []string{"docker", "flannel", "cni"}
)

type NetworkAddr struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

func (na NetworkAddr) String() string {
	return fmt.Sprintf("%s:%d", na.Ip, na.Port)
}

func (na NetworkAddr) FullString() string {
	return fmt.Sprintf("%s://%s:%d", na.Protocol, na.Ip, na.Port)
}

func (na *NetworkAddr) ParseString(addrString string) bool {
	protoAndAddr := strings.Split(addrString, "://")
	if len(protoAndAddr) != 2 {
		return false
	}
	na.Protocol = protoAndAddr[0]
	ipAndPort := strings.Split(protoAndAddr[1], ":")
	if len(ipAndPort) != 2 {
		return false
	}
	na.Ip = ipAndPort[0]
	port, err := strconv.ParseInt(ipAndPort[1], 10, 32)
	if err != nil {
		return false
	}
	na.Port = int(port)
	return true
}
func (na NetworkAddr) UdpConn() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", na.Ip, na.Port))
	if err != nil {
		logger.Warnf("[network][%s]error:%s", fmt.Sprintf("%s:%d", na.Ip, na.Port), err.Error())
		return nil
	}
	conn, dialError := net.DialUDP("udp", nil, addr)
	if dialError != nil {
		logger.Warnf("[network][%s]error:%s", fmt.Sprintf("%s:%d", na.Ip, na.Port), dialError.Error())
		return nil
	} else {
		return conn
	}
}

func IsSameNetworkAddr(addrs []NetworkAddr) bool {
	if len(addrs) <= 1 {
		return true
	}
	firstAddr := addrs[0]
	for index := 1; index < len(addrs); index++ {
		addr := addrs[index]
		if addr.Ip != firstAddr.Ip {
			return false
		}
		if addr.Port != firstAddr.Port {
			return false
		}
		if addr.Protocol != firstAddr.Protocol {
			return false
		}
	}
	return true
}

func NetworkAddrsToString(addrs []NetworkAddr, sep string) string {
	addrStrs := []string{}
	for _, addr := range addrs {
		addrStrs = append(addrStrs, addr.FullString())
	}
	return strings.Join(addrStrs, sep)
}

// 判断是否为公网ip
func IsPublicIPv4(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// 判断是否为内网ip
func IsPrivateIPv4(IP net.IP) bool {
	return !IsPublicIPv4(IP)
}

func IsPublicIPv4Str(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return IsPublicIPv4(ip)
}

func IsPrivateIPv4Str(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return IsPrivateIPv4(ip)
}

// 获取网卡ip
func GetInterfaceIpv4(filter string) ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return []string{}, err
	}
	ips := []string{}
	for _, i := range interfaces {
		name := i.Name
		isVirtual := false
		for _, virtualName := range VIRTUAL_INTERFACE {
			if strings.Contains(name, virtualName) {
				isVirtual = true
				break
			}
		}
		// interface is virtual or belongs to a point-to-point link
		if isVirtual || i.Flags&net.FlagPointToPoint != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var (
				ip net.IP
			)
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ipv4 := ip.To4(); ipv4 != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast() && (filter == "" || strings.Contains(name, filter)) {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips, nil
}

// 该函数只能在能访问公网环境下使用，否则会出问题
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func SameNetworkAddr(first NetworkAddr, second NetworkAddr) bool {
	if first.Ip == second.Ip && first.Port == second.Port && first.Protocol == second.Protocol {
		return true
	}
	return false
}

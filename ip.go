package util

import (
	"fmt"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	Ip_APIUrl         string = "http://ip-api.com/json/"
	IpStackApiUrl     string = "http://api.ipstack.com"
	APIIpUrl          string = "http://apiip.net/api"
	IPApiUrl          string = "http://api.ipapi.com/api/"
	IpUserAgentInfo   string = "https://ip.useragentinfo.com/json"
	IpStackApiKey1    string = "7526b5001e2cc6fbc854feddc19e4a76"
	IpStackApiKey2    string = "e3dcfe9ed9635455e3333ce8eadb9ea3"
	APIIpKey          string = "3f740f6d-7ff3-41d0-bcf7-2f844d6832f5"
	IPApiKey          string = "8557773513ffb5242020ad75fdf76e97"
	IpQueryMaxTimeout        = time.Duration(60) * time.Second
)

var (
	IpStackApiKey string = IpStackApiKey2
)

type IPLanguage struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

type IPCurencyRates struct {
	EURUSD float32 `json:"EURUSD,omitempty"`
}

type IPCurency struct {
	Code   string         `json:"code,omitempty"`
	Name   string         `json:"name,omitempty"`
	Symbol string         `json:"symbol,omitempty"`
	Number string         `json:"number,omitempty"`
	Rates  IPCurencyRates `json:"rates,omitempty"`
}
type IPTimeZone struct {
	Id           string    `json:"id,omitempty"`
	CurrentTime  time.Time `json:"currentTime,omitempty"`
	Code         string    `json:"code,omitempty"`
	TimeZoneName string    `json:"timeZoneName,omitempty"`
	UtcOffset    int       `json:"utcOffset,omitempty"`
}
type IPUserAgent struct {
	IsMobile        bool   `json:"isMobile,omitempty"`
	IsiPod          bool   `json:"isiPod,omitempty"`
	IsTablet        bool   `json:"isTablet,omitempty"`
	IsDesktop       bool   `json:"isDesktop,omitempty"`
	IsSmartTV       bool   `json:"isSmartTV,omitempty"`
	IsRaspberry     bool   `json:"isRaspberry,omitempty"`
	IsBot           bool   `json:"isBot,omitempty"`
	Browser         string `json:"browser,omitempty"`
	BrowserVersion  string `json:"browserVersion,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
	Platform        string `json:"platform,omitempty"`
	Source          string `json:"source,omitempty"`
}

type IPConnection struct {
	ASN int    `json:"asn,omitempty"`
	ISP string `json:"isp,omitempty"`
}
type IPSecurity struct {
	IsProxy       bool   `json:"isProxy,omitempty"`
	IsBogon       bool   `json:"isBogon,omitempty"`
	IsTorExitNode bool   `json:"isTorExitNode,omitempty"`
	IsCloud       bool   `json:"isCloud,omitempty"`
	IsHosting     bool   `json:"isHosting,omitempty"`
	IsSpamhaus    bool   `json:"isSpamhaus,omitempty"`
	Suggestion    string `json:"suggestion,omitempty"`
	Network       string `json:"network,omitempty"`
}
type APIIPMessage struct {
	Code int    `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
	Info string `json:"info,omitempty"`
}

type IPInfo struct {
	Status                  string    `json:"status"`
	Country                 string    `json:"country"`
	CountryCode             string    `json:"countryCode"`
	Region                  string    `json:"region"`
	RegionName              string    `json:"regionName"`
	City                    string    `json:"city"`
	Zip                     string    `json:"zip"`
	Lat                     float32   `json:"lat"`
	Lon                     float32   `json:"lon"`
	Timezone                string    `json:"timezone"`
	ISP                     string    `json:"isp"`
	Org                     string    `json:"org"`
	As                      string    `json:"as"`
	Query                   string    `json:"query"`
	CountryFlag             string    `json:"country_flag"`
	CountryFlagEmoji        string    `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string    `json:"country_flag_emoji_unicode"`
	UpdateTS                time.Time `json:"update_ts"`
}

type APIIPInfo struct {
	Success                bool         `json:"success,omitempty"`
	Message                APIIPMessage `json:"message,omitempty"`
	Ip                     string       `json:"ip,omitempty"`
	ContinentCode          string       `json:"continentCode,omitempty"`
	ContinentName          string       `json:"continentName,omitempty"`
	CountryCode            string       `json:"countryCode,omitempty"`
	CountryName            string       `json:"countryName,omitempty"`
	CountryNameNative      string       `json:"countryNameNative,omitempty"`
	OfficialCountryName    string       `json:"officialCountryName,omitempty"`
	RegionCode             string       `json:"regionCode,omitempty"`
	RegionName             string       `json:"regionName,omitempty"`
	City                   string       `json:"city,omitempty"`
	PostalCode             string       `json:"postalCode,omitempty"`
	Latitude               float32      `json:"latitude,omitempty"`
	Longitude              float32      `json:"longitude,omitempty"`
	Capital                string       `json:"capital,omitempty"`
	PhoneCode              string       `json:"phoneCode,omitempty"`
	CountryFlagEmoj        string       `json:"countryFlagEmoj,omitempty"`
	CountryFlagEmojUnicode string       `json:"countryFlagEmojUnicode,omitempty"`
	IsEu                   bool         `json:"isEu,omitempty"`
	Borders                [2]string    `json:"borders,omitempty"`
	TopLevelDomains        []string     `json:"topLevelDomains,omitempty"`
	Languages              IPLanguage   `json:"languages,omitempty"`
	Currency               IPCurency    `json:"currency,omitempty"`

	TimeZone   IPTimeZone   `json:"timeZone,omitempty"`
	UserAgent  IPUserAgent  `json:"userAgent,omitempty"`
	Connection IPConnection `json:"connection,omitempty"`
	Security   IPSecurity   `json:"security,omitempty"`
	UpdateTS   time.Time    `json:"update_ts,omitempty"`
}

type IPLocationInfo struct {
	GeonameId               int32        `json:"geoname_id"`
	Capital                 string       `json:"capital"`
	Languages               []IPLanguage `json:"languages"`
	CountryFlag             string       `json:"country_flag"`
	CountryFlagEmoji        string       `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string       `json:"country_flag_emoji_unicode"`
	CallingCode             string       `json:"calling_code"`
	ISEU                    bool         `json:"is_eu"`
}
type IPStackInfo struct {
	Success bool `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"error,omitempty"`
	Ip            string         `json:"ip"`
	Type          string         `json:"type"`
	ContinentCode string         `json:"continent_code"`
	ContinentName string         `json:"continent_name"`
	CountryCode   string         `json:"country_code"`
	CountryName   string         `json:"country_name"`
	RegionCode    string         `json:"region_code"`
	RegionName    string         `json:"region_name"`
	City          string         `json:"city"`
	Zip           string         `json:"zip"`
	Latitude      float32        `json:"latitude"`
	Longitude     float32        `json:"longitude"`
	Location      IPLocationInfo `json:"location"`
}

func IPQuery(ip string) *IPInfo {
	res := APIIpQuery(ip)
	if res == nil {
		res = IPStackQuery(ip)
		if res == nil {
			return IP_ApiQuery(ip)
		} else {
			return res
		}
	} else {
		return res
	}
}

func IP_ApiQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s%s", Ip_APIUrl, ip)
	ipInfo, err := HttpGetJson(url, &IPInfo{}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IP_ApiQuery] new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	if ipInfo.(*IPInfo).Status == "fail" {
		return nil
	}
	ipInfo.(*IPInfo).UpdateTS = time.Now()
	return ipInfo.(*IPInfo)
}
func APIIpQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s/check?ip=%s&accessKey=%s", APIIpUrl, ip, APIIpKey)
	apiIpInfo := &APIIPInfo{Success: true}
	var err error
	_, err = HttpGetJson(url, apiIpInfo, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPApiQuery] new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	if !apiIpInfo.Success {
		logger.Warnf("[util][IPApiQuery] new query ip:%s failed:%s, type:%s", ip, apiIpInfo.Message.Info, apiIpInfo.Message.Type)
		return nil
	}
	apiIpInfo.UpdateTS = time.Now()
	ipInfo := &IPInfo{}
	ipInfo.City = apiIpInfo.City
	ipInfo.Country = apiIpInfo.CountryName
	ipInfo.CountryCode = apiIpInfo.CountryCode
	ipInfo.CountryFlag = apiIpInfo.CountryFlagEmojUnicode
	ipInfo.CountryFlagEmoji = apiIpInfo.CountryFlagEmoj
	ipInfo.CountryFlagEmojiUnicode = apiIpInfo.CountryFlagEmojUnicode
	ipInfo.Lat = apiIpInfo.Latitude
	ipInfo.Lon = apiIpInfo.Longitude
	ipInfo.Query = apiIpInfo.Ip
	ipInfo.Region = apiIpInfo.RegionName
	ipInfo.RegionName = apiIpInfo.RegionCode
	ipInfo.Zip = apiIpInfo.PostalCode
	ipInfo.UpdateTS = apiIpInfo.UpdateTS
	return ipInfo
}

func IPUserAgentInfoQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s?ip=%s", IpUserAgentInfo, ip)
	_, err := HttpGetJson(url, &IPInfo{}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPUserAgentInfoQuery] new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	return nil
}

func IPStackQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s/%s/?access_key=%s", IpStackApiUrl, ip, IpStackApiKey)
	ipInfoIf, err := HttpGetJson(url, &IPStackInfo{Success: true}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPStackQuery]new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	ipStackInfo := ipInfoIf.(*IPStackInfo)
	if !ipStackInfo.Success && ipStackInfo.Error.Code == 104 {
		logger.Warnf("[util][IPStackQuery]api response error code:%d error message:%s switch api key", ipStackInfo.Error.Code, ipStackInfo.Error.Info)
		if IpStackApiKey == IpStackApiKey1 {
			IpStackApiKey = IpStackApiKey2
		} else {
			IpStackApiKey = IpStackApiKey1
		}
		return IPStackQuery(ip)
	} else if !ipStackInfo.Success {
		logger.Warnf("[util][IPStackQuery]api response error code:%d error message:%s", ipStackInfo.Error.Code, ipStackInfo.Error.Info)
		return nil
	}
	ipInfo := &IPInfo{}
	ipInfo.City = ipStackInfo.City
	ipInfo.Country = ipStackInfo.CountryName
	ipInfo.CountryCode = ipStackInfo.CountryCode
	ipInfo.CountryFlag = ipStackInfo.Location.CountryFlag
	ipInfo.CountryFlagEmoji = ipStackInfo.Location.CountryFlagEmoji
	ipInfo.CountryFlagEmojiUnicode = ipStackInfo.Location.CountryFlagEmojiUnicode
	ipInfo.Lat = ipStackInfo.Latitude
	ipInfo.Lon = ipStackInfo.Longitude
	ipInfo.Query = ipStackInfo.Ip
	ipInfo.Region = ipStackInfo.RegionName
	ipInfo.RegionName = ipStackInfo.RegionCode
	ipInfo.Zip = ipStackInfo.Zip
	ipInfo.UpdateTS = time.Now()
	return ipInfo
}

func GetAddrFromNetworkAddr(addr string) string {
	if strings.Contains(addr, "[") && strings.Contains(addr, "]") {
		values := strings.Split(addr, "]:")
		return strings.Trim(strings.Trim(values[0], "["), "]")
	} else if strings.Contains(addr, ".") {
		values := strings.Split(addr, ":")
		return values[0]
	} else {
		return ""
	}
}

func GetAddrRegion(addr string) *IPInfo {
	ip := GetAddrFromNetworkAddr(addr)
	if ip == "" {
		return nil
	}
	return IPQuery(ip)
}

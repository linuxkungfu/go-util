package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/linuxkungfu/go-util/dep/countries"
	"github.com/redis/go-redis/v9"
	logger "github.com/sirupsen/logrus"
)

const (
	Ip_APIUrl          string = "http://ip-api.com/json/"
	IpStackApiUrl      string = "http://api.ipstack.com"
	APIIpUrl           string = "http://apiip.net/api"
	IPApiUrl           string = "http://api.ipapi.com/api/"
	IpUserAgentInfoUrl string = "https://ip.useragentinfo.com/json"
	IPToLocationUrl    string = "https://api.ip2location.io/"
	IPRegistryUrl      string = "https://api.ipregistry.co/"
	IPInfoUrl          string = "https://ipinfo.io/"

	IpStackApiKey1    string = "7526b5001e2cc6fbc854feddc19e4a76"
	IpStackApiKey2    string = "e3dcfe9ed9635455e3333ce8eadb9ea3"
	APIIpKey          string = "3f740f6d-7ff3-41d0-bcf7-2f844d6832f5"
	IPApiKey          string = "8557773513ffb5242020ad75fdf76e97"
	IPToLocationKey   string = "5DFCDB79756CE10039FCE40E36EB632D"
	IPRegistryKey     string = "3fbw90yjv8v0pog4"
	IPInfoKey         string = "c4e6f9d5b42c85"
	IpQueryMaxTimeout        = time.Duration(60) * time.Second
)

var (
	IpStackApiKey    string             = IpStackApiKey2
	ipInfoMap        map[string]*IPInfo = map[string]*IPInfo{}
	ipQueryFunc                         = [...]func(string) *IPInfo{IPToLocationQuery, IPInfoIoQuery, IPRegistryQuery, APIIpQuery, IPStackQuery, IP_ApiQuery}
	ipQueryFuncIndex                    = 0
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
type IPRegistryCountryFlag struct {
	Emoji        string `json:"emoji"`
	EmojiUnicode string `json:"emoji_unicode"`
	EmojiTwo     string `json:"emojitwo"`
	Noto         string `json:"noto"`
	Twemoji      string `json:"twemoji"`
	Wikimedia    string `json:"wikimedia"`
}
type IPRegistryCountry struct {
	Name              string                `json:"name"`
	Code              string                `json:"code"`
	Capital           string                `json:"capital"`
	Area              string                `json:"area"`
	Borders           string                `json:"borders"`
	CallingCode       string                `json:"calling_code"`
	Population        int                   `json:"population"`
	PopulationDensity int                   `json:"population_density"`
	Flag              IPRegistryCountryFlag `json:"flag"`
	Languages         []IPLanguage          `json:"languages"`
	Tld               string                `json:"tld"`
}
type IPRegistryRegion struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type IPRegisterLocationContinent struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type IPRegistryLocation struct {
	Continent IPRegisterLocationContinent `json:"continent"`
	Country   IPRegistryCountry           `json:"country"`
	Region    IPRegistryRegion            `json:"region"`
	City      string                      `json:"city"`
	Postal    string                      `json:"postal"`
	Latitude  float32                     `json:"latitude"`
	Longitude float32                     `json:"longitude"`
	Language  IPLanguage                  `json:"language"`
	In_eu     bool                        `json:"in_eu"`
}
type IPRegistrySecurity struct {
	Is_anonymous      bool `json:"is_anonymous"`
	Is_abuser         bool `json:"is_abuser"`
	Is_attacker       bool `json:"is_attacker"`
	Is_bogon          bool `json:"is_bogon"`
	Is_cloud_provider bool `json:"is_cloud_provider"`
	Is_proxy          bool `json:"is_proxy"`
	Is_relay          bool `json:"is_relay"`
	Is_threat         bool `json:"is_threat"`
	Is_tor            bool `json:"is_tor"`
	Is_tor_exit       bool `json:"is_tor_exit"`
	Is_vpn            bool `json:"is_vpn"`
}
type IPRegistryTimeZone struct {
	Id                 string `json:"id"`
	Abbreviation       string `json:"abbreviation"`
	CurrentTime        string `json:"current_time"`
	Name               string `json:"name"`
	Offset             int    `json:"offset"`
	In_daylight_saving bool   `json:"in_daylight_saving"`
}
type IPRegistryCurrencyFormatNegative struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}
type IPRegistryCurrencyFormatPositive struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type IPRegistryCurrencyFormat struct {
	Negative IPRegistryCurrencyFormatNegative `json:"negative"`
	Positive IPRegistryCurrencyFormatPositive `json:"positive"`
}

type IPRegistryCurrency struct {
	Code         string                   `json:"code"`
	Name         string                   `json:"name"`
	NameNative   string                   `json:"name_native"`
	Plural       string                   `json:"plural"`
	PluralNative string                   `json:"plural_native"`
	Symbol       string                   `json:"symbol"`
	SymbolNative string                   `json:"symbol_native"`
	Format       IPRegistryCurrencyFormat `json:"format"`
}

type IPRegistryCarrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}
type IPRegistryCompany struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}
type IPRegistryConnection struct {
	Asn          string `json:"asn"`
	Domain       string `json:"domain"`
	Organization string `json:"organization"`
	Route        string `json:"route"`
	Type         string `json:"type"`
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

type IPTOLacationInfo struct {
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionName  string  `json:"region_name"`
	CityName    string  `json:"city_name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Zip         string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	ASN         string  `json:"asn"`
	As          string  `json:"as"`
	IsProxy     bool    `json:"is_proxy"`
}

type IPRegistryInfo struct {
	Ip         string               `json:"ip"`
	Type       string               `json:"type"`
	Carrier    IPRegistryCarrier    `json:"carrier"`
	Company    IPRegistryCompany    `json:"company"`
	Connection IPRegistryConnection `json:"connection"`
	Currency   IPRegistryCurrency   `json:"currency"`
	Location   IPRegistryLocation   `json:"location"`
	Security   IPRegistrySecurity   `json:"security"`
	TimeZone   IPRegistryTimeZone   `json:"time_zone"`
}

type IPInfoIo struct {
	Ip          string `json:"ip"`
	City        string `json:"city"`
	Region      string `json:"region"`
	CountryCode string `json:"country"`
	Loc         string `json:"loc"`
	Org         string `json:"org"`
	Timezone    string `json:"timezone"`
}

func IPQuery(ip string) *IPInfo {
	index := ipQueryFuncIndex
	for {
		res := ipQueryFunc[index](ip)
		if res != nil {
			ipQueryFuncIndex = index
			return res
		} else {
			index = (index + 1) % len(ipQueryFunc)
			if index == ipQueryFuncIndex {
				return nil
			}
		}
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
	url := fmt.Sprintf("%s?ip=%s", IpUserAgentInfoUrl, ip)
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

func IPToLocationQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s?key=%s&ip=%s", IPToLocationUrl, IPToLocationKey, ip)
	ipInfoIf, err := HttpGetJson(url, &IPTOLacationInfo{}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPToLocationQuery]new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	ipToLocationInfo := ipInfoIf.(*IPTOLacationInfo)
	ipInfo := &IPInfo{}
	ipInfo.City = ipToLocationInfo.CityName
	ipInfo.Country = countries.ByName(ipToLocationInfo.CountryCode).String()
	ipInfo.CountryCode = ipToLocationInfo.CountryCode
	ipInfo.CountryFlagEmoji, ipInfo.CountryFlagEmojiUnicode = GetFlag(ipToLocationInfo.CountryCode)
	ipInfo.Lat = ipToLocationInfo.Latitude
	ipInfo.Lon = ipToLocationInfo.Longitude
	ipInfo.Query = ipToLocationInfo.Ip
	// ipInfo.Region = ipToLocationInfo.Regioname
	ipInfo.RegionName = ipToLocationInfo.RegionName
	ipInfo.Zip = ipToLocationInfo.Zip
	ipInfo.UpdateTS = time.Now()
	return ipInfo
}
func IPRegistryQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s%s?key=%s", IPRegistryUrl, ip, IPRegistryKey)
	ipInfoIf, err := HttpGetJson(url, &IPRegistryInfo{}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPToLocationQuery]new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	ipRegistryInfo := ipInfoIf.(*IPRegistryInfo)
	ipInfo := &IPInfo{}

	ipInfo.City = ipRegistryInfo.Location.City
	ipInfo.Country = ipRegistryInfo.Location.Country.Name
	ipInfo.CountryFlag = ipRegistryInfo.Location.Country.Flag.Noto
	ipInfo.CountryCode = ipRegistryInfo.Location.Country.Code
	ipInfo.CountryFlagEmoji = ipRegistryInfo.Location.Country.Flag.Emoji
	ipInfo.CountryFlagEmojiUnicode = ipRegistryInfo.Location.Country.Flag.EmojiUnicode
	ipInfo.Lat = ipRegistryInfo.Location.Latitude
	ipInfo.Lon = ipRegistryInfo.Location.Longitude
	ipInfo.Query = ipRegistryInfo.Ip
	ipInfo.Region = ipRegistryInfo.Location.Region.Code
	ipInfo.RegionName = ipRegistryInfo.Location.Region.Name
	ipInfo.Zip = ipRegistryInfo.Location.Postal
	ipInfo.Timezone = ipRegistryInfo.TimeZone.Name
	ipInfo.ISP = ipRegistryInfo.Company.Name
	ipInfo.UpdateTS = time.Now()
	return ipInfo
}

func IPInfoIoQuery(ip string) *IPInfo {
	url := fmt.Sprintf("%s%s?token=%s", IPInfoUrl, ip, IPInfoKey)
	ipInfoIf, err := HttpGetJson(url, &IPInfoIo{}, IpQueryMaxTimeout)
	if err != nil {
		logger.Warnf("[util][IPToLocationQuery]new query ip:%s failed:%s", ip, err.Error())
		return nil
	}
	ipInfoIo := ipInfoIf.(*IPInfoIo)
	ipInfo := &IPInfo{}

	ipInfo.City = ipInfoIo.City
	ipInfo.Country = countries.ByName(ipInfoIo.CountryCode).String()
	ipInfo.CountryCode = ipInfoIo.CountryCode
	ipInfo.CountryFlagEmoji, ipInfo.CountryFlagEmojiUnicode = GetFlag(ipInfoIo.CountryCode)
	locArray := strings.Split(ipInfoIo.Loc, ",")
	value, _ := strconv.ParseFloat(locArray[0], 64)
	ipInfo.Lat = float32(value)
	value, _ = strconv.ParseFloat(locArray[1], 64)
	ipInfo.Lon = float32(value)
	ipInfo.Query = ipInfoIo.Ip
	ipInfo.RegionName = ipInfoIo.Region
	ipInfo.Timezone = ipInfoIo.Timezone
	ipInfo.ISP = ipInfoIo.Org
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

func QueryIpInfo(readClient *redis.Client, writeClient *redis.Client, ip string) *IPInfo {
	// 先不设置过期时间
	ipInfo, ok := ipInfoMap[ip]
	if ok {
		if ipInfo.Country != "unknown" && ipInfo.CountryFlagEmoji != "" {
			return ipInfo
		} else {
			logger.Warnf("[util][QueryIpInfo] new query ip:%s country:%s or emjo is empty", ip, ipInfo.Country)
		}
	}
	key := fmt.Sprintf("ip_query_%s", ip)
	if readClient != nil && writeClient != nil {
		lockValue := AcquireSpinLock(writeClient, key, time.Duration(5)*time.Second, time.Duration(3)*time.Second)
		if lockValue != 0 {
			defer func() {
				ReleaseSpinLock(writeClient, key, lockValue)
			}()
			ipInfo = &IPInfo{}
			object := GetObjectFromRedis(readClient, key, ipInfo)
			if object != nil {
				ipInfo = object.(*IPInfo)
				if ipInfo.Country != "unknown" && ipInfo.CountryFlagEmoji != "" {
					ipInfoMap[ip] = ipInfo
					return ipInfo
				} else {
					logger.Warnf("[util][QueryIpInfo] new query ip:%s country:%s or emjo is empty", ip, ipInfo.Country)
				}

			}
		}
	}
	ipInfo = IPQuery(ip)
	if ipInfo != nil {
		ipInfoMap[ip] = ipInfo
		if writeClient != nil {
			SetObjectToRedis(writeClient, key, ipInfo, time.Duration(2400)*time.Hour)
		}
	}
	return ipInfo
}

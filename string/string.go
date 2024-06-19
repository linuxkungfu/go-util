package string

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spaolacci/murmur3"
)

func MapToStruct(m map[string]interface{}, s interface{}) error {
	return InterfaceToStruct(m, s)
}

func InterfaceToStruct(i interface{}, s interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	} else {
		err := json.Unmarshal(data, s)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

var sessionId uint64 = 0

var charList = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// charsList := [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

var charLen int = len(charList)

func initSessionId() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	hashValue := fmt.Sprintf("%s:%d", hostname, time.Now().Unix())
	hValue1 := murmur3.Sum32WithSeed([]byte(hashValue), uint32(time.Now().Nanosecond()))
	sessionId = uint64(hValue1) << 24
}

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func CapacityToInteger(capacityString string) uint64 {
	strLen := len(capacityString)
	endIndex := strLen
	for index := 0; index < len(capacityString); index++ {
		if capacityString[index] < 48 || capacityString[index] > 57 {
			endIndex = index
			break
		}
	}
	unit := capacityString[endIndex:strLen]
	number := capacityString[0:endIndex]
	if strings.EqualFold(unit, "KB") || strings.EqualFold(unit, "K") {
		size, _ := strconv.ParseUint(number, 10, 64)
		return size * 1024
	} else if strings.EqualFold(unit, "MB") || strings.EqualFold(unit, "M") {
		size, _ := strconv.ParseUint(number, 10, 64)
		return size * 1024 * 1024
	} else if strings.EqualFold(unit, "GB") || strings.EqualFold(unit, "G") {
		size, _ := strconv.ParseUint(number, 10, 64)
		return size * 1024 * 1024 * 1024
	} else if strings.EqualFold(unit, "TB") || strings.EqualFold(unit, "T") {
		size, _ := strconv.ParseUint(number, 10, 64)
		return size * 1024 * 1024 * 1024 * 1024
	} else if strings.EqualFold(unit, "PB") || strings.EqualFold(unit, "P") {
		size, _ := strconv.ParseUint(number, 10, 64)
		return size * 1024 * 1024 * 1024 * 1024
	} else {
		size, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			return 0
		}
		return size
	}
}

// StringToTime convert string to time format
func StringToTime(stringTime string) time.Duration {
	lastChar := stringTime[len(stringTime)-1:]
	valueStr := stringTime[0 : len(stringTime)-1]
	value, err := strconv.ParseInt(valueStr, 10, 0)
	if err != nil {
		return 0
	}
	if strings.EqualFold(lastChar, "d") {
		return time.Hour * time.Duration(value*24)
	} else if strings.EqualFold(lastChar, "h") {
		return time.Hour * time.Duration(value)
	} else if strings.EqualFold(lastChar, "m") {
		return time.Minute * time.Duration(value)
	} else if strings.EqualFold(lastChar, "s") {
		return time.Minute * time.Duration(value)
	} else {
		// logger.Warnf("[config][]StringToTime failed, time:%s", stringTime)
		return -1
	}
}

func FormatFloat(num float64, decimal int) float64 {
	d := float64(1)
	if decimal > 0 {
		d = math.Pow10(decimal)
	}
	return math.Trunc(num*d) / d
}

func FormatFloatString(num float64, decimal int) string {
	return strconv.FormatFloat(FormatFloat(num, decimal), 'f', -1, 64)
}

func CreateSessionId() uint64 {
	if sessionId == 0 {
		initSessionId()
	}
	sessionId += 1
	return sessionId
}

func CreateRandString(strLen int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(source)
	buffer := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		buffer[i] = charList[rd.Intn(charLen)]
	}
	return string(buffer)
}

func StrArray2IntArray(arr []string) []int {
	intArray := []int{}
	for _, e := range arr {
		v, _ := strconv.ParseInt(e, 10, 32)
		intArray = append(intArray, int(v))
	}
	return intArray
}

func GetStructTypeName(st interface{}) string {
	if t := reflect.TypeOf(st); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func AnyToStr(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch v := value.(type) {
	case float64:
		ft := v
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := v
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := v
		key = strconv.Itoa(it)
	case uint:
		it := v
		key = strconv.Itoa(int(it))
	case int8:
		it := v
		key = strconv.Itoa(int(it))
	case uint8:
		it := v
		key = strconv.Itoa(int(it))
	case int16:
		it := v
		key = strconv.Itoa(int(it))
	case uint16:
		it := v
		key = strconv.Itoa(int(it))
	case int32:
		it := v
		key = strconv.Itoa(int(it))
	case uint32:
		it := v
		key = strconv.Itoa(int(it))
	case int64:
		it := v
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := v
		key = strconv.FormatUint(it, 10)
	case string:
		key = v
	case []byte:
		key = string(v)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func StringToDuration(durStr string) time.Duration {
	re := regexp.MustCompile(`^([0-9]+)([a-zA-Z]+)$`)
	values := re.FindStringSubmatch(durStr)
	if len(values) != 3 {
		return 0
	}
	value, err := strconv.Atoi(values[1])
	if err != nil {
		return 0
	}
	switch values[2] {
	case "m", "min", "minute":
		return time.Duration(value) * time.Minute
	case "s", "sec", "second", "S":
		return time.Duration(value) * time.Second
	case "h", "hour", "H":
		return time.Duration(value) * time.Hour
	case "d", "day", "D":
		return time.Duration(value) * time.Hour * 24
	case "w", "week", "W":
		return time.Duration(value) * time.Hour * 24 * 7
	case "M", "month":
		return time.Duration(value) * time.Hour * 24 * 30
	case "y", "year", "Y":
		return time.Duration(value) * time.Hour * 24 * 30 * 365
	}
	return 0
}

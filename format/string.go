package format

import (
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func IntToStr(content int) string {
	return strconv.Itoa(content)
}

func Int64ToStr(content int64) string {
	return strconv.FormatInt(content, 10)
}

func StrToInt(content string) int {
	result, err := strconv.Atoi(content)
	if nil != err {
		return 0
	} else {
		return result
	}
}

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func StructToStr(content interface{}) string {
	str, _ := json.Marshal(content)
	return string(str)
}

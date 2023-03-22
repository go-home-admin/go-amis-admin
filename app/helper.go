package app

import (
	"crypto/md5"
	"fmt"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// PushQueue 投递进入队列的简短函数
func PushQueue(msg interface{}) {
	providers.GetBean("queue").(*servers.Queue).Push(msg)
}

// GetStructFieldsInfo 获取模型信息
func GetStructFieldsInfo(s interface{}) (fields []map[string]interface{}) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		fieldInfo := make(map[string]interface{})
		fieldInfo["name"] = GetTagInfoWitch(tag, "column")
		fieldInfo["comment"] = GetTagInfoWitch(tag, "comment")

		fields = append(fields, fieldInfo)
	}
	return
}

func GetTagInfoWitch(tag, search string) string {
	arr := strings.Split(tag, ";")
	for _, s := range arr {
		arr2 := strings.Split(s, ":")
		if len(arr2) == 2 && search == arr2[0] {
			return strings.ReplaceAll(arr2[1], "'", "")
		}
	}

	return ""
}

func MD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Int32 强制转换
func Int32(v interface{}) int32 {
	switch val := v.(type) {
	case int32:
		return val
	case int:
		return int32(val)
	case int64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		num, err := strconv.Atoi(val)
		if err != nil {
			// 无法解析为数字，返回 0
			return 0
		}
		if num > math.MaxInt32 || num < math.MinInt32 {
			// 超出 int32 的范围，返回 0
			return 0
		}
		return int32(num)
	default:
		s := fmt.Sprintf("%d", v)
		return Int32(s)
	}
}

// UInt32 强制转换
func UInt32(v interface{}) uint32 {
	switch val := v.(type) {
	case uint32:
		return val
	case uint:
		return uint32(val)
	case uint64:
		return uint32(val)
	case float32:
		return uint32(val)
	case float64:
		return uint32(val)
	case string:
		num, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return uint32(num)
	default:
		s := fmt.Sprintf("%d", v)
		return UInt32(s)
	}
}

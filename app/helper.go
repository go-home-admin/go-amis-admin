package app

import (
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"reflect"
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

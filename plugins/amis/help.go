package amis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

type Url map[string]interface{}

func (u Url) Method(v string) Url {
	u["method"] = v
	return u
}

func (u Url) Url(v string) Url {
	u["url"] = v
	return u
}

func (u Url) Append(v string) Url {
	u["url"] = u.String() + v
	return u
}

func (u Url) Data(v interface{}) Url {
	u["data"] = v
	return u
}

func (u Url) String() string {
	v := u["url"]
	return v.(string)
}

// GetUrl action = /list | /edit | /del
func GetUrl(ctx *gin.Context, action string) Url {
	return Url{
		"method": "get",
		"url":    app.Config("app.url", "http://127.0.0.1") + ctx.Request.URL.RequestURI() + action,
	}
}

func GetInt(ctx *gin.Context, k string, def int) int {
	v := ctx.Query(k)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		logrus.Error("GetInt", err)
		return 0
	}
	return i
}

func NewPageForm(form *Form) *Page {
	page := NewPage("创建")
	page.Body = form
	return page
}

func getPrimaryKey(model interface{}) (string, interface{}) {
	// 使用反射获取主键字段
	reflectValue := reflect.ValueOf(model).Elem()
	var primaryKey reflect.Value
	var primaryName string
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Type().Field(i)
		tag := field.Tag.Get("gorm")
		if strings.Index(tag, "primaryKey") != -1 {
			primaryKey = reflectValue.Field(i)
			primaryName = field.Tag.Get("json")
			break
		}
	}
	if !primaryKey.IsValid() {
		return "", nil
	}

	// 获取主键值
	return primaryName, primaryKey.Interface()
}

// ToCamelCase 转驼峰
func ToCamelCase(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	result := strings.Join(words, "")
	return result
}

func structToMap(input interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	// 获取 struct 的反射值
	val := reflect.ValueOf(input)

	// 如果不是 struct 类型或者是空指针，直接返回空 map
	if val.Kind() != reflect.Struct || val.IsNil() {
		return output
	}

	// 获取 struct 的类型信息
	typ := val.Type()

	// 遍历 struct 的字段
	for i := 0; i < val.NumField(); i++ {
		// 获取字段名和值
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()

		// 将字段名和值存储到 map 中
		output[fieldName] = fieldValue
	}

	return output
}

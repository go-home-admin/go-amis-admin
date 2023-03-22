package amis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

// GetUrl action = /list | /edit | /del
func GetUrl(ctx *gin.Context, action string) string {
	domain := app.Config("app.url", "http://127.0.0.1")
	return domain + ctx.Request.URL.RequestURI() + action
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

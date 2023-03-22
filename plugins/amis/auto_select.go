package amis

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// 字段名称格式是 table.field 时，自动补充上去
// 识别方案
// 1。识别模型的依赖关系
func (c *CurdController) authSelect(list []map[string]interface{}, columns []interface{}) {
	if len(list) == 0 {
		// 没有数据需要分析
		return
	}
	// 需要自动读取其他表数据的信息map[表][前端字段] = 表字段
	tableToIds := map[string]map[string]string{}
	for _, columnT := range columns {
		if column, ok := columnT.(*ColumnConfig); ok {
			if strings.Index(column.Name, ".") != -1 {
				arr := strings.Split(column.Name, ".")
				tableName := arr[0]
				fieldName := arr[1]
				if _, ok := tableToIds[tableName]; !ok {
					tableToIds[tableName] = make(map[string]string)
				}
				tableToIds[tableName][column.Name] = fieldName
			}
		}
	}
	// 读取主表的依赖关系
	indexInfos := getModelInfo(c.model.GetTableInfo())
	for tableName, _ := range tableToIds {
		c.setListAndSelect(list, tableName, indexInfos)
	}
}

func (c *CurdController) setListAndSelect(list []map[string]interface{}, readTable string, myIndexList map[string]indexInfo) {
	for _, info := range myIndexList {
		otherTable := info.tags[info.index]
		switch info.index {
		case "belongs_to":
		case "has_one":
		case "has_many":
		case "many2many":
			if ToCamelCase(readTable) == info.targetModelName {
				c.many2many(list, readTable, otherTable, info)
			}
		}
	}
}

func (c *CurdController) many2many(list []map[string]interface{}, readTable, connectTable string, info indexInfo) {
	JoinForeignKey := info.tags["joinForeignKey"] //本表在连接表的外键
	JoinTargetKey := info.tags["joinReferences"]  //关联表在连接表的外键
	if JoinForeignKey == "" || JoinTargetKey == "" {
		logrus.Error("many2many配置, 没有joinForeignKey｜JoinTargetKey")
		return
	}

	// 连接表的数据
	connList := make([]map[string]interface{}, 0)
	c.DB().Table(connectTable).Where(fmt.Sprintf("`%s` in (%s)", JoinForeignKey, toInSql(list, c.GetPrimary()))).Find(&connList)
	connMap := toMapList(connList, JoinForeignKey)
	// 目标表的数据
	gotList := make([]map[string]interface{}, 0)
	idKey := c.getTablePrimary(readTable)
	c.DB().Table(readTable).Where(fmt.Sprintf("`%s` in (%s)", idKey, toInSql(connList, JoinTargetKey))).Find(&gotList)
	gotMap := toMap(gotList, idKey)

	for i, m := range list {
		priVal := m[c.GetPrimary()]
		for _, got := range connMap[fmt.Sprintf("%v", priVal)] {
			gotID := fmt.Sprintf("%v", got[JoinTargetKey])
			if v, ok := gotMap[gotID]; ok {
				m[readTable] = v
			}
		}
		list[i] = m
	}
}

func toInSql(list []map[string]interface{}, key string) string {
	var inIds string
	uniqueValues := make(map[interface{}]bool)
	for _, m := range list {
		v := m[key]
		if v != nil && !uniqueValues[v] {
			if inIds != "" {
				inIds = inIds + ","
			}
			inIds = inIds + fmt.Sprintf("'%v'", v)
			uniqueValues[v] = true
		}
	}
	return inIds
}

// 转成map
func toMap(list []map[string]interface{}, key string) map[interface{}]map[string]interface{} {
	got := make(map[interface{}]map[string]interface{})
	for _, m := range list {
		got[fmt.Sprintf("%v", m[key])] = m
	}
	return got
}

// 转成map
func toMapList(list []map[string]interface{}, key string) map[string][]map[string]interface{} {
	got := make(map[string][]map[string]interface{})
	for _, m := range list {
		var k string = fmt.Sprintf("%v", m[key])
		if _, ok := got[k]; !ok {
			got[k] = make([]map[string]interface{}, 0)
		}
		got[k] = append(got[k], m)
	}
	return got
}

// 索引信息
type indexInfo struct {
	// 索引类型
	index string
	// 目标表模型的名称
	targetModelName string
	// 索引配置
	tags map[string]string
}

// 获取模型的索引信息
func getModelInfo(model interface{}) map[string]indexInfo {
	got := make(map[string]indexInfo)
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		gorm := field.Tag.Get("gorm")
		if gorm == "" {
			continue
		}
		info := indexInfo{tags: map[string]string{}}
		for _, s := range strings.Split(gorm, ";") {
			arr := strings.Split(s, ":")
			if len(arr) == 2 {
				info.tags[arr[0]] = arr[1]
				switch arr[0] {
				case "belongs_to", "has_one", "has_many":
					info.index = arr[0]
				case "many2many":
					info.index = arr[0]
					tt := field.Type.Elem().String()
					ttArr := strings.Split(tt, ".")
					info.targetModelName = ttArr[len(ttArr)-1]
				}
			}
		}
		if info.index != "" {
			got[field.Name] = info
		}
	}

	return got
}

func (c *CurdController) getTablePrimary(table string) string {
	// TODO 可以读取数据库表信息确认
	return "id"
}

func (c *CurdController) DB() *gorm.DB {
	return c.model.GetDB().Session(&gorm.Session{NewDB: true})
}

package amis

import (
	"testing"
)

func TestNewCurd(t *testing.T) {
	curd := NewCurd()

	curd.Column("自增", "id")
	curd.Column("文本", "text")
	curd.Column("图片", "image").Image()
	curd.Column("日期", "date").Date()
	curd.Column("进度", "progress").Progress()
	curd.Column("状态", "status").Status()
	curd.Column("开关", "switch").Switch()
	curd.Column("映射", "mapping").Mapping(map[string]string{})
	curd.Column("List", "list").List()
}

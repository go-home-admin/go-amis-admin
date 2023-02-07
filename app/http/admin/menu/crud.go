package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/entity/admin"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmAdminMenu())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.AutoGenerateFilter()
	curd.Column("ID", "id").SearchableInput()
	curd.Column("父级", "parent_id")
	curd.Column("排序", "order")
	curd.Column("组件名称", "name")
	curd.Column("组件", "component")
	curd.Column("地址", "path")
	curd.Column("重定向", "redirect")
	curd.Column("元数据", "meta").Json()
	curd.Column("排序", "sort")
	curd.Column("created_at", "created_at").Date()
	curd.Column("updated_at", "updated_at").Date()
}

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("parent_id", "父级").SaveInt()
	form.Input("name", "组件名称")
	form.Input("component", "组件")
	form.Input("path", "地址")
	form.Input("redirect", "重定向")
	form.EditorJson("meta", "元数据")
	form.Input("sort", "排序")
	form.Input("api_list", "api")
	form.InputDatetime("created_at", "created_at")
	form.InputDatetime("updated_at", "updated_at")

}

func (c *Controller) GinHandleCurd(ctx *gin.Context) {
	var crud = &CrudContext{}
	crud.CurdController.Context = ctx
	crud.CurdController.Crud = crud
	amis.GinHandleCurd(ctx, crud)
}

type CrudContext struct {
	amis.CurdController
}

package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/admin"
	"github.com/go-home-admin/go-admin/plugins/amis"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmAdminMenu())

	form := amis.NewForm()
	form.Input("n", "测试")
	form.Api = amis.Api{
		Method: "post",
		Url:    amis.GetUrl(c.Context),
	}
	c.GetPage().AddBody(form)
}

func (c *CrudContext) Table(curd *amis.Crud) {
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

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("text", "文本")
	form.Input("image", "图片")

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

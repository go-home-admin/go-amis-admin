package admin_user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/entity/admin"
	"github.com/go-home-admin/home/bootstrap/utils"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmAdminUsers())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.AutoGenerateFilter()
	curd.EnSelect()
	curd.Column("id", "id")
	curd.Column("账户", "username")
	curd.Column("显示名称", "name")
	curd.Column("头像", "avatar").Image().Height("50px")
	curd.Column("创建时间", "created_at").Date()
}

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("username", "账户")
	form.InputPassword("password", "密码").SkipEmpty().SaveMd5()
	form.Input("name", "显示名称")
	form.InputImage("avatar", "头像").Update("avatar")
	form.InputOptions("roles", "角色").SetModel(admin.NewOrmAdminRoles().Select("id as value, name as label"))
	form.AddCreatedAndUpdatedAt()

	form.SaveAfter(func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context) {
		roles, _ := post["roles"]

		utils.Dump(roles)
	})
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

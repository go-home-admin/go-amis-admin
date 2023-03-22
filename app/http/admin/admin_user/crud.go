package admin_user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app"
	"github.com/go-home-admin/go-admin/app/entity/admin"
	"github.com/sirupsen/logrus"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmAdminUsers())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.EnOnlySelect()
	curd.Column("id", "id").SearchableInput()
	curd.Column("账户", "username").SearchableInput()
	curd.Column("显示名称", "name")
	curd.Column("所属角色", "admin_roles.name")
	curd.Column("头像", "avatar").Image().Height("50px")
	curd.Column("创建时间", "created_at").Date()
}

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("username", "账户")
	form.InputPassword("password", "密码").SkipEmpty().SaveMd5()
	form.Input("name", "显示名称")
	form.InputImage("avatar", "头像").Update("avatar")
	form.InputOptions("admin_roles.id", "角色").SetModel(admin.NewOrmAdminRoles().Select("id as value, name as label"))
	form.AddCreatedAndUpdatedAt()
	// 同时更新依赖表
	form.SaveAfter(func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context) {
		roles, has := post["admin_roles"].(map[string]interface{})["id"]
		userID := app.Int32(primaryVal)
		admin.NewOrmAdminRoleUsers().WhereUserId(userID).Delete()
		if has {
			if role, ok := roles.(string); ok {
				if err := admin.NewOrmAdminRoleUsers().Insert(&admin.AdminRoleUsers{
					RoleId: app.Int32(role),
					UserId: userID,
				}); err != nil {
					logrus.Error(err)
				}
			}
		}
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

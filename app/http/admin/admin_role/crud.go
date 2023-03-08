package admin_role

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/entity/admin"
)

func (c *CrudContext) Common() {
	c.SetDb(admin.NewOrmAdminRoles())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.Column("id", "id")
	curd.Column("角色名", "name")
	curd.Column("备注", "remark")
	curd.Column("默认权限", "slug")
	curd.Column("created_at", "created_at")
	curd.Column("updated_at", "updated_at")
}

func (c *CrudContext) Form(form *amis.Form) {
	form.Input("name", "角色名")
	form.Input("slug", "默认权限")
	form.Textarea("remark", "备注")
	form.AddCreatedAt()
	form.AddUpdatedAt()
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

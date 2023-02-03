package amis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	"gorm.io/gorm"
)

type Index interface {
	Index(ctx *gin.Context)
}
type Create interface {
	Create(ctx *gin.Context)
}
type Update interface {
	Update(ctx *gin.Context)
}
type Delete interface {
	Delete(ctx *gin.Context)
}

type CurdSave interface {
	Common()
	Table(curd *Crud)
	Form(form *Form)
}

func NewCrudDisplay() *CurdController {
	return &CurdController{}
}

type CurdController struct {
	*gin.Context

	page  *Page
	title string
	model Model
	Crud  CurdSave
}

type Model interface {
	TableName() string
	GetDB() *gorm.DB
}

func (c *CurdController) GetPage() *Page {
	if c.page == nil {
		c.page = NewPage("")
	}
	return c.page
}

func (c *CurdController) SetTitle(t string) {
	c.title = t
}

func (c *CurdController) SetDb(model Model) {
	c.model = model
}

func GinHandleCurd(ctx *gin.Context, controller CurdSave) {
	c := &CurdController{
		Context: ctx,
		Crud:    controller,
	}
	controller.Common()

	switch ctx.Param("action") {
	case "", "index": // 列表页面
		if i, ok := controller.(Index); ok {
			i.Index(ctx)
		} else {
			c.Index(ctx)
		}
	case "list": // 列表数据
		c.List(ctx)
	case "create": // 创建
		if i, ok := controller.(Create); ok {
			i.Create(ctx)
		} else {
			c.Create(ctx)
		}
	case "edit": // edit
		if i, ok := controller.(Update); ok {
			i.Update(ctx)
		} else {
			c.Update(ctx)
		}
	case "del": // edit
		if i, ok := controller.(Delete); ok {
			i.Delete(ctx)
		} else {
			c.Delete(ctx)
		}
	default:
		http.NewContext(ctx).Fail(errors.New("不支持的路由"))
	}
}

// Index 默认还在页面信息
func (c *CurdController) Index(ctx *gin.Context) {
	context := http.NewContext(ctx)

	crud := NewCurd(ctx)
	form := NewForm()

	c.Crud.Table(crud)
	c.Crud.Form(form)

	// 把form放入到curd的按钮中
	crud.AddCreate(form)
	crud.Operation().AddButton("编辑").SetDialog(form)
	crud.Operation().AddButton("删除").SetClassName("text-danger").SetAjax("确定要删除？", ctx.Request.RequestURI).Method = "delete"

	page := c.GetPage()
	page.AddBody(crud.ToAmisJson())

	context.Success(page)
}

func (c *CurdController) List(ctx *gin.Context) {
	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Create(ctx *gin.Context) {

	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Update(ctx *gin.Context) {

	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Delete(ctx *gin.Context) {

	http.NewContext(ctx).Success(nil)
}

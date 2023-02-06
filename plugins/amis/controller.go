package amis

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Index interface {
	Index(ctx *gin.Context)
}
type List interface {
	List(ctx *gin.Context)
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

type GetPrimary interface {
	GetPrimary() string
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

func (c *CurdController) GetPrimary() string {
	return "id"
}

func (c *CurdController) GetFromData(ctx *gin.Context) (map[string]interface{}, map[string]interface{}) {
	by, _ := ctx.GetRawData()
	var m map[string]interface{}
	err := json.Unmarshal(by, &m)
	if err != nil {
		logrus.Error(err)
		return nil, nil
	}
	data := map[string]interface{}{}

	return data, m
}

func GinHandleCurd(ctx *gin.Context, controller CurdSave) {
	c := &CurdController{
		Context: ctx,
		Crud:    controller,
	}
	controller.Common()
	key := ctx.Param("action")
	if key == "" {
		c.CurdAll()
	} else {
		c.CurdOne(key)
	}
}

func (c *CurdController) CurdAll() {
	switch c.Request.Method {
	case "POST":
		// 创建
		if i, ok := c.Crud.(Create); ok {
			i.Create(c.Context)
		} else {
			c.Create(c.Context)
		}
	default:
		// 页面显示
		if i, ok := c.Crud.(Index); ok {
			i.Index(c.Context)
		} else {
			c.Index(c.Context)
		}
	}
}

func (c *CurdController) CurdOne(action string) {
	ctx := c.Context
	controller := c.Crud
	switch action {
	case "list": // 列表数据
		if i, ok := controller.(List); ok {
			i.List(ctx)
		}
	case "edit": // edit
		if i, ok := controller.(Update); ok {
			i.Update(ctx)
		}
	case "del": // edit
		if i, ok := controller.(Delete); ok {
			i.Delete(ctx)
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
	delUrl := GetUrl(ctx, "/del?${") + c.GetPrimary() + "}"
	crud.Operation().AddButton("删除").SetClassName("text-danger").SetAjax("确定要删除？", delUrl).Method = "delete"

	page := c.GetPage()
	page.AddBody(crud.ToAmisJson())

	context.Success(page)
}

func (c *CurdController) List(ctx *gin.Context) {
	got := NewCurdData()
	list := make([]map[string]interface{}, 0)
	c.model.GetDB().Count(&got.Total)
	if got.Total > 0 {
		Page := GetInt(c.Context, "page", 1)
		PageSize := GetInt(c.Context, "pageSize", 20)
		tx := c.model.GetDB().Offset((Page - 1) * PageSize).Limit(PageSize).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
		// TODO 后端值转换

		for _, m := range list {
			got.Items = append(got.Items, m)
		}
	}

	http.NewContext(ctx).Success(got)
}

func (c *CurdController) Create(ctx *gin.Context) {
	by, _ := ctx.GetRawData()
	var data map[string]interface{}
	err := json.Unmarshal(by, &data)
	if err != nil {
		logrus.Error(err)
		return
	}

	td := c.model.GetDB().Create(&data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("创建失败"))
		return
	}
	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Update(ctx *gin.Context) {
	data, all := c.GetFromData(ctx)
	primary := c.GetPrimary()
	primaryValStringOrFloat64, ok := all[primary]
	if !ok {
		logrus.Error("必须要有主键数据才能更新, 当前的主建=" + primary)
		return
	}
	var primaryVal interface{}
	switch primaryValStringOrFloat64.(type) {
	case string:
		primaryVal = primaryValStringOrFloat64
	case float64:
		primaryVal = int(primaryValStringOrFloat64.(float64))
	}

	td := c.model.GetDB().Where(primary+" = ?", primaryVal).Updates(&data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("更新失败"))
		return
	}

	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Delete(ctx *gin.Context) {
	primary := c.GetPrimary()
	var primaryVal interface{}
	primaryVal, ok := ctx.GetQuery(primary)
	if !ok {
		logrus.Error("url必须要有主键数据才能删除, 当前的主建=" + primary)
		return
	}

	td := c.model.GetDB().Delete(c.model, primaryVal)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("删除失败"))
		return
	}

	http.NewContext(ctx).Success(nil)
}

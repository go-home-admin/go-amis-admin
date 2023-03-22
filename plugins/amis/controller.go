package amis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
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
	GetTableInfo() interface{}
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

// GetFromData 预处理提交参数
func (c *CurdController) GetFromData(ctx *gin.Context, form *Form) (interface{}, map[string]interface{}, error) {
	by, _ := ctx.GetRawData()
	post := map[string]interface{}{}
	err := json.Unmarshal(by, &post)
	if err != nil {
		logrus.Error(err)
		return nil, post, err
	}
	data := form.data
	for _, item := range form.Items() {
		v := item.GetValue(post)
		if v != nil {
			data[item.GetName()] = v
		}
	}
	postStr, _ := json.Marshal(data)
	m := c.model.GetTableInfo()
	_ = json.Unmarshal(postStr, &m)
	return m, post, nil
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

// 便捷闭包函数
const _route_ = "_route_"

func AddCallRoutes(ctx *gin.Context, action string, call func(ctx *gin.Context)) {
	var routes map[string]func(ctx2 *gin.Context)
	t, has := ctx.Get(_route_)
	if has {
		routes = t.(map[string]func(ctx2 *gin.Context))
	} else {
		routes = map[string]func(ctx2 *gin.Context){}
	}
	routes[action] = call
	ctx.Set(_route_, routes)
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
		// 自定义的路由
		form := NewForm(ctx)
		c.Crud.Form(form)
		routes, has := ctx.Get(_route_)
		if has {
			if displayRoutes, ok := routes.(map[string]func(*gin.Context)); ok {
				if call, ok := displayRoutes[action]; ok {
					call(ctx)
					return
				}
			}
		}

		http.NewContext(ctx).Fail(errors.New("不支持的路由"))
	}
}

// Index 默认还在页面信息
func (c *CurdController) Index(ctx *gin.Context) {
	context := http.NewContext(ctx)

	crud := NewCurd(ctx)
	form := NewForm(ctx)

	c.Crud.Table(crud)
	c.Crud.Form(form)

	// 把form放入到curd的按钮中
	priStr := c.GetPrimary() + "=${" + c.GetPrimary() + "}"
	crud.AddCreate(form)
	crud.Operation().AddButton("编辑").SetDialogForm(form.SetApi(GetUrl(ctx, "/edit?"+priStr), "put"))
	delUrl := GetUrl(ctx, "/del?"+priStr)
	crud.Operation().AddButton("删除").SetClassName("text-danger").SetAjax("确定要删除？", delUrl).Method = "delete"

	page := c.GetPage()
	page.AddBody(crud.ToAmisJson())

	context.Success(page)
}

func (c *CurdController) List(ctx *gin.Context) {
	got := NewCurdData()
	got.Page = GetInt(c.Context, "page", 1)
	list := make([]map[string]interface{}, 0)

	queryDB := c.model.GetDB().Session(&gorm.Session{})
	queryDB.Count(&got.Total)
	if got.Total > 0 {
		crud := NewCurd(ctx)
		c.Crud.Table(crud)

		PageSize := GetInt(c.Context, "perPage", 20)
		queryDB = queryDB.Offset((got.Page - 1) * PageSize).Limit(PageSize)
		if crud.enSelect {
			query := ""
			// 只读取设置了columns才读取数据库
			for _, columnT := range crud.columns {
				if column, ok := columnT.(*ColumnConfig); ok && !column.skip {
					if strings.Index(column.Name, ".") == -1 {
						query += column.Name + ","
					}
				}
			}
			queryDB = queryDB.Select(strings.Trim(query, ","))
		}
		tx := queryDB.Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
		// 自动查询分表数据
		c.authSelect(list, crud.columns)
		// 后端值转换, 后台类型的状态转换登陆
		for _, m := range list {
			for _, columnT := range crud.columns {
				if column, ok := columnT.(*ColumnConfig); ok {
					if column.display != nil {
						m[column.Name] = column.display(m[column.Name])
					}
				}
			}
			got.Items = append(got.Items, m)
		}
	}

	http.NewContext(ctx).Success(got)
}

// Create 创建数据
func (c *CurdController) Create(ctx *gin.Context) {
	form := NewForm(ctx)
	c.Crud.Form(form)
	for _, f := range form.createBefore {
		f(form)
	}
	data, post, _ := c.GetFromData(ctx, form)
	td := c.model.GetDB().Create(data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("创建失败"))
		return
	}
	for _, f := range form.createAfter {
		_, primaryVal := getPrimaryKey(data)
		f(primaryVal, post, ctx)
	}

	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Update(ctx *gin.Context) {
	form := NewForm(ctx)
	c.Crud.Form(form)
	if form.updateBefore != nil {

	}
	data, post, _ := c.GetFromData(ctx, form)
	key, primaryVal := getPrimaryKey(data)
	if key == "" {
		http.NewContext(ctx).Fail(errors.New("curl只能更新有自增字段的模型"))
		return
	}

	if primaryVal == nil || primaryVal == "" || primaryVal == 0 || fmt.Sprintf("%v", primaryVal) == "0" || fmt.Sprintf("%v", primaryVal) == "" {
		primaryVal = ctx.Query(key)
		if primaryVal == "" {
			logrus.Error("必须要有主键数据才能更新, 当前的主建=" + key)
			return
		}
	}

	td := c.model.GetDB().Where(key+" = ?", primaryVal).Updates(data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("更新失败"))
		return
	}

	for _, f := range form.updateAfter {
		f(primaryVal, post, ctx)
	}

	http.NewContext(ctx).Success(nil)
}

func (c *CurdController) Delete(ctx *gin.Context) {
	primary, primaryVal := getPrimaryKey(c.model.GetTableInfo())
	primaryVal, ok := ctx.GetQuery(primary)
	if !ok {
		logrus.Error("url必须要有主键数据才能删除, 当前的主建=" + primary)
		return
	}

	td := c.model.GetDB().Delete(c.model.GetTableInfo(), primaryVal)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("删除失败"))
		return
	}

	http.NewContext(ctx).Success(nil)
}

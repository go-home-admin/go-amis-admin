package amis

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/database"
	"github.com/sirupsen/logrus"
	"strconv"
)

func NewForm(ctx *gin.Context) *Form {
	return &Form{
		ctx:      ctx,
		Type:     "form",
		Body:     make([]interface{}, 0),
		itemList: make(map[string]IsFormItem),
		data:     map[string]interface{}{},
	}
}

type Form struct {
	ctx   *gin.Context
	Id    string        `json:"id,omitempty"`
	Type  string        `json:"type,omitempty"`
	Title string        `json:"title,omitempty"`
	Body  []interface{} `json:"body,omitempty"`
	// 水平模式horizontal 内联模式inline
	Mode string `json:"mode,omitempty"`
	// 如果水平模式, 还可以设置
	Horizontal *FormModeHorizontal `json:"horizontal,omitempty"`

	// 初始化表单
	InitApi string `json:"initApi,omitempty"`
	// 如果需要设置提交连接
	Api Url `json:"api,omitempty"`

	itemList map[string]IsFormItem
	// 如果包装到dialog, 可以设置到dialog组件
	size string
	// 默认更新的值
	data map[string]interface{}

	createBefore []func(form *Form)
	updateBefore []func(form *Form)

	createAfter []func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context)
	updateAfter []func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context)
}

type FormModeHorizontal struct {
	LeftFixed string `json:"leftFixed"`
}

// CreateBefore 创建前执行
func (f *Form) CreateBefore(fun func(form *Form)) *Form {
	if f.createBefore == nil {
		f.createBefore = make([]func(form *Form), 0)
	}
	f.createBefore = append(f.createBefore, fun)
	return f
}

// UpdateBefore 更新前执行
func (f *Form) UpdateBefore(fun func(form *Form)) *Form {
	if f.createBefore == nil {
		f.updateBefore = make([]func(form *Form), 0)
	}
	f.updateBefore = append(f.updateBefore, fun)
	return f
}

// SaveAfter 创建或者更新后执行
func (f *Form) SaveAfter(fun func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context)) *Form {
	f.CreateAfter(fun)
	f.UpdateAfter(fun)
	return f
}

// CreateAfter 创建后执行
func (f *Form) CreateAfter(fun func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context)) *Form {
	if f.createBefore == nil {
		f.createAfter = make([]func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context), 0)
	}
	f.createAfter = append(f.createAfter, fun)
	return f
}

// UpdateAfter 更新后执行
func (f *Form) UpdateAfter(fun func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context)) *Form {
	if f.createBefore == nil {
		f.updateAfter = make([]func(primaryVal interface{}, post map[string]interface{}, ctx *gin.Context), 0)
	}
	f.updateAfter = append(f.updateAfter, fun)
	return f
}

// SetData 设置默认值
func (f *Form) SetData(i map[string]interface{}) *Form {
	f.data = i
	return f
}

func (f *Form) AddData(k string, v interface{}) *Form {
	f.data[k] = v
	return f
}

func (f *Form) SetSize(i string) *Form {
	switch i {
	case "sm", "lg", "xl", "full":
		f.size = i
	default:
		logrus.Warning("SetSize = sm | lg | xl | full")
	}
	return f
}

func (f *Form) GetDialogSize() string {
	if f.size == "" {
		f.size = "lg"
	}
	return f.size
}

func (f *Form) AddBody(i interface{}) {
	f.Body = append(f.Body, i)

	if v, ok := i.(IsFormItem); ok {
		f.itemList[v.GetName()] = v
	}
}

func (f *Form) SetApi(url Url, method string) *Form {
	newForm := *f
	url.Method(method)
	newForm.Api = url

	return &newForm
}

func (f *Form) Items() map[string]IsFormItem {
	return f.itemList
}

func (f *Form) AddCreatedAndUpdatedAt() {
	f.AddCreatedAt()
	f.AddUpdatedAt()
}

// AddCreatedAt 添加创建时间字段
func (f *Form) AddCreatedAt() {
	f.CreateBefore(func(form *Form) {
		form.AddData("created_at", database.Now())
	})
}

// AddUpdatedAt 添加更新时间字段
func (f *Form) AddUpdatedAt() {
	f.UpdateBefore(func(form *Form) {
		form.AddData("updated_at", database.Now())
	})
}

type IsFormItem interface {
	GetName() string
	GetValue(row map[string]interface{}) interface{}
}

type FormItem struct {
	IsFormItem `json:"-"`
	Id         string `json:"id,omitempty"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	Name       string `json:"name"`

	save []func(old interface{}) interface{}
	opt  map[string]interface{}
}

func (f *FormItem) SetOptions(k string, v interface{}) {
	f.opt[k] = v
}

// MarshalJSON opt 会合并到整个ColumnConfig上再输出到前端
func (f *FormItem) MarshalJSON() ([]byte, error) {
	newStruct := *f
	if len(f.opt) == 0 {
		return json.Marshal(newStruct)
	}
	mm := map[string]interface{}{}
	by, _ := json.Marshal(newStruct)
	_ = json.Unmarshal(by, &mm)
	for k, v := range f.opt {
		mm[k] = v
	}
	return json.Marshal(mm)
}

func NewItem(name, label, Type string) *FormItem {
	return &FormItem{
		Id:    "",
		Label: label,
		Type:  Type,
		Name:  name,
		opt:   map[string]interface{}{},
	}
}

func (f *FormItem) GetName() string {
	return f.Name
}

func (f *FormItem) GetValue(row map[string]interface{}) interface{} {
	var got interface{}
	v, has := row[f.Name]
	if has {
		got = v
	}
	if f.save != nil {
		for _, f2 := range f.save {
			if got != nil {
				got = f2(got)
			}
		}
	}
	return got
}

func (f *FormItem) Placeholder(v string) *FormItem {
	f.SetOptions("placeholder", v)
	return f
}

// SkipEmpty 不保存空的值
func (f *FormItem) SkipEmpty() *FormItem {
	f.SetSave(func(old interface{}) interface{} {
		if old != nil {
			switch old.(type) {
			case string:
				if old.(string) == "" {
					return nil
				}
			}
		}
		return old
	})
	return f
}

// NotSave 不保存, 只是显示组件
func (f *FormItem) NotSave() *FormItem {
	f.SetSave(func(old interface{}) interface{} {
		return nil
	})
	return f
}

func (f *FormItem) SetSave(fun func(old interface{}) interface{}) {
	f.save = append(f.save, fun)
}

func (f *FormItem) SaveInt() {
	f.SetSave(func(old interface{}) interface{} {
		got := 0
		var err error
		switch old.(type) {
		case string:
			got, err = strconv.Atoi(old.(string))
			if err != nil {
				logrus.Error(err)
			}
		case float64: // 经过json后可能是float64
			got = int(old.(float64))
		}

		return got
	})
}

// SaveMd5 使用md5保存
func (f *FormItem) SaveMd5(sec ...string) {
	f.SetSave(func(old interface{}) interface{} {
		if s, ok := old.(string); ok {
			secStr := ""
			for _, s2 := range sec {
				secStr += s2
			}
			data := []byte(s + secStr)
			return fmt.Sprintf("%x", md5.Sum(data))
		}
		return ""
	})
}

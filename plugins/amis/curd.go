package amis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/database"
	"time"
)

type CurdData struct {
	Items []interface{} `json:"items"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
}

func NewCurdData() *CurdData {
	return &CurdData{
		Items: make([]interface{}, 0),
		Total: 0,
		Page:  0,
	}
}

func (c *CurdData) SetItems(items []interface{}) {
	c.Items = items
}

func (c *CurdData) AddItems(item interface{}) {
	c.Items = append(c.Items, item)
}

func NewCurd(ctx *gin.Context) *Crud {
	return &Crud{
		ctx:     ctx,
		columns: make([]interface{}, 0),
		opt:     map[string]interface{}{},
	}
}

// Crud 这个对象会存储一些后台内部的配置，不能直接作为amis结构
type Crud struct {
	ctx           *gin.Context
	columns       []interface{}
	headerToolbar []interface{}
	// 操作列的索引
	operation int
	opt       map[string]interface{}
}

func (c *Crud) SetOptions(k string, v interface{}) {
	c.opt[k] = v
}

// CurdJsonConfig 这个对象直接响应到前端json
type CurdJsonConfig struct {
	Type          string        `json:"type,omitempty"`
	Api           string        `json:"api,omitempty"`
	SyncLocation  bool          `json:"syncLocation,omitempty"`
	Columns       []interface{} `json:"columns,omitempty"`
	HeaderToolbar []interface{} `json:"headerToolbar,omitempty"`
	opt           map[string]interface{}
}

func (p *CurdJsonConfig) MarshalJSON() ([]byte, error) {
	newStruct := *p
	if len(p.opt) == 0 {
		return json.Marshal(newStruct)
	}
	mm := map[string]interface{}{}
	by, _ := json.Marshal(newStruct)
	_ = json.Unmarshal(by, &mm)
	for k, v := range p.opt {
		mm[k] = v
	}
	return json.Marshal(mm)
}

func (c *Crud) ToAmisJson() *CurdJsonConfig {
	got := &CurdJsonConfig{
		Type:          "crud",
		SyncLocation:  false,
		HeaderToolbar: c.headerToolbar,
		Api:           GetUrl(c.ctx, "/list"),
		opt:           c.opt,
	}
	got.Columns = c.columns

	return got
}

func (c *Crud) AutoGenerateFilter() {
	c.SetOptions("autoGenerateFilter", true)
}

func (c *Crud) Column(label string, name string) *ColumnConfig {
	config := &ColumnConfig{
		curl:  c,
		Name:  name,
		Label: label,
		Type:  "text",
		opt:   map[string]interface{}{},
	}
	c.columns = append(c.columns, config)
	return config
}

// Operation 添加操作栏
func (c *Crud) Operation() *OperationConfig {
	if c.operation == 0 {
		config := &OperationConfig{
			Label:   "操作",
			Type:    "operation",
			Buttons: []*Button{},
		}
		c.columns = append(c.columns, config)
		c.operation = len(c.columns) - 1
		return config
	}

	return c.columns[c.operation].(*OperationConfig)
}

// AddCreate 创建按钮
func (c *Crud) AddCreate(form *Form) {
	f := form.SetApi(GetUrl(c.ctx, ""), "post")
	button := NewButton("创建").SetDialogForm(f)
	button.Level = "primary"
	c.headerToolbar = []interface{}{
		button,
		"bulkActions",
	}
}

type OperationConfig struct {
	Id      string    `json:"id,omitempty"`
	Type    string    `json:"type,omitempty"`
	Label   string    `json:"label,omitempty"`
	Buttons []*Button `json:"buttons,omitempty"`
}

func (o *OperationConfig) AddButton(label string) *Button {
	b := NewButton(label)
	o.Buttons = append(o.Buttons, b)
	return b
}

// ColumnConfig 如果不需要响应到前端要加-
type ColumnConfig struct {
	curl *Crud

	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
	// 宽度
	Width string `json:"width,omitempty" form:"width"`

	Searchable interface{} `json:"searchable,omitempty"`
	Format     string      `json:"format,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	// opt 会合并到整个ColumnConfig上再输出到前端
	opt     map[string]interface{}
	display func(v interface{}) interface{}
}

// MarshalJSON opt 会合并到整个ColumnConfig上再输出到前端
func (c *ColumnConfig) MarshalJSON() ([]byte, error) {
	newStruct := *c
	if len(c.opt) == 0 {
		return json.Marshal(newStruct)
	}
	mm := map[string]interface{}{}
	by, _ := json.Marshal(newStruct)
	_ = json.Unmarshal(by, &mm)
	for k, v := range c.opt {
		mm[k] = v
	}
	return json.Marshal(mm)
}

func (c *ColumnConfig) Display(f func(v interface{}) interface{}) *ColumnConfig {
	c.display = f
	return c
}

func (c *ColumnConfig) Image() *ColumnConfig {
	c.Type = "image"
	return c
}

// Date 输出到前端，应该要时间戳，方便其他组件读取
func (c *ColumnConfig) Date(formats ...string) *ColumnConfig {
	c.Type = "date"
	if len(formats) == 0 {
		c.Format = "YYYY年MM月DD日 HH时mm分ss秒"
		c.Width = "160"
	} else {
		c.Format = formats[0]
	}
	c.Display(func(v interface{}) interface{} {
		switch v.(type) {
		case time.Time:
			return v.(time.Time).Unix()
		}
		return database.StrToTime(fmt.Sprintf("%v", v)).Unix()
	})
	return c
}

func (c *ColumnConfig) Progress() *ColumnConfig {
	c.Type = "progress"
	return c
}

func (c *ColumnConfig) Status() *ColumnConfig {
	c.Type = "status"
	return c
}

func (c *ColumnConfig) Switch() *ColumnConfig {
	c.Type = "switch"
	return c
}

func (c *ColumnConfig) Mapping(list map[string]string) *ColumnConfig {
	c.Type = "mapping"
	return c
}
func (c *ColumnConfig) List() *ColumnConfig {
	c.Type = "list"
	return c
}

// SearchableInput 自动生成查询
// SearchableInput(label, name...)
func (c *ColumnConfig) SearchableInput(opts ...string) *FormItemText {
	name := c.Name
	label := c.Label
	switch len(opts) {
	case 2:
		name = opts[1]
		label = opts[0]
	case 1:
		label = opts[0]
	}

	item := &FormItemText{
		FormItem: NewItem(name, label, "input-text"),
	}
	c.Type = ""
	c.Searchable = item
	return item
}

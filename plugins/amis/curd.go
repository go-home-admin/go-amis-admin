package amis

import "github.com/gin-gonic/gin"

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
	}
}

// Crud 这个对象会存储一些后台内部的配置，不能直接作为amis结构
type Crud struct {
	ctx           *gin.Context
	columns       []interface{}
	headerToolbar []interface{} `json:"headerToolbar,omitempty"`
	// 操作列的索引
	operation int
}

// CurdJsonConfig 这个对象直接响应到前端json
type CurdJsonConfig struct {
	Type          string        `json:"type,omitempty"`
	Api           string        `json:"api,omitempty"`
	SyncLocation  bool          `json:"syncLocation,omitempty"`
	Columns       []interface{} `json:"columns,omitempty"`
	HeaderToolbar []interface{} `json:"headerToolbar,omitempty"`
}

func (c *Crud) ToAmisJson() CurdJsonConfig {
	got := CurdJsonConfig{
		Type:          "crud",
		SyncLocation:  false,
		HeaderToolbar: c.headerToolbar,
		Api:           GetUrl(c.ctx, "/list"),
	}
	got.Columns = c.columns

	return got
}

func (c *Crud) Column(label string, name string) *ColumnConfig {
	config := &ColumnConfig{
		Name:  name,
		Label: label,
		Type:  "text",
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
	f := *form
	f.Api = Api{
		Method: "post",
		Url:    GetUrl(c.ctx, ""),
	}
	page := NewPage("创建")
	page.Body = f
	button := NewButton(page.Title).SetDialog(page)
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
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
	// 宽度
	Width string `json:"width,omitempty" form:"width"`

	Searchable interface{} `json:"searchable,omitempty"`
}

func (c *ColumnConfig) Image() *ColumnConfig {
	c.Type = "image"
	return c
}

func (c *ColumnConfig) Date() *ColumnConfig {
	c.Type = "date"
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
func (c *ColumnConfig) SearchableInput(opts ...string) *FormItemText {
	name := c.Name
	label := c.Label
	switch len(opts) {
	case 2:
		name = opts[0]
		label = opts[1]
	case 1:
		name = opts[0]
	}

	item := &FormItemText{
		FormItem: NewItem(name, label, "input-text"),
	}
	c.Searchable = item
	return item
}

package amis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func NewForm() *Form {
	return &Form{
		Type:     "form",
		Body:     make([]interface{}, 0),
		itemList: make(map[string]IsFormItem),
		data:     map[string]interface{}{},
	}
}

type Form struct {
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
	Api Api `json:"api,omitempty"`

	itemList map[string]IsFormItem
	// 如果包装到dialog, 可以设置到dialog组件
	size string
	// 默认更新的值
	data map[string]interface{}
}

type FormModeHorizontal struct {
	LeftFixed string `json:"leftFixed"`
}

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

func (f *Form) SetApi(url, method string) *Form {
	newForm := *f
	newForm.Api = Api{
		Method: method,
		Url:    url,
	}

	return &newForm
}

func (f *Form) Items() map[string]IsFormItem {
	return f.itemList
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

	save func(old interface{}) interface{}
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
		got = f.save(got)
	}
	return got
}

func (f *FormItem) Placeholder(v string) *FormItem {
	f.SetOptions("placeholder", v)
	return f
}

func (f *FormItem) SetSave(fun func(old interface{}) interface{}) {
	f.save = fun
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

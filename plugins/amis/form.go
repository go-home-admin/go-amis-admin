package amis

func NewForm() *Form {
	return &Form{
		Type: "form",
		Body: make([]interface{}, 0),
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
}

type FormModeHorizontal struct {
	LeftFixed string `json:"leftFixed"`
}

func (f *Form) AddBody(i interface{}) {
	f.Body = append(f.Body, i)
}

type FormItem struct {
	Id    string `json:"id,omitempty"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

func NewItem(name, label, Type string) *FormItem {
	return &FormItem{
		Id:    "",
		Label: label,
		Type:  Type,
		Name:  name,
	}
}

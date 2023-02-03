package amis

type AmisPage interface {
	AddBody(i interface{})
}

type Api struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

func NewPage(title string) *Page {
	return &Page{
		Id:    "",
		Type:  "",
		Title: title,
		Body:  nil,
		Data:  nil,
	}
}

type Page struct {
	Id      string      `json:"id,omitempty"`
	Type    string      `json:"type,omitempty"`
	Title   string      `json:"title,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	OnEvent *OnEvent    `json:"onEvent,omitempty"`
	InitApi *InitApi    `json:"initApi,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type InitApi struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type OnEvent struct {
	Init struct {
		Weight  int           `json:"weight"`
		Actions []interface{} `json:"actions"`
	} `json:"init"`
	Inited struct {
		Weight  int           `json:"weight"`
		Actions []interface{} `json:"actions"`
	} `json:"inited"`
}

func (p *Page) AddBody(i interface{}) {
	if _, ok := p.Body.([]interface{}); !ok {
		p.Body = []interface{}{}
	}
	p.Body = append(p.Body.([]interface{}), i)
}

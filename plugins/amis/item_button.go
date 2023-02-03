package amis

func NewButton(label string) *Button {
	return &Button{
		Label:      label,
		Type:       "button",
		ActionType: "dialog",
		Level:      "primary",
	}
}

type Button struct {
	Id          string      `json:"id,omitempty"`
	Label       string      `json:"label"`
	Type        string      `json:"type,omitempty"`
	ActionType  string      `json:"actionType,omitempty"`
	Level       string      `json:"level,omitempty"`
	Dialog      interface{} `json:"dialog,omitempty"`
	ClassName   string      `json:"className,omitempty"`
	ConfirmText string      `json:"confirmText,omitempty"`
	Api         *Api        `json:"api,omitempty"`
}

func (b *Button) SetClassName(v string) *Button {
	b.ClassName = v
	return b
}

func (b *Button) SetDialog(page interface{}) *Button {
	b.ActionType = "dialog"
	b.Dialog = page
	return b
}

func (b *Button) SetAjax(confirmText, url string) *Api {
	b.ConfirmText = confirmText
	b.ActionType = "ajax"
	b.Api = &Api{
		Method: "post",
		Url:    url,
	}
	return b.Api
}

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
	Api         Url         `json:"api,omitempty"`
}

func (b *Button) SetClassName(v string) *Button {
	b.ClassName = v
	return b
}

func (b *Button) SetAjax(confirmText string, url Url) Url {
	b.ConfirmText = confirmText
	b.ActionType = "ajax"
	b.Api = url
	return b.Api
}

func (b *Button) SetDialog(page interface{}) *Button {
	b.ActionType = "dialog"
	b.Dialog = page
	return b
}

func (b *Button) SetDialogForm(f *Form) *Button {
	b.ActionType = "dialog"
	page := NewPageForm(f)
	page.Title = b.Label
	page.SetOptions("size", f.GetDialogSize())
	b.Dialog = page
	return b
}

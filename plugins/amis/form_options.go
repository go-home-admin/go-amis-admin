package amis

func (f *Form) Options(name, label string) *FormItemOptions {
	item := &FormItemOptions{
		FormItem: NewItem(name, label, "select"),
	}
	f.AddBody(item)
	return item
}

type FormItemOptions struct {
	*FormItem
	Options []Options `json:"options"`
}

type Options struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

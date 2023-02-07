package amis

func (f *Form) JsonSchema(name, label string) *FormItemText {
	item := &FormItemText{
		FormItem: NewItem(name, label, "json-schema"),
	}
	f.AddBody(item)
	return item
}

type FormItemJsonSchema struct {
	*FormItem
}

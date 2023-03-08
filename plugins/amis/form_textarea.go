package amis

func (f *Form) Textarea(name, label string) *FormItemTextarea {
	item := &FormItemTextarea{
		FormItem: NewItem(name, label, "textarea"),
	}
	f.AddBody(item)
	return item
}

type FormItemTextarea struct {
	*FormItem
}

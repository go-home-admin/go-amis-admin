package amis

func (f *Form) Input(name, label string) *FormItemText {
	item := &FormItemText{
		FormItem: NewItem(name, label, "input-text"),
	}
	f.AddBody(item)
	return item
}

type FormItemText struct {
	*FormItem
}

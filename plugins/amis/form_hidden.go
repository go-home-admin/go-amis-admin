package amis

func (f *Form) Hidden(name, label string) *FormItemText {
	item := &FormItemText{
		FormItem: NewItem(name, label, "hidden"),
	}
	f.AddBody(item)
	return item
}

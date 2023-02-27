package amis

func (f *Form) Password(name, label string) *FormItemText {
	item := &FormItemText{
		FormItem: NewItem(name, label, "input-password"),
	}
	f.AddBody(item)
	return item
}

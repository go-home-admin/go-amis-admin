package amis

func (f *Form) InputPassword(name, label string) *FormItemText {
	item := &FormItemText{
		FormItem: NewItem(name, label, "input-password"),
	}
	f.AddBody(item)
	return item
}

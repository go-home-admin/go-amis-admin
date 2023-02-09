package amis

func (f *Form) InputNumber(name, label string) *FormItemNumber {
	item := &FormItemNumber{
		FormItem: NewItem(name, label, "input-number"),
	}
	f.AddBody(item)
	return item
}

type FormItemNumber struct {
	*FormItem
}

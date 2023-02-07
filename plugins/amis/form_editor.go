package amis

func (f *Form) Editor(name, label string) *FormItemEditor {
	item := &FormItemEditor{
		FormItem: NewItem(name, label, "editor"),
	}
	f.AddBody(item)
	return item
}

// EditorJson 编辑json
func (f *Form) EditorJson(name, label string) *FormItemEditor {
	item := &FormItemEditor{
		FormItem: NewItem(name, label, "editor"),
	}
	item.Language("json")
	f.AddBody(item)
	return item
}

type FormItemEditor struct {
	*FormItem
}

func (f *FormItemEditor) Language(v string) *FormItemEditor {
	f.opt["language"] = v
	return f
}

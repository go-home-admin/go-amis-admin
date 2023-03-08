package amis

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func (f *Form) InputOptions(name, label string) *FormItemOptions {
	item := &FormItemOptions{
		FormItem: NewItem(name, label, "select"),
	}
	f.AddBody(item)
	return item
}

type FormItemOptions struct {
	*FormItem
}

type Options struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Value 默认值
func (i *FormItemOptions) Value(v string) *FormItemOptions {
	i.FormItem.SetOptions("value", v)
	return i
}

// AddOptions 设置静态选项
func (i *FormItemOptions) AddOptions(opts []Options) *FormItemOptions {
	i.FormItem.SetOptions("options", opts)
	return i
}

func (i *FormItemOptions) SetModel(model Model) *FormItemOptions {
	list := make([]map[string]interface{}, 0)
	ret := model.GetDB().Find(&list)
	if ret.Error != nil {
		logrus.Errorf("FromModel err=%v", ret.Error)
		return i
	}
	opts := make([]Options, 0)
	for _, m := range list {
		label := fmt.Sprintf("%s", m["value"])
		if _, ok := m["value"]; !ok {
			logrus.Error("必须设置select test as value 等, InputOptions组件需要有value、label字段")
			return i
		}
		if _, ok := m["label"]; ok {
			label = fmt.Sprintf("%s", m["label"])
		}
		opts = append(opts, Options{
			Label: label,
			Value: fmt.Sprintf("%s", m["value"]),
		})
	}

	return i.AddOptions(opts)
}

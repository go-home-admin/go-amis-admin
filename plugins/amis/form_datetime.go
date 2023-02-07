package amis

import (
	"github.com/go-home-admin/home/bootstrap/services/database"
	"strconv"
)

func (f *Form) InputDatetime(name, label string) *FormInputDatetime {
	item := &FormInputDatetime{
		FormItem: NewItem(name, label, "input-datetime"),
	}
	f.AddBody(item)
	item.SetSave(func(old interface{}) interface{} {
		switch old.(type) {
		case string:
			i, _ := strconv.Atoi(old.(string))
			if i != 0 {
				return database.UnixToTime(int64(i)).YmdHis()
			}
		}
		return "2023-01-01 08:00:00"
	})

	return item
}

type FormInputDatetime struct {
	*FormItem
}

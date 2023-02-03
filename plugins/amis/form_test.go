package amis

import (
	"testing"
)

func TestNewForm(t *testing.T) {
	form := NewForm()
	form.Input("name", "姓名")
	form.Input("email", "邮箱")
}

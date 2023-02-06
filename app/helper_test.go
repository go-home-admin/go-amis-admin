package app

import (
	"fmt"
	"github.com/go-home-admin/go-admin/app/entity/admin"
	"testing"
)

func TestGetStructFieldsInfo(t *testing.T) {
	gotFields := GetStructFieldsInfo(admin.AdminMenu{})

	for _, field := range gotFields {
		fmt.Printf("curd.Column(\"%v\", \"%v\")\n", field["comment"], field["name"])
	}
	fmt.Println("")
	for _, field := range gotFields {
		fmt.Printf("form.Input(\"%v\", \"%v\")\n", field["name"], field["comment"])
	}
}

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	personMap := map[string]interface{}{
		"Name":    "Alice",
		"Age":     "30",
		"Address": "123 Main St",
	}

	var p Person

	for k, v := range personMap {
		field := reflect.ValueOf(&p).Elem().FieldByName(k)
		if field.IsValid() {
			if field.CanSet() {
				value := reflect.ValueOf(v)
				if value.Type().AssignableTo(field.Type()) {
					field.Set(value)
				}
			}
		}
	}

	fmt.Printf("%+v\n", p)
}

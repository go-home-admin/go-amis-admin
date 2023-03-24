package amis

import "encoding/json"

func (f *Form) JsonSchema(name, label string) *FormItemJsonSchema {
	item := &FormItemJsonSchema{
		FormItem: NewItem(name, label, "json-schema"),
	}
	f.AddBody(item)
	return item
}

type FormItemJsonSchema struct {
	*FormItem
}

func (f *FormItemJsonSchema) GetValue(row map[string]interface{}) interface{} {
	v, ok := row[f.Name]
	if !ok {
		return nil
	}

	switch v.(type) {
	case string:
		vv := map[string]interface{}{}
		json.Unmarshal([]byte(v.(string)), &vv)
		return vv
	default:
		return v
	}
}

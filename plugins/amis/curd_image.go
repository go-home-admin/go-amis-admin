package amis

func (c *ColumnConfig) Image() *TableImage {
	image := &TableImage{c}
	image.Type = "image"
	return image
}

type TableImage struct {
	*ColumnConfig
}

func (t *TableImage) Width(v string) *TableImage {
	t.SetOptions("width", v)
	return t
}

func (t *TableImage) Height(v string) *TableImage {
	t.SetOptions("height", v)
	return t
}

// EnlargeAble 放大功能
func (t *TableImage) EnlargeAble() *TableImage {
	t.SetOptions("enlargeAble", true)
	return t
}

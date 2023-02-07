package amis

type ColumnConfigJson struct {
	c *ColumnConfig
}

func (c *ColumnConfig) Json() *ColumnConfigJson {
	c.Type = "json"
	got := &ColumnConfigJson{
		c: c,
	}
	return got.LevelExpand(0)
}

// JsonTheme 可配置jsonTheme，指定显示主题，可选twilight和eighties，默认为twilight。
func (c *ColumnConfigJson) JsonTheme(v int) *ColumnConfigJson {
	c.c.opt["jsonTheme"] = v
	return c
}

// LevelExpand 配置默认展开层级
func (c *ColumnConfigJson) LevelExpand(v int) *ColumnConfigJson {
	c.c.opt["levelExpand"] = v
	return c
}

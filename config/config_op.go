package config

func (c *Config) GetObj(key string) interface{} {
	unit := c.UnitsGet(key)
	if unit == nil {
		return nil
	}
	return unit.obj
}

func (c *Config) SetObj(key string, obj interface{}) {
	c.UnitsOverwriteAdd(&Unit{
		key: key,
		obj: obj,
	}, c.UnitsEqual)
}

package config

func (c *Config) GetObj(key string) interface{} {
	unit := c.UnitsGet(key)
	if unit == nil {
		return nil
	}
	return unit.obj
}

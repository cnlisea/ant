package config

import (
	"container/list"
)

func (c *Config) UnitsGet(key string) *Unit {
	var ret *Unit
	c.UnitsRange(func(u *Unit) bool {
		if u.key == key {
			ret = u
		}
		return ret == nil
	})
	return ret
}

func (c *Config) UnitsRange(fn func(u *Unit) bool) {
	if c.units == nil || c.units.Len() == 0 {
		return
	}

	var (
		e  *list.Element
		u  *Unit
		ok bool
	)
	for e = c.units.Front(); e != nil; e = e.Next() {
		if u, ok = e.Value.(*Unit); ok && !fn(u) {
			break
		}
	}
}

func (c *Config) UnitsEqual(a, b *Unit) bool {
	if a == nil || b == nil {
		return false
	}

	return a.key == b.key
}

func (c *Config) UnitsAdd(us ...*Unit) {
	if c.units == nil {
		c.units = list.New()
	}

	for i := range us {
		c.units.PushBack(us[i])
	}
}

func (c *Config) UnitsExist(unit *Unit, equal func(a, b *Unit) bool) bool {
	var exist bool
	c.UnitsRange(func(u *Unit) bool {
		if equal(unit, u) {
			exist = true
		}
		return !exist
	})
	return exist
}

package config

import (
	"github.com/cnlisea/ant/typex"
	"github.com/cnlisea/crypto"
)

func (u *Unit) Update(data []byte) error {
	sign := crypto.EncryptMD5(data)
	if u.sign != nil {
		if typex.BytesToString(sign) == typex.BytesToString(u.sign) {
			return nil
		}
	}

	if err := u.UpdateObj(data); err != nil {
		return err
	}

	if u.sign != nil && u.updateHook != nil {
		u.updateHook(u.obj)
	}
	u.sign = sign
	return nil
}

func (u *Unit) UpdateObj(data []byte) error {
	return u.Layout.Parse(data, u.obj)
}

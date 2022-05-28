package typex

import "container/list"

func ListToInts(l *list.List) []int {
	if l == nil || l.Len() == 0 {
		return nil
	}

	var (
		r   = make([]int, 0, l.Len())
		val int
		ok  bool
		e   *list.Element
	)
	for e = l.Front(); e != nil; e = e.Next() {
		if val, ok = e.Value.(int); ok {
			r = append(r, val)
		}
	}

	return r
}

func ListToBytes(l *list.List) []byte {
	if l == nil || l.Len() == 0 {
		return nil
	}

	var (
		r   = make([]byte, 0, l.Len())
		val byte
		ok  bool
		e   *list.Element
	)
	for e = l.Front(); e != nil; e = e.Next() {
		if val, ok = e.Value.(byte); ok {
			r = append(r, val)
		}
	}

	return r
}

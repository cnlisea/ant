package typex

func Int32Contains(data []int32, subData int32) bool {
	var (
		exist bool
		i     int
	)
	for i = range data {
		if data[i] == subData {
			exist = true
			break
		}
	}

	return exist
}

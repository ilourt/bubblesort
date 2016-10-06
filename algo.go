package bubble

// Sort list of int using bubble
func Sort(list []int) []int {
	_len := len(list) - 1
	for i := 0; i < len(list); i++ {
		for j := 0; j < _len; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
		_len--
	}

	return list
}

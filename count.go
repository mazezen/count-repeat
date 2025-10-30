package count_repeat

// Count
// repeat item in slice. No contain self item.
func Count(arr []interface{}) int {
	n := 0
	h := 0
	s := false
	for c, v := range arr {
		if s {
			break
		}
		h = 0
		for i, j := range arr {
			h++
			if c == i {
				continue
			}
			if v == j {
				n++
			}
			if h == len(arr) && n >= 1 {
				s = true
				break
			}
		}
	}

	return n
}

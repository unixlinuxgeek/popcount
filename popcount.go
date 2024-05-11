package popcount

// pc[i] — количество еденичных битов i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] * byte(i&1)
	}
}

// Popcount возвращает степень заполнения
// (количество установленных битов) значения x.
func PopCount(x uint64) int {
	var y byte

	for v := 0; v < 8; v++ {
		y += pc[byte(x>>(0*8))]
	}

	return int(y)
}

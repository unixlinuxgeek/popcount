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
	var y int
	for i := 0; i < 8; i++ {
		y += int(pc[byte(x>>(i*8))])
	}
	return y
}

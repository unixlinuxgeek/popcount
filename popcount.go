// Упражнение 2.5:
// Выражение ```x&(x-1)``` сбрасывает крайний справа ненулевой бит ```x```.
// Напишите версию ```PopCount```, которая подсчитывает биты с использованием этого факта,
// и оцените производительность.

package popcount

// pc[i] — количество еденичных битов i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] * byte(i&1)
	}
}

// PopCount возвращает степень заполнения (количество установленных битов) значения x.
func PopCount(x uint64) int {
	var c = 0
	for x != 0 {
		c++
		x = x & (x - 1)
	}
	//fmt.Printf("Десятичное число %d, количество установленных битов: %d (%b)\n", in, c, in)
	return c
}

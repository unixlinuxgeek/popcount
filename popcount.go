// Package popcount Упражнение 2.4
// Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента по всем 64 позициям,
// проверяя при каждом сдвиге крайний справа бит. Сравните производительность этой версии с выборкой из таблицы.

package popcount

// pc[i] — количество еденичных битов i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount возвращает степень заполнения
// (количество установленных битов) значения x.
func PopCount(x uint64) int {
	var v int
	v = int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	return v
}

func PopCountVer2(x uint64) int {
	var v int
	for i := 0; i < 64; i++ {
		f := int(pc[byte(x>>(i*8))])
		v += f
	}
	//fmt.Printf("Количество установленных битов: %d\n", v)

	return v
}

func PopCountVer3(x uint64) int {
	var c = 0

	for x != 0 {
		if (x & 1) == 1 {
			c++
		}
		x = x >> 1
	}
	//fmt.Printf("Количество установленных битов: %d\n", c)
	return c
}

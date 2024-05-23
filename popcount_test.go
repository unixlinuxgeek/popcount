// Упражнение 2.4
// Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента по всем 64 позициям,
// проверяя при каждом сдвиге крайний справа бит. Сравните производительность этой версии с выборкой из таблицы.

package popcount

import (
	"testing"
)

func TestPopCountNew(t *testing.T) {
	n := uint64(2050) // 100000000010 bit count is 2

	if PopCountNew(n) == 2 { // bit count
		t.Logf("%s Десятичное число %d, количество установленных битов: %d (%b)\n", t.Name(), n, PopCountNew(n), n)
	} else {
		t.Errorf("Error %b not equal %d", n, n)
	}
}

func TestPopCountOld(t *testing.T) {
	n := uint64(2050) // 100000000010 bit count is 2

	if PopCountOld(n) == 2 { // bit count
		t.Logf("%s Десятичное число %d, количество установленных битов: %d (%b)\n", t.Name(), n, PopCountNew(n), n)
	} else {
		t.Errorf("Error %b not equal %d", n, n)
	}
}

func BenchmarkPopCountOld(t *testing.B) {
	PopCountOld(25)
	//for i := 0; i < t.N; i++ {
	//	PopCountOld(25)
	//}
}

func BenchmarkPopCountNew(t *testing.B) {
	PopCountNew(25)
	//for i := 0; i < t.N; i++ {
	//	PopCountNew(25)
	//}
}

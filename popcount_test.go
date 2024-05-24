package popcount

import (
	"testing"
)

// Десятичное число 2050, количество установленных битов: 2 (100000000010)
func TestPopCountNew1(t *testing.T) {
	d := uint64(2050)
	bitCnt := 2
	cnt := PopCount(d)
	if bitCnt == cnt {
		t.Logf("%s: Десятичное число %d, количество установленных битов: %d (%b)\n", t.Name(), d, cnt, d)
	} else {
		t.Errorf("Ошибка число %d не равно  %d", cnt, bitCnt)
	}
}

// Десятичное число 3, количество установленных битов: 2 (11)
func TestPopCountNew2(t *testing.T) {
	d := uint64(3)
	bitCnt := 2
	cnt := PopCount(d)
	if bitCnt == cnt {
		t.Logf("%s: Десятичное число %d, количество установленных битов: %d (%b)\n", t.Name(), d, cnt, d)
	} else {
		t.Errorf("Ошибка число %d не равно  %d", cnt, bitCnt)
	}
}

// Десятичное число 55, количество установленных битов: 5 (110111)
func TestPopCountNew3(t *testing.T) {
	d := uint64(55)
	bitCnt := 5
	cnt := PopCount(d)
	if bitCnt == cnt {
		t.Logf("%s: Десятичное число %d, количество установленных битов: %d (%b)\n", t.Name(), d, cnt, d)
	} else {
		t.Errorf("Ошибка число %d не равно  %d", cnt, bitCnt)
	}
}

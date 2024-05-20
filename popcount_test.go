package popcount

import (
	"testing"
)

func TestPopCount1(t *testing.T) {
	num := uint64(2500)
	x := PopCount(num)
	t.Logf("%s число %d, двоичное представление: %b, клличество установленных битов: %d\n", t.Name(), num, num, x)
}

func TestPopCount2(t *testing.T) {
	num := uint64(2500)
	x := PopCountVer2(num)
	t.Logf("%s число %d, двоичное представление: %b, клличество установленных битов: %d\n", t.Name(), num, num, x)
}

func TestPopCount3(t *testing.T) {
	num := uint64(2500)
	x := PopCountVer3(num)
	t.Logf("%s число %d, двоичное представление: %b, клличество установленных битов: %d\n", t.Name(), num, num, x)
}

// 0.0000003 ns/op
func BenchmarkPopCount(b *testing.B) {
	s := uint64(2500)
	PopCount(s)
}

// 0.0000002 ns/op
func BenchmarkPopCount2(b *testing.B) {
	s := uint64(2500)
	PopCountVer2(s)
}

// 0.0000003 ns/op
func BenchmarkPopCount3(b *testing.B) {
	s := uint64(2500)
	PopCountVer3(s)
}

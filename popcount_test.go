package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	d := 2050
	cnt := PopCountNew(2050)
	t.Logf("Десятичное число %d, количество установленных битов: %d (%b)\n", d, cnt, d)
}

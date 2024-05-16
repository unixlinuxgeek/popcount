package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	d := uint64(10) // output: 1010
	// 1010 = 2 установленных бита
	if PopCount(d) == 2 {
		t.Logf("Test Passed: TestPopCount(%b) == 1010", d)
	} else {
		t.Errorf("Test Error: TestPopCount(%b) not equal 1010", d)
	}
}
func TestPopCountNew(t *testing.T) {
	d := uint64(10) // output: 1010
	// 1010 = 2 установленных бита
	if PopCountNew(d) == PopCount(d) {
		t.Logf("Test Passed: PopCountNew(%b) equal PopCount(%b)", d, d)
	} else {
		t.Errorf("Test Error: PopCountNew(%b) not equal PopCount(%b)", d, d)
	}
}

package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	x := uint64(18446744073709551615)

	v1 := PopCount(x)
	v2 := PopCount1(x)

	fmt.Printf("%d = %d\n", x, v1)
	fmt.Printf("%d = %d\n", x, v2)
	if v1 == v2 {
		t.Logf("Test Passed: %d equal %d", v1, v2)
	} else {
		t.Errorf("Test Not Passed!!!: %d not equal %d", v1, v2)
	}
}

package unittest

import (
	"testing"
)

func TestCalcAdd(t *testing.T) {
	result := CalcAdd(10)
	if result != 45 {
		t.Fatalf("结果错误！")
	}

	t.Logf("结果正确！")
}

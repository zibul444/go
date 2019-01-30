package math

import "testing"

type TestPair struct {
	value   []float64
	average float64
}

var testpair = []TestPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range testpair {
		v := Average(pair.value)
		if v != pair.average {
			t.Error(
				"For", pair.value,
				"expected", pair.average,
				"got", v,
			)
		}
	}

	//v := Average([]float64{1,2})
	//if v != 1.5 {
	//	t.Error("Expected 1.5, got ", v)
	//}
}

func TestMax(t *testing.T) {
	v := Max([]float64{1, 2})
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{1, 2})
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

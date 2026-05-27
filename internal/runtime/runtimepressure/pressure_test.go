package runtimepressure

import "testing"

func TestPressure(
	t *testing.T,
) {
	var p Pressure

	p.Inc()

	if p.Load() != 1 {
		t.Fatal("pressure increment failed")
	}

	p.Dec()

	if p.Load() != 0 {
		t.Fatal("pressure decrement failed")
	}
}

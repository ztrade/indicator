package indicator

import (
	"math"
	"strings"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestStochUsesFirstSampleInWindow(t *testing.T) {
	s := NewStoch(14, 3, 3)

	if err := s.Update(100); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := s.Update(90); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !floatEquals(s.Result(), 0) {
		t.Fatalf("unexpected stoch result, got %.6f, want 0", s.Result())
	}
}

func TestRegisterIndicatorCaseInsensitive(t *testing.T) {
	name := "unit_test_custom_indicator"
	upper := strings.ToUpper(name)
	delete(ExtraIndicators, upper)
	defer delete(ExtraIndicators, upper)

	RegisterIndicator(name, func(params ...int) (CommonIndicator, error) {
		return NewMixed(NewSMA(2), nil), nil
	})

	_, err := NewCommonIndicator(name, 2)
	if err != nil {
		t.Fatalf("unexpected register indicator error: %v", err)
	}
}

func TestNewCommonIndicatorRejectsNonPositiveParams(t *testing.T) {
	tests := []struct {
		name   string
		params []int
	}{
		{name: "SMA", params: []int{0}},
		{name: "EMA", params: []int{-1}},
		{name: "MACD", params: []int{12, 26, 0}},
		{name: "STOCHRSI", params: []int{14, 14, 0, 3}},
		{name: "BOLL", params: []int{20, 0}},
	}

	for _, tt := range tests {
		_, err := NewCommonIndicator(tt.name, tt.params...)
		if err == nil {
			t.Fatalf("expected error for %s with params %v", tt.name, tt.params)
		}
	}
}

func TestConstructorsNormalizeNonPositivePeriods(t *testing.T) {
	if NewSMA(0).winLen != 1 {
		t.Fatalf("SMA period should fallback to 1")
	}
	if NewEMA(0).winLen != 1 {
		t.Fatalf("EMA period should fallback to 1")
	}
	if NewSMMA(0).winLen != 1 {
		t.Fatalf("SMMA period should fallback to 1")
	}
	if NewRSI(0).winLen != 1 {
		t.Fatalf("RSI period should fallback to 1")
	}

	st := NewStoch(0, 0, 0)
	if st.winLen != 1 {
		t.Fatalf("Stoch window should fallback to 1")
	}

	b := NewBoll(0, 0)
	if b.winLen != 1 || b.k != 1 {
		t.Fatalf("Boll settings should fallback to 1")
	}

	macd := NewMACD(0, 0, 0)
	if macd == nil {
		t.Fatalf("NewMACD should return a valid instance")
	}
}

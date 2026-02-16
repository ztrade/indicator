package indicator

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestATRUpdateOHLC(t *testing.T) {
	atr := NewATR(3)
	bars := []struct {
		h float64
		l float64
		c float64
	}{
		{h: 10, l: 9, c: 9.5},
		{h: 11, l: 9.5, c: 10.5},
		{h: 12, l: 10, c: 11},
		{h: 11.5, l: 10.5, c: 11},
	}

	atr.UpdateOHLC(bars[0].h, bars[0].l, bars[0].c)
	if !almostEqual(atr.Result(), 0) {
		t.Fatalf("unexpected initial atr value, got %.12f", atr.Result())
	}

	atr.UpdateOHLC(bars[1].h, bars[1].l, bars[1].c)
	if !almostEqual(atr.Result(), 0) {
		t.Fatalf("unexpected atr value before warm-up, got %.12f", atr.Result())
	}

	atr.UpdateOHLC(bars[2].h, bars[2].l, bars[2].c)
	if !almostEqual(atr.Result(), 1.5) {
		t.Fatalf("unexpected atr after warm-up, got %.12f, want 1.5", atr.Result())
	}

	atr.UpdateOHLC(bars[3].h, bars[3].l, bars[3].c)
	if !almostEqual(atr.Result(), 1.3333333333333333) {
		t.Fatalf("unexpected atr smoothed value, got %.12f", atr.Result())
	}
	if !almostEqual(atr.TR(), 1) {
		t.Fatalf("unexpected last true range, got %.12f", atr.TR())
	}
}

func TestADXUpdateOHLC(t *testing.T) {
	adx := NewADX(3)
	bars := []struct {
		h float64
		l float64
		c float64
	}{
		{h: 30, l: 28, c: 29},
		{h: 32, l: 29, c: 31},
		{h: 33, l: 30, c: 32},
		{h: 31, l: 29, c: 30},
		{h: 30, l: 27, c: 28},
		{h: 29, l: 26, c: 27},
		{h: 31, l: 27, c: 30},
		{h: 32, l: 28, c: 31},
	}

	for idx, bar := range bars {
		adx.UpdateOHLC(bar.h, bar.l, bar.c)
		if idx == 4 && !almostEqual(adx.Result(), 0) {
			t.Fatalf("adx should still be warming up, got %.12f", adx.Result())
		}
	}

	if !almostEqual(adx.PlusDI(), 27.430555555555554) {
		t.Fatalf("unexpected +DI, got %.12f", adx.PlusDI())
	}
	if !almostEqual(adx.MinusDI(), 11.574074074074073) {
		t.Fatalf("unexpected -DI, got %.12f", adx.MinusDI())
	}
	if !almostEqual(adx.DX(), 40.652818991097924) {
		t.Fatalf("unexpected DX, got %.12f", adx.DX())
	}
	if !almostEqual(adx.Result(), 33.14106550382515) {
		t.Fatalf("unexpected ADX, got %.12f", adx.Result())
	}
}

func TestNewCommonIndicatorSupportsATRAndADX(t *testing.T) {
	if _, err := NewCommonIndicator("ATR", 14); err != nil {
		t.Fatalf("expected ATR to be supported, got error: %v", err)
	}
	if _, err := NewCommonIndicator("ADX", 14); err != nil {
		t.Fatalf("expected ADX to be supported, got error: %v", err)
	}
}

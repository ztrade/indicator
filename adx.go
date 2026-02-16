package indicator

import "math"

// ADX Average Directional Index.
// Use UpdateOHLC for standard ADX, or Update for close-only fallback.
type ADX struct {
	winLen int

	trMA      *wilderAverage // Wilder smoothing for TR
	plusDMMA  *wilderAverage // Wilder smoothing for +DM
	minusDMMA *wilderAverage // Wilder smoothing for -DM
	adxMA     *wilderAverage // Wilder smoothing for DX

	lastHigh  float64
	lastLow   float64
	lastClose float64
	hasLast   bool

	plusDI  float64
	minusDI float64
	dx      float64
	result  float64
}

func NewADX(winLen int) *ADX {
	winLen = normalizePeriod(winLen)
	return &ADX{
		winLen:    winLen,
		trMA:      newWilderAverage(winLen),
		plusDMMA:  newWilderAverage(winLen),
		minusDMMA: newWilderAverage(winLen),
		adxMA:     newWilderAverage(winLen),
	}
}

func (a *ADX) Update(values ...float64) {
	o, h, l, c := getOHLC(values)
	a.UpdateOHLC(o, h, l, c)
}

func (a *ADX) UpdateOHLC(open, high, low, close float64) {
	if !a.hasLast {
		a.lastHigh = high
		a.lastLow = low
		a.lastClose = close
		a.hasLast = true
		return
	}

	tr := math.Abs(high - low)
	tr = math.Max(tr, math.Abs(high-a.lastClose))
	tr = math.Max(tr, math.Abs(low-a.lastClose))

	upMove := high - a.lastHigh
	downMove := a.lastLow - low

	plusDM := 0.0
	if upMove > downMove && upMove > 0 {
		plusDM = upMove
	}
	minusDM := 0.0
	if downMove > upMove && downMove > 0 {
		minusDM = downMove
	}

	smoothedTR := a.trMA.Update(tr)
	smoothedPlusDM := a.plusDMMA.Update(plusDM)
	smoothedMinusDM := a.minusDMMA.Update(minusDM)

	a.plusDI = 0
	a.minusDI = 0
	a.dx = 0
	if a.trMA.Ready() && smoothedTR != 0 {
		a.plusDI = 100 * smoothedPlusDM / smoothedTR
		a.minusDI = 100 * smoothedMinusDM / smoothedTR
		diSum := a.plusDI + a.minusDI
		if diSum != 0 {
			a.dx = 100 * math.Abs(a.plusDI-a.minusDI) / diSum
		}
		a.result = a.adxMA.Update(a.dx)
	}

	a.lastHigh = high
	a.lastLow = low
	a.lastClose = close
	a.hasLast = true
}

func (a *ADX) Result() float64 {
	return a.result
}

func (a *ADX) PlusDI() float64 {
	return a.plusDI
}

func (a *ADX) MinusDI() float64 {
	return a.minusDI
}

func (a *ADX) DX() float64 {
	return a.dx
}

func (a *ADX) FastResult() float64 {
	return a.plusDI
}

func (a *ADX) SlowResult() float64 {
	return a.minusDI
}

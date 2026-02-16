package indicator

import "math"

// ATR Average True Range
// Use UpdateOHLC for standard ATR, or Update for close-only fallback.
type ATR struct {
	winLen       int
	trMA         *wilderAverage // Wilder smoothing for True Range
	lastClose    float64
	hasLastClose bool
	tr           float64
	result       float64
}

func NewATR(winLen int) *ATR {
	winLen = normalizePeriod(winLen)
	return &ATR{winLen: winLen, trMA: newWilderAverage(winLen)}
}

func (a *ATR) Update(values ...float64) {
	o, h, l, c := getOHLC(values)
	a.UpdateOHLC(o, h, l, c)
}

func (a *ATR) UpdateOHLC(open, high, low, close float64) {
	tr := math.Abs(high - low)
	if a.hasLastClose {
		tr = math.Max(tr, math.Abs(high-a.lastClose))
		tr = math.Max(tr, math.Abs(low-a.lastClose))
	}
	a.tr = tr
	a.result = a.trMA.Update(tr)
	a.lastClose = close
	a.hasLastClose = true
}

func (a *ATR) Result() float64 {
	return a.result
}

func (a *ATR) TR() float64 {
	return a.tr
}

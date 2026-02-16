package indicator

// wilderAverage implements Welles Wilder's smoothing method.
//
// It has two phases:
//  1. Warm-up: collect `period` samples and initialize with their SMA.
//  2. Recursive smoothing:
//     next = (prev*(period-1) + current) / period
//
// ATR uses one wilderAverage over TR;
// ADX uses wilderAverage over TR/+DM/-DM, and then another one over DX.
type wilderAverage struct {
	period int
	count  int
	sum    float64
	value  float64
}

func newWilderAverage(period int) *wilderAverage {
	return &wilderAverage{period: normalizePeriod(period)}
}

// Update pushes one sample and returns the latest smoothed value.
// During warm-up (count < period), the return value remains 0 until enough
// samples are collected to build the initial SMA.
func (w *wilderAverage) Update(value float64) float64 {
	if w.count < w.period {
		w.count++
		w.sum += value
		if w.count == w.period {
			w.value = w.sum / float64(w.period)
		}
		return w.value
	}
	w.value = (w.value*float64(w.period-1) + value) / float64(w.period)
	return w.value
}

// Ready reports whether warm-up is complete.
func (w *wilderAverage) Ready() bool {
	return w.count >= w.period
}

package indicator

// Stoch just test with StochRSI
type Stoch struct {
	winLen  int
	prices  []float64
	lowest  float64
	highest float64
	age     int
	bFirst  bool
	result  float64
	kSMA    *SMA
	dSMA    *SMA
}

func NewStoch(winLen, periodK, periodD int) *Stoch {
	s := new(Stoch)
	s.winLen = winLen
	s.prices = make([]float64, s.winLen)
	s.bFirst = true
	s.kSMA = NewSMA(periodK)
	s.dSMA = NewSMA(periodD)
	return s
}

func (s *Stoch) Update(price float64) {
	if s.bFirst {
		s.lowest = price
		s.highest = price
		s.bFirst = false
		return
	}
	s.prices[s.age] = price
	s.age = (s.age + 1) % s.winLen
	s.highest = highest(s.prices)
	s.lowest = lowest(s.prices)
	if s.highest == s.lowest {
		return
	}
	s.result = (100 * (price - s.lowest)) / (s.highest - s.lowest)
	s.kSMA.Update(s.result)
	s.dSMA.Update(s.kSMA.Result())
}

func (s *Stoch) Result() float64 {
	return s.result
}

func (s *Stoch) KResult() float64 {
	return s.kSMA.Result()
}

func (s *Stoch) DResult() float64 {
	return s.dSMA.Result()
}
func lowest(prices []float64) (ret float64) {
	ret = prices[0]
	for _, v := range prices {
		if v < ret {
			ret = v
		}
	}
	return
}

func highest(prices []float64) (ret float64) {
	ret = prices[0]
	for _, v := range prices {
		if v > ret {
			ret = v
		}
	}
	return
}

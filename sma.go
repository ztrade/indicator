package indicator

type SMA struct {
	MABase
	prices []float64
	sum    float64
	age    int
}

func NewSMA(winLen int) *SMA {
	s := new(SMA)
	s.winLen = winLen
	s.prices = make([]float64, s.winLen)
	s.age = 0
	return s
}

func (s *SMA) Update(price float64) {
	s.result = price
	tail := s.prices[s.age]
	s.prices[s.age] = price
	s.sum += price - tail
	s.result = s.sum / float64(s.winLen)
	s.age = (s.age + 1) % s.winLen
}

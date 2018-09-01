package indicator

type EMA struct {
	MABase
	alpha  float64
	bFirst bool
}

func NewEMA(winLen int) *EMA {
	e := new(EMA)
	e.winLen = winLen
	e.alpha = 2 / float64((e.winLen + 1))
	e.bFirst = true
	return e
}

func (e *EMA) Update(price float64) {
	if e.bFirst {
		e.result = price
		e.bFirst = false
	}
	e.cal(price)
}

// cal
// EMA = alpha * x + (1 - alpha) * EMA[1]
// alpha = 2 / (y + 1)
func (e *EMA) cal(price float64) {
	oldResult := e.result
	e.result = e.alpha*price + (1-e.alpha)*oldResult
}

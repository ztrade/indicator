package indicator

type RSI struct {
	winLen    int
	avgU      *SMMA
	avgD      *SMMA
	u         float64
	d         float64
	lastClose *float64
	rs        float64
	result    float64
}

func NewRSI(winLen int) *RSI {
	r := new(RSI)
	r.winLen = winLen
	r.avgU = NewSMMA(r.winLen)
	r.avgD = NewSMMA(r.winLen)
	return r
}

func (r *RSI) Update(price float64) {
	if r.lastClose == nil {
		r.lastClose = &price
		return
	}
	if price > *r.lastClose {
		r.u = price - *r.lastClose
		r.d = 0
	} else {
		r.u = 0
		r.d = *r.lastClose - price
	}
	r.avgU.Update(r.u)
	r.avgD.Update(r.d)
	uResult := r.avgU.Result()
	dResult := r.avgD.Result()
	r.rs = uResult / dResult
	r.result = 100 - (100 / (1 + r.rs))
	if dResult == 0 {
		if uResult != 0 {
			r.result = 100
		} else {
			r.result = 0
		}
	}
	r.lastClose = &price
}

func (r *RSI) Result() float64 {
	return r.result
}

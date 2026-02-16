package indicator

// StochRSI
// result test with aicoin's bitmex data
// maybe it's difference with different website
type StochRSI struct {
	winLen int
	r      *RSI
	st     *Stoch
}

func NewStochRSI(winLen, rsiWinLen, k, d int) *StochRSI {
	sr := new(StochRSI)
	sr.winLen = normalizePeriod(winLen)
	sr.r = NewRSI(normalizePeriod(rsiWinLen))
	sr.st = NewStoch(sr.winLen, normalizePeriod(k), normalizePeriod(d))
	return sr
}

func (sr *StochRSI) Update(price float64) {
	sr.r.Update(price)
	sr.st.Update(sr.r.Result())
}

func (sr *StochRSI) KResult() float64 {
	return sr.st.KResult()
}

func (sr *StochRSI) DResult() float64 {
	return sr.st.DResult()
}

func (sr *StochRSI) Result() float64 {
	return sr.st.Result()
}

func (sr *StochRSI) FastResult() float64 {
	return sr.st.KResult()
}

func (sr *StochRSI) SlowResult() float64 {
	return sr.st.DResult()
}

package indicator

type MACD struct {
	long   *EMA
	short  *EMA
	signal *EMA
	dif    float64
	dea    float64
	result float64
}

func NewMACD(short, long, signal int) *MACD {
	ma := new(MACD)
	ma.short = NewEMA(short)
	ma.long = NewEMA(long)
	ma.signal = NewEMA(signal)
	return ma
}

func (ma *MACD) Update(price float64) {
	ma.long.Update(price)
	ma.short.Update(price)

	ma.dif = ma.short.Result() - ma.long.Result()
	ma.signal.Update(ma.dif)
	ma.dea = ma.signal.Result()
	ma.result = 2 * (ma.dif - ma.dea)
}

func (ma *MACD) Result() float64 {
	return ma.result
}

func (ma *MACD) DIF() float64 {
	return ma.dif
}

func (ma *MACD) DEA() float64 {
	return ma.dea
}

func (ma *MACD) FastResult() float64 {
	return ma.dif
}

func (ma *MACD) SlowResult() float64 {
	return ma.dea
}

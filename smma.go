package indicator

// SMMA Smoothed Moving Average (SMMA)
type SMMA struct {
	MABase
	sma    *SMA
	age    int
	bFirst bool
}

func NewSMMA(winLen int) *SMMA {
	sm := new(SMMA)
	sm.winLen = winLen
	sm.sma = NewSMA(sm.winLen)
	sm.bFirst = true
	return sm
}

func (sm *SMMA) Update(price float64) {
	if sm.bFirst {
		nLen := sm.age + 1
		if nLen < sm.winLen {
			sm.sma.Update(price)
		} else if nLen == sm.winLen {
			sm.sma.Update(price)
			sm.result = sm.sma.Result()
			sm.bFirst = false
		}
		sm.age++
	} else {
		sm.result = (sm.result*float64(sm.winLen-1) + price) / float64(sm.winLen)
	}
}

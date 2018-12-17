package indicator

import (
	"math"
)

type Boll struct {
	*SMA
	k      int
	mid    float64
	top    float64
	bottom float64
}

func NewBoll(winLen, k int) *Boll {
	b := new(Boll)
	b.SMA = NewSMA(winLen)
	b.k = k
	return b
}

func (b *Boll) Result() float64 {
	return b.mid
}

func (b *Boll) Update(price float64) {
	b.SMA.Update(price)
	b.Cal()
}

func (b *Boll) Cal() {
	b.mid = b.SMA.result
	var sd float64
	for j := 0; j < b.winLen; j++ {
		sd += math.Pow(b.prices[j]-b.mid, 2)
	}
	sd = math.Sqrt(sd/float64(b.winLen)) * float64(b.k)
	b.top = b.mid + sd
	b.bottom = b.mid - sd
}

func (b *Boll) Top() float64 {
	return b.top
}

func (b *Boll) Bottom() float64 {
	return b.bottom
}

func (b *Boll) Indicator() map[string]float64 {
	return map[string]float64{"result": b.Result(), "top": b.Top(), "bottom": b.Bottom()}
}

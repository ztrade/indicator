package indicator

type Updater interface {
	Update(price float64)
}

type Indicator interface {
	Updater
	Result() float64
}

type Crosser interface {
	Updater
	SlowResult() float64
	FastResult() float64
}

type OHLCUpdater interface {
	UpdateOHLC(open, high, low, close float64)
}
type OHLCIndicator interface {
	OHLCUpdater
	Result() float64
}

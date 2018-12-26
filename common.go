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


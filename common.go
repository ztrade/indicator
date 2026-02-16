package indicator

type Updater interface {
	Update(price ...float64)
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

func getPrice(args []float64) (price float64, isOHLC bool) {
	switch len(args) {
	case 1:
		price = args[0]
		isOHLC = false
	case 4:
		price = args[3] // close price
		isOHLC = true
	default:
		panic("invalid number of arguments, expected 1 or 4")
	}
	return
}

func getOHLC(args []float64) (open, high, low, close float64) {
	if len(args) != 4 {
		panic("invalid number of arguments, expected 4 for OHLC")
	}
	open, high, low, close = args[0], args[1], args[2], args[3]
	return
}

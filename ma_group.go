package indicator

type MAGroup struct {
	fast Indicator
	slow Indicator
}

func NewMAGroup(fast, slow Indicator) *MAGroup {
	mg := new(MAGroup)
	mg.fast = fast
	mg.slow = slow
	return mg
}

func (mg *MAGroup) Update(values ...float64) error {
	price, _, err := getPrice(values)
	if err != nil {
		return err
	}
	mg.fast.Update(price)
	mg.slow.Update(price)
	return nil
}

func (mg *MAGroup) FastResult() float64 {
	return mg.fast.Result()
}

func (mg *MAGroup) SlowResult() float64 {
	return mg.slow.Result()
}

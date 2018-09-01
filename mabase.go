package indicator

type MABase struct {
	winLen int //window length
	result float64
}

func (m *MABase) Result() float64 {
	return m.result
}

package indicator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// MixIndicator
type MixIndicator interface {
	Updater
	SlowResult() float64
	FastResult() float64
	Result() float64
	SupportSlowFast() bool
	SupportResult() bool
}

func NewMixIndicator(name string, params ...int) (ind MixIndicator, err error) {
	name = strings.ToUpper(name)
	nLen := len(params)
	if nLen == 0 {
		err = fmt.Errorf("%s params can't be empty", name)
		return
	}
	switch name {
	case "EMA":
		if nLen >= 2 {
			maGroup := NewMAGroup(NewEMA(params[0]), NewEMA(params[1]))
			ind = NewMixed(nil, maGroup)
		} else {
			ema := NewEMA(params[0])
			ind = NewMixed(ema, nil)
		}
	case "MACD":
		if nLen < 3 {
			err = fmt.Errorf("%s params not enough", name)
		} else {
			macd := NewMACD(params[0], params[1], params[2])
			ind = NewMixed(macd, macd)
		}
	case "SMA":
		if nLen >= 2 {
			maGroup := NewMAGroup(NewSMA(params[0]), NewSMA(params[1]))
			ind = NewMixed(nil, maGroup)
		} else {
			sma := NewSMA(params[0])
			ind = NewMixed(sma, nil)
		}
	case "SMMA":
		if nLen >= 2 {
			maGroup := NewMAGroup(NewSMMA(params[0]), NewSMMA(params[1]))
			ind = NewMixed(nil, maGroup)
		} else {
			smma := NewSMMA(params[0])
			ind = NewMixed(smma, nil)
		}
	case "STOCHRSI":
		if nLen < 4 {
			err = fmt.Errorf("%s params not enough", name)
		} else {
			stochRSI := NewStochRSI(params[0], params[1], params[2], params[3])
			ind = NewMixed(stochRSI, stochRSI)
		}
	default:
		err = fmt.Errorf("%s indicator not support", name)
	}
	return
}

type Mixed struct {
	indicator      Indicator
	crossIndicator Crosser
	isSameOne      bool
	crossTool      *CrossTool
}

func NewMixed(indicator Indicator, crossIndicator Crosser) *Mixed {
	m := new(Mixed)
	m.indicator = indicator
	m.crossIndicator = crossIndicator
	if m.crossIndicator != nil {
		m.crossTool = NewCrossTool(m.crossIndicator)
	}
	m.checkSameOne()
	return m
}

func (m *Mixed) checkSameOne() {
	if reflect.ValueOf(m.indicator) == reflect.ValueOf(m.crossIndicator) {
		m.isSameOne = true
	} else {
		m.isSameOne = false
	}
}

func (m *Mixed) Update(price float64) {
	if m.crossTool != nil {
		m.crossTool.Update(price)
	}
	if !m.isSameOne && m.indicator != nil {
		m.indicator.Update(price)
	}
}

func (m *Mixed) FastResult() float64 {
	if m.crossIndicator != nil {
		return m.crossIndicator.FastResult()
	}
	return 0
}

func (m *Mixed) SlowResult() float64 {
	if m.crossIndicator != nil {
		return m.crossIndicator.SlowResult()
	}
	return 0
}

func (m *Mixed) IsCrossUp() bool {
	if m.crossTool == nil {
		fmt.Println("cross tool is nil")
		return false
	}
	return m.crossTool.IsCrossUp()
}
func (m *Mixed) IsCrossDown() bool {
	if m.crossTool == nil {
		return false
	}
	return m.crossTool.IsCrossDown()
}

func (m *Mixed) Result() float64 {
	if m.indicator != nil {
		return m.indicator.Result()
	}
	return 0
}

func (m *Mixed) SupportResult() bool {
	if m.indicator != nil {
		return true
	}
	return false
}

func (m *Mixed) SupportSlowFast() bool {
	if m.crossIndicator != nil {
		return true
	}
	return false
}

func (m *Mixed) MarshalJSON() (buf []byte, err error) {
	ret := make(map[string]interface{})
	if m.SupportResult() {
		ret["result"] = m.Result()
	}
	if m.SupportSlowFast() {
		ret["fast"] = m.FastResult()
		ret["slow"] = m.SlowResult()
		if m.IsCrossDown() {
			ret["crossDown"] = true
		} else {
			ret["crossDown"] = false
		}

		if m.IsCrossUp() {
			ret["crossUp"] = true
		} else {
			ret["crossUp"] = false
		}
	}
	buf, err = json.Marshal(ret)
	return
}

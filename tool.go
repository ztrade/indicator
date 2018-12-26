package indicator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type NewCommonIndicatorFunc func(params ...int) (CommonIndicator, error)

var (
	ExtraIndicators = map[string]NewCommonIndicatorFunc{}
)

// CommonIndicator
type CommonIndicator interface {
	Indicator
	Indicator() map[string]float64
}

func RegisterIndicator(name string, fn NewCommonIndicatorFunc) {
	ExtraIndicators[name] = fn
}

type JsonIndicator struct {
	CommonIndicator
}

func NewJsonIndicator(m CommonIndicator) *JsonIndicator {
	return &JsonIndicator{CommonIndicator: m}
}

func (j *JsonIndicator) MarshalJSON() (buf []byte, err error) {
	ret := j.Indicator()
	buf, err = json.Marshal(ret)
	return
}

func NewCommonIndicator(name string, params ...int) (ind CommonIndicator, err error) {
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
	case "SMAMACD":
		if nLen < 3 {
			err = fmt.Errorf("%s params not enough", name)
		} else {
			macd := NewMACDWithSMA(params[0], params[1], params[2])
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
	case "RSI":
		if nLen >= 2 {
			maGroup := NewMAGroup(NewRSI(params[0]), NewRSI(params[1]))
			ind = NewMixed(nil, maGroup)
		} else {
			rsi := NewRSI(params[0])
			ind = NewMixed(rsi, nil)
		}
	case "BOLL":
		if nLen >= 2 {
			boll := NewBoll(params[0], params[1])
			ind = boll
		} else {
			err = fmt.Errorf("%s params not enough", name)
		}
	default:
		fn, ok := ExtraIndicators[name]
		if !ok {
			err = fmt.Errorf("%s indicator not support", name)
		} else {
			ind, err = fn(params...)
		}
	}
	if err == nil {
		ind = NewJsonIndicator(ind)
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

func (m *Mixed) Indicator() map[string]float64 {
	ret := make(map[string]float64)
	if m.SupportResult() {
		ret["result"] = m.Result()
	}
	if m.SupportSlowFast() {
		ret["fast"] = m.FastResult()
		ret["slow"] = m.SlowResult()
		if m.IsCrossDown() {
			ret["crossDown"] = 1
		} else {
			ret["crossDown"] = 0
		}

		if m.IsCrossUp() {
			ret["crossUp"] = 1
		} else {
			ret["crossUp"] = 0
		}
	}
	return ret
}

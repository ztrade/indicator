package indicator

type CrossTool struct {
	crosser Crosser
	fasts   [3]float64 // prev fasts
	slows   [3]float64 // prev slows
}

func NewCrossTool(crosser Crosser) *CrossTool {
	ct := new(CrossTool)
	ct.crosser = crosser
	return ct
}

func (ct *CrossTool) Update(price float64) {
	ct.fasts[2] = ct.fasts[1]
	ct.fasts[1] = ct.fasts[0]
	ct.slows[2] = ct.slows[1]
	ct.slows[1] = ct.slows[0]

	ct.crosser.Update(price)

	ct.fasts[0] = ct.crosser.FastResult()
	ct.slows[0] = ct.crosser.SlowResult()
}

func (ct *CrossTool) IsCrossUp() bool {
	prevFast, prevSlow := ct.getPrev()
	fast, slow := ct.getCurrent()
	if prevFast < prevSlow && fast > slow {
		// fmt.Println("up:", ct.fasts, ct.slows)
		return true
	}
	return false
}

func (ct *CrossTool) IsCrossDown() bool {
	prevFast, prevSlow := ct.getPrev()
	fast, slow := ct.getCurrent()
	if prevFast > prevSlow && fast < slow {
		// fmt.Println("down:", ct.fasts, ct.slows)
		return true
	}
	return false
}

func (ct *CrossTool) getCurrent() (fast, slow float64) {
	fast, slow = ct.fasts[0], ct.slows[0]
	return
}
func (ct *CrossTool) getPrev() (prevFast, prevSlow float64) {
	prevFast = ct.fasts[1]
	prevSlow = ct.slows[1]
	// if prev fast is same of prev slow,see prev one
	if prevFast == prevSlow {
		prevFast = ct.fasts[2]
		prevSlow = ct.slows[2]
	}
	return
}

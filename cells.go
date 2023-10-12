package main

type Cell struct {
	state       bool
	nextState   bool
	liveTime    int
	aroundCells []*Cell
}

func (c *Cell) Set(state bool) {
	c.state = state
}
func (c *Cell) Link(aroundCells []*Cell) {
	c.aroundCells = aroundCells
}
func (c *Cell) Unlink() {
	c.aroundCells = []*Cell{}
}
func (c *Cell) CalcNextState() {
	var nextState bool
	count := 0
	for _, ac := range c.aroundCells {
		if ac.state {
			count++
		}
	}
	if c.state {
		if count <= 1 || count >= 4 {
			nextState = false
		} else {
			nextState = true
		}
	} else {
		if count == 3 {
			nextState = true
		} else {
			nextState = false
		}
	}
	c.nextState = nextState
}

func (c *Cell) Flush() {
	if c.state && c.nextState {
		c.liveTime++
	} else {
		c.liveTime = 0
	}
	c.state = c.nextState
}
func (c *Cell) Switch() {
	c.state = !c.state
}

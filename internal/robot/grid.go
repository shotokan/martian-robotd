package robot

// the upper-right coordinates of the rectangular world
type Grid struct {
	top    int32
	right  int32
	bottom int32
	left   int32
	Scents []Position
}

func NewGrid(right, top, left, bottom int32) *Grid {
	scents := make([]Position, 0, 10)
	return &Grid{
		top:    top,
		right:  right,
		bottom: bottom,
		left:   left,
		Scents: scents,
	}
}

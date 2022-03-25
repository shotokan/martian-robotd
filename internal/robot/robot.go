package robot

import (
	"fmt"
)

type Position struct {
	x int32
	y int32
}

type Robot struct {
	position    Position
	orientation int32
	isLost      bool
	surface     *Grid
}

func (r *Robot) String() string {
	return fmt.Sprintf("    %d %d %s %s", r.position.x, r.position.y, formatRadiansOrientation(r.orientation), formatLost(r.isLost))
}

// Method that start executing commands
func (r *Robot) On(commands []Command) {
	for _, command := range commands {
		command.Execute(r)
		if r.isLost {
			break
		}
	}
}

func (r *Robot) PrintLastPosition() {
	fmt.Println("-- last position --")
	fmt.Println(r.String())
	fmt.Println("-------------------")
}

// Method that creates a new robot with the existence surface
func NewRobot(x, y int32, orientation int32, surface *Grid) Robot {
	currentPosition := Position{x: x, y: y}
	return Robot{
		position:    currentPosition,
		orientation: orientation,
		surface:     surface,
	}
}

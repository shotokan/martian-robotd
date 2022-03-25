package robot

import (
	"fmt"
	"math"
)

type Command interface {
	Execute(robot *Robot)
}

type LeftCommand struct{}

// Turns left 90 degrees
func (lc LeftCommand) Execute(robot *Robot) {
	robot.orientation += 90
}

type RightCommand struct{}

// Turns right 90 degrees
func (rc RightCommand) Execute(robot *Robot) {
	robot.orientation -= 90
}

type ForwardCommand struct{}

// Calculate next step for the robot
// and validates that the next step is within the surface and
// in case it is outside that coordinates are stored in the scents
// to prevent other robots move to that point
func (fc ForwardCommand) Execute(robot *Robot) {
	orientation_radians := float64(robot.orientation) * math.Pi / 180.0

	newPositionX := robot.position.x + int32(math.Cos(orientation_radians))
	newPositionY := robot.position.y + int32(math.Sin(orientation_radians))

	if fc.isPositionOutsideArea(newPositionX, newPositionY, *robot.surface) {

		// Function that search if new position has been visited before
		containsPosition := func(x, y int32) bool {
			for _, visitedPos := range robot.surface.Scents {
				if fc.isInProtectedArea(x, y, visitedPos.x, visitedPos.y) {
					return true
				}
			}
			return false
		}

		// Validate if the new point is outside
		if !containsPosition(newPositionX, newPositionY) {
			newPosition := Position{newPositionX, newPositionY}
			robot.surface.Scents = append(robot.surface.Scents, newPosition)
			robot.isLost = true
		}

	} else {
		// Move the robot
		robot.position.x = newPositionX
		robot.position.y = newPositionY
	}
}

func (fc ForwardCommand) isPositionOutsideArea(newPositionX, newPositionY int32, surface Grid) bool {
	return newPositionX < surface.left || newPositionX > surface.right || newPositionY < surface.bottom || newPositionY > surface.top
}

func (fc ForwardCommand) isInProtectedArea(newPositionX, newPositionY int32, visitedX, visitedY int32) bool {
	return newPositionX == visitedX && newPositionY == visitedY
}

func NewCommand(command rune) (Command, error) {
	switch string(command) {
	case "L":
		return LeftCommand{}, nil
	case "R":
		return RightCommand{}, nil
	case "F":
		return ForwardCommand{}, nil
	default:
		return nil, fmt.Errorf("Wrong command type passed")
	}
}

package robot

import (
	"testing"

	"github.com/matryer/is"
)

func TestExecuteCommand(t *testing.T) {
	is := is.New(t)
	surface := NewGrid(5, 3, 0, 0)
	surface2 := NewGrid(5, 3, 0, 0)
	positionVisited := Position{x: 3, y: 4}
	surface2.Scents = append(surface2.Scents, positionVisited)
	newRobotWithScents := NewRobot(0, 3, 180, surface2)
	newRobot := NewRobot(1, 1, 0, surface)
	commands := []struct {
		robot       Robot
		commands    []Command
		expected    Robot
		description string
	}{
		{robot: newRobot, commands: []Command{RightCommand{}}, expected: NewRobot(1, 1, -90, surface), description: "Should change orientation from E to S"},
		{robot: newRobot, commands: []Command{LeftCommand{}}, expected: NewRobot(1, 1, 90, surface), description: "Should change orientation from E to N"},
		{robot: newRobot, commands: []Command{ForwardCommand{}}, expected: NewRobot(2, 1, 0, surface), description: "Should change position from 1, 1 to 2, 1 E"},
		{robot: newRobot, commands: []Command{ForwardCommand{}, RightCommand{}}, expected: NewRobot(2, 1, -90, surface), description: "Should change position from 1, 1 to 2, 1 S"},
		{robot: newRobotWithScents, commands: []Command{LeftCommand{}, LeftCommand{}, ForwardCommand{}, ForwardCommand{}, ForwardCommand{}, LeftCommand{}, ForwardCommand{}, LeftCommand{}, ForwardCommand{}, LeftCommand{}}, expected: NewRobot(2, 3, 630, surface2), description: "Should change position from 1, 1 to 2, 1 S"},
	}

	for _, tt := range commands {
		t.Run(tt.description, func(t *testing.T) {
			tt.robot.On(tt.commands)
			is.Equal(tt.robot.orientation, tt.expected.orientation)
			is.Equal(tt.robot.position.x, tt.expected.position.x)
			is.Equal(tt.robot.position.y, tt.expected.position.y)
		})
	}
}

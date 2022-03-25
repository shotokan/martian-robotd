package robot

import (
	"testing"

	"github.com/matryer/is"
)

func TestParseSurface(t *testing.T) {
	is := is.New(t)

	inputs := []struct {
		input       string
		expected    *Grid
		description string
	}{
		{input: "5 3", expected: NewGrid(5, 3, 0, 0), description: "Should return Grid with top 4 and right 5"},
		{input: " 30 10 ", expected: NewGrid(30, 10, 0, 0), description: "Should return Grid with top 30 and right 10"},
		{input: " 30   10 ", expected: NewGrid(30, 10, 0, 0), description: "Should return Grid with top 30 and right 10 when there are more than one space"},
	}

	for _, tt := range inputs {
		t.Run(tt.description, func(t *testing.T) {
			surface := ParseSurface(tt.input)
			is.Equal(surface.top, tt.expected.top)
			is.Equal(surface.right, tt.expected.right)
		})
	}
}

func TestParseRobotPosition(t *testing.T) {
	is := is.New(t)

	type RobotPositionExpected struct {
		x           int32
		y           int32
		orientation int32
	}

	inputs := []struct {
		input       string
		expected    RobotPositionExpected
		description string
	}{
		{input: "1 1 E", expected: RobotPositionExpected{x: 1, y: 1, orientation: 0}, description: "Should return initial position x: 1 y:1 orientation: 0"},
		{input: "1  1 E ", expected: RobotPositionExpected{x: 1, y: 1, orientation: 0}, description: "Should return initial position when extra spaces are passed"},
	}

	for _, tt := range inputs {
		t.Run(tt.description, func(t *testing.T) {
			x, y, orientation := ParseRobotPosition(tt.input)
			is.Equal(x, tt.expected.x)
			is.Equal(y, tt.expected.y)
			is.Equal(orientation, tt.expected.orientation)
		})
	}
}

func TestParseCommand(t *testing.T) {
	is := is.New(t)

	inputs := []struct {
		input       string
		expected    []Command
		description string
	}{
		{input: "R", expected: []Command{RightCommand{}}, description: "Should return RightCommand"},
		{input: "RL", expected: []Command{RightCommand{}, LeftCommand{}}, description: "Should return RightCommand and LeftCommand"},
		{input: "RLF", expected: []Command{RightCommand{}, LeftCommand{}, ForwardCommand{}}, description: "Should return RightCommand, LeftCommand and ForwardCommand"},
	}

	for _, tt := range inputs {
		t.Run(tt.description, func(t *testing.T) {
			commands := ParseCommands(tt.input)
			is.Equal(commands, tt.expected)
		})
	}
}

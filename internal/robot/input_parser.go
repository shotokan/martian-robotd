package robot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	coordinateMaximumValueErrorMsg = "Maximum value for any coordinate is 50"
	notNumberErrorMsg              = "Not a number"
)

// Parse Orientation to radians
func parseOrientationToRadians(orientation string) int32 {
	switch orientation {
	case "N":
		return 90
	case "E":
		return 0
	case "S":
		return 270
	case "W":
		return 180
	default:
		panic("Missing orientation")
	}
}

type generator func() (string, error)

// Function which split values and iterates over them one by one
func Iterator(s string) generator {
	whitespaces := regexp.MustCompile(`\s+`)
	s = whitespaces.ReplaceAllString(s, " ")
	s = strings.Replace(strings.TrimSpace(s), "\r\n", "\n", -1)
	values := strings.Split(s, " ")
	maxIterations := len(values)
	currentIndex := 0
	return func() (string, error) {
		if currentIndex >= maxIterations {
			return "", fmt.Errorf("No more values")
		}
		value := values[currentIndex]
		currentIndex++
		return value, nil
	}
}

// Parses the upper-right coordinates and returns a Grid object
// that is a rectangular and bounded
func ParseSurface(rawSurface string) *Grid {
	nextValue := Iterator(rawSurface)
	xInput, err := nextValue()
	if err != nil {
		panic(err)
	}
	x := parseCoordinate(xInput)

	yInput, err := nextValue()
	if err != nil {
		panic(err)
	}
	y := parseCoordinate(yInput)

	return NewGrid(x, y, 0, 0)
}

// Parses the position and orientation input to integers
func ParseRobotPosition(positionInput string) (int32, int32, int32) {
	nextValue := Iterator(positionInput)

	xInput, err := nextValue()
	if err != nil {
		panic(err)
	}
	x := parseCoordinate(xInput)

	yInput, err := nextValue()
	if err != nil {
		panic(err)
	}
	y := parseCoordinate(yInput)

	orientation, err := nextValue()
	if err != nil {
		panic(err)
	}

	orientationRadians := parseOrientationToRadians(orientation)

	return x, y, orientationRadians
}

// Parses a string number to an integer
func parseCoordinate(v string) int32 {
	if number, err := strconv.Atoi(v); err == nil {
		if number > 50 {
			panic(coordinateMaximumValueErrorMsg)
		}
		return int32(number)
	}
	panic(notNumberErrorMsg)
}

// Parse instruction input to chars to create
// commands that understand the robot and returns a slice of commands
func ParseCommands(rawCommands string) []Command {
	if len(rawCommands) > 100 {
		panic("All instruction strings will be less than 100 characters in length.")
	}
	commands := make([]Command, 0, len(rawCommands))
	for _, rawCommand := range rawCommands {
		command, err := NewCommand(rawCommand)
		if err != nil {
			panic(err)
		}
		commands = append(commands, command)
	}
	return commands
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shotokan/martian-robots/internal/robot"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var surface *robot.Grid
	var newRobot robot.Robot
	initializedRobot := false
	var commands []robot.Command
	var outputs []string

	for {
		fmt.Println()
		scanner.Scan()
		data := scanner.Text()

		if data == "" {
			break
		}

		if surface == nil {
			surface = robot.ParseSurface(data)
			continue
		}

		if !initializedRobot {
			fmt.Println("-- New Robot --")
			fmt.Println(data)
			x, y, orientation := robot.ParseRobotPosition(data)
			newRobot = robot.NewRobot(x, y, orientation, surface)
			initializedRobot = true
			continue
		}

		if initializedRobot {
			fmt.Println(data)
			commands = robot.ParseCommands(data)
			newRobot.On(commands)
			outputs = append(outputs, newRobot.String())
			newRobot.PrintLastPosition()
		}
		initializedRobot = false
	}

	fmt.Println("-- Outputs --")
	for _, o := range outputs {

		fmt.Println(o)
	}
}

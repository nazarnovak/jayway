package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	pkgRobot "github.com/nazarnovak/jayway/backend/pkg/robot"
	pkgRoom "github.com/nazarnovak/jayway/backend/pkg/room"
)

func handleCLIMode() error {
	if err := inputRoomSize(); err != nil {
		return err
	}

	if err := inputRobotPosition(); err != nil {
		return err
	}

	if err := inputRobotNavigation(); err != nil {
		return err
	}

	fmt.Println("Report:", robot.Width, robot.Depth, robot.Orientation)

	return nil
}

// getUserParams parses incoming user input, and separates incoming values into a slice.
func getUserParams() ([]string, error) {
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("Problem with parsing input string: %s", err)
	}

	// Remove the newline character in the end
	input := strings.Replace(line, "\n", "", -1)

	// Separate all incoming values
	params := strings.Split(input, " ")

	// Remove empty space elements
	cleanParams := []string{}
	for _, param := range params {
		if param != "" {
			cleanParams = append(cleanParams, strings.TrimSpace(strings.ToUpper(param)))
		}
	}

	return cleanParams, nil
}

func inputRoomSize() error {
	fmt.Println("Please provide room size (<width> <depth>):")

	params, err := getUserParams()
	if err != nil {
		return err
	}

	if len(params) != 2 {
		return fmt.Errorf("Expecting 2 params, but got %d", len(params))
	}

	width, err := strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		return fmt.Errorf("Problem parsing room width (%s): %s", params[0], err)
	}

	depth, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return fmt.Errorf("Problem parsing depth (%s): %s", params[1], err)
	}

	if err := pkgRoom.ValidateSize(width, depth); err != nil {
		return err
	}

	room = pkgRoom.New(width, depth)

	return nil
}

func inputRobotPosition() error {
	fmt.Println("Please provide starting position for the robot (<width> <depth> <orientation>):")

	params, err := getUserParams()
	if err != nil {
		return err
	}

	if len(params) != 3 {
		return fmt.Errorf("Expecting 3 inputs, but got %d", len(params))
	}

	width, err := strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		return fmt.Errorf("Problem parsing width position (%s): %s", params[0], err)
	}

	depth, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return fmt.Errorf("Problem parsing depth position (%s): %s", params[1], err)
	}

	// First we must make sure the provided coordinates are within the existing room limits
	if width >= room.Width {
		return fmt.Errorf("Robot start width position (%d) cannot be equal or larger than room width (%d)", width,
			room.Width)
	}

	if depth >= room.Depth {
		return fmt.Errorf("Robot start depth position (%d) cannot be equal or larger than room depth (%d)", depth,
			room.Depth)
	}

	orientation := pkgRobot.Orientation(params[2])

	// Note that robot position starts from 0,0
	if err := pkgRobot.ValidateValues(width, depth, orientation); err != nil {
		return err
	}

	robot = pkgRobot.New(width, depth, orientation)

	return nil
}

func inputRobotNavigation() error {
	fmt.Println("Please provide navigation instructions (<INSTRUCTIONS>):")

	params, err := getUserParams()
	if err != nil {
		return err
	}

	if len(params) != 1 {
		return fmt.Errorf("Navigation cannot contain spaces (%s)", strings.Join(params, " "))
	}

	instructions := []pkgRobot.Instruction{}
	for _, letter := range params[0] {
		instructions = append(instructions, pkgRobot.Instruction(letter))
	}

	if err := pkgRobot.ValidateInstructions(instructions); err != nil {
		return err
	}

	if err := robot.Move(instructions, room.Width, room.Depth); err != nil {
		return err
	}

	return nil
}


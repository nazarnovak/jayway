package robot

import (
	"fmt"
	"testing"
)

type moveTestCase struct {
	testName           string
	startingRobot Robot

	// Room limits
	widthLimit int64
	depthLimit int64

	instructions []Instruction
	expectedErrorMessage string
	finishedRobot Robot
}

var moveTestCases = []moveTestCase{
	{
		testName: "Test success 1",
		startingRobot: Robot{
			Width: 1,
			Depth: 2,
			Orientation: North,
		},
		widthLimit: 5,
		depthLimit: 5,
		instructions: []Instruction{right, forward, right, forward, forward, right, forward, right, forward}, // RFRFFRFRF
		expectedErrorMessage: "",
		finishedRobot: Robot{
			Width: 1,
			Depth: 3,
			Orientation: North,
		},
	},
	{
		testName: "Test success 2",
		startingRobot: Robot{
			Width: 0,
			Depth: 0,
			Orientation: East,
		},
		widthLimit: 5,
		depthLimit: 5,
		instructions: []Instruction{right, forward, left, forward, forward, left, right, forward}, // RFLFFLRF
		expectedErrorMessage: "",
		finishedRobot: Robot{
			Width: 3,
			Depth: 1,
			Orientation: East,
		},
	},
	{
		testName: "Unknown instruction",
		startingRobot: Robot{
			Width: 0,
			Depth: 0,
			Orientation: East,
		},
		widthLimit: 5,
		depthLimit: 5,
		instructions: []Instruction{"bad", "instructions"},
		expectedErrorMessage: fmt.Sprintf(errUnknownInstruction, "bad"),
		finishedRobot: Robot{
			Width: 3,
			Depth: 1,
			Orientation: East,
		},
	},
}

type valuesTestCase struct {
	testName string
	width int64
	depth int64
	orientation Orientation
	expectedErrorMessage string
}

var valuesTestCases = []valuesTestCase{
	{
		testName: "Negative width",
		width: -1,
		depth: 1,
		orientation: North,
		expectedErrorMessage: "Width position cannot be negative (-1)",
	},
	{
		testName: "Negative depth",
		width: 1,
		depth: -1,
		orientation: North,
		expectedErrorMessage: "Depth position cannot be negative (-1)",
	},
	{
		testName: "Empty orientation",
		width: 1,
		depth: 1,
		orientation: "",
		expectedErrorMessage: "Orientation cannot be empty",
	},
}

func TestMove(t *testing.T) {
	for _, testCase := range moveTestCases {
		sr := testCase.startingRobot

		err := sr.Move(testCase.instructions, testCase.depthLimit, testCase.widthLimit)

		// If we don't expect an error but it's present
		if testCase.expectedErrorMessage == "" && err != nil {
			t.Errorf("Test '%s': Not expecting an error, but got: '%s'", testCase.testName,
				testCase.expectedErrorMessage)
		}

		if testCase.expectedErrorMessage != "" {
			// If we expect an error but it's not present
			if err == nil {
				t.Errorf("Test '%s': Expecting an error (%s), but error is nil", testCase.testName,
					testCase.expectedErrorMessage)
			}

			// Now we're sure there's an error, we need to compare error messages
			if testCase.expectedErrorMessage != err.Error() {
				t.Errorf("Test '%s': Expecting error message (%s), but got (%s)", testCase.testName,
					testCase.expectedErrorMessage, err.Error())
			}
		}

		// Code below assertions need to run only if we don't have an error, so we can short circuit here
		if testCase.expectedErrorMessage != "" {
			return
		}

		if testCase.finishedRobot.Width != sr.Width {
			t.Errorf("Test '%s': Expecting finished robot width to be '%d', got: '%d'", testCase.testName,
				testCase.finishedRobot.Width, sr.Width)
		}

		if testCase.finishedRobot.Depth != sr.Depth {
			t.Errorf("Test '%s': Expecting finished robot depth to be '%d', got: '%d'", testCase.testName,
				testCase.finishedRobot.Depth, sr.Depth)
		}

		if testCase.finishedRobot.Orientation != sr.Orientation {
			t.Errorf("Test '%s': Expecting finished robot orientation to be '%s', got: '%s'", testCase.testName,
				testCase.finishedRobot.Orientation, sr.Orientation)
		}
	}
}

func TestValidateValues(t *testing.T) {
	for _, testCase := range valuesTestCases {
		err := ValidateValues(testCase.width, testCase.depth, testCase.orientation)

		// If we don't expect an error but it's present
		if testCase.expectedErrorMessage == "" && err != nil {
			t.Errorf("Test '%s': Not expecting an error, but got: '%s'", testCase.testName,
				testCase.expectedErrorMessage)
		}

		if testCase.expectedErrorMessage != "" {
			// If we expect an error but it's not present
			if err == nil {
				t.Errorf("Test '%s': Expecting an error (%s), but error is nil", testCase.testName,
					testCase.expectedErrorMessage)
			}

			// Now we're sure there's an error, we need to compare error messages
			if testCase.expectedErrorMessage != err.Error() {
				t.Errorf("Test '%s': Expecting error message (%s), but got (%s)", testCase.testName,
					testCase.expectedErrorMessage, err.Error())
			}
		}
	}
}
package robot

import (
	"fmt"
)

const (
	North Orientation = "N"
	East  Orientation = "E"
	South Orientation = "S"
	West  Orientation = "W"

	forward Instruction = "F"
	left    Instruction = "L"
	right   Instruction = "R"
)

type Orientation string
type Instruction string

var (
	allowedOrientations = []Orientation{North, East, South, West}
	allowedInstructions = []Instruction{forward, left, right}

	errUnknownInstruction = "Unknown instruction (%s)"
)

type Robot struct {
	Width          int64 `json:"width"`
	Depth          int64 `json:"depth"`
	Orientation    Orientation `json:"orientation"`
	orientationKey int
}

// New creates a new instance of Robot.
func New(width, depth int64, orientation Orientation) *Robot {
	return &Robot{
		Width:       width,
		Depth:       depth,
		Orientation: orientation,
	}
}

// setOrientationKey is a helper function, that maps the orientation to a number, which makes it easier to perform
// rotations.
func (r *Robot) setOrientationKey() {
	// Reset the key, if it was already set
	r.orientationKey = 0

	for k, allowedOrientation := range allowedOrientations {
		if allowedOrientation == r.Orientation {
			r.orientationKey = k
			break
		}
	}
}

// rotateLeft rotates the robot left.
func (r *Robot) rotateLeft() {
	r.orientationKey--

	if r.orientationKey < 0 {
		r.orientationKey = len(allowedOrientations) - 1
	}

	r.Orientation = allowedOrientations[r.orientationKey]
}

// rotateLeft rotates the robot right.
func (r *Robot) rotateRight() {
	r.orientationKey++

	if r.orientationKey > len(allowedOrientations)-1 {
		r.orientationKey = 0
	}

	r.Orientation = allowedOrientations[r.orientationKey]
}

// moveForward moves the robot forward.
func (r *Robot) moveForward(widthLimit, depthLimit int64) error {
	switch r.Orientation {
	case North:
		if r.Depth > 0 {
			r.Depth--
		}
	case South:
		if r.Depth < depthLimit-1 {
			r.Depth++
		}
	case West:
		if r.Width > 0 {
			r.Width--
		}
	case East:
		if r.Width < widthLimit-1 {
			r.Width++
		}
	default:
		return fmt.Errorf(errUnknownInstruction, r.Orientation)
	}

	return nil
}

// Move performs any robot movement or rotations.
func (r *Robot) Move(insts []Instruction, widthLimit, depthLimit int64) error {
	r.setOrientationKey()

	for _, inst := range insts {
		switch inst {
		case left:
			r.rotateLeft()
		case right:
			r.rotateRight()
		case forward:
			if err := r.moveForward(widthLimit, depthLimit); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unknown instruction (%s)", inst)
		}
	}

	return nil
}

// ValidateValues checks if width, depth, and orientation are valid for a robot.
func ValidateValues(width, depth int64, orientation Orientation) error {
	if width < 0 {
		return fmt.Errorf("Width position cannot be negative (%d)", width)
	}

	if depth < 0 {
		return fmt.Errorf("Depth position cannot be negative (%d)", depth)
	}

	if orientation == "" {
		return fmt.Errorf("Orientation cannot be empty")
	}

	isValid := false
	for _, allowedOrientation := range allowedOrientations {
		if allowedOrientation == orientation {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("Invalid orientation (%s). Valid orientations are: %s", orientation, allowedOrientations)
	}

	return nil
}

// ValidateInstructions validates any rotation or movement instructions.
func ValidateInstructions(insts []Instruction) error {
	if len(insts) == 0 {
		return fmt.Errorf("No instructions provided")
	}

	for _, inst := range insts {
		isAllowed := false

		for _, allowedInstruction := range allowedInstructions {
			if inst == allowedInstruction {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return fmt.Errorf("Invalid instruction: %s", inst)
		}
	}

	return nil
}

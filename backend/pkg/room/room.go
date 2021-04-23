package room

import "fmt"

type Room struct {
	Width int64 `json:"width"`
	Depth int64 `json:"depth"`
}

func New(width, depth int64) Room {
	return Room{
		Width: width,
		Depth: depth,
	}
}

// ValidateSize checks if width and depth are valid for a room.
func ValidateSize(width, depth int64) error {
	if width < 1 {
		return fmt.Errorf("Width cannot be negative or 0 (%d)", width)
	}

	if depth < 1 {
		return fmt.Errorf("Depth cannot be negative or 0 (%d)", depth)
	}

	return nil
}

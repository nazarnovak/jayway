package api

import (
	"encoding/json"
	"fmt"
	pkgRobot "github.com/nazarnovak/jayway/backend/pkg/robot"
	pkgRoom "github.com/nazarnovak/jayway/backend/pkg/room"
	"net/http"
)

type RobotRequest struct {
	pkgRoom.Room   `json:"room"`
	pkgRobot.Robot `json:"robot"`
	Instructions   string `json:"instructions"`
}

type RobotResponse struct {
	Error   bool           `json:"error"`
	Message string         `json:"message"`
	Report  pkgRobot.Robot `json:"report,omitempty"`
}

func RobotHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rr := RobotRequest{}

		if err := json.NewDecoder(r.Body).Decode(&rr); err != nil {
			fmt.Println(err)

			sendResponse(w, http.StatusBadRequest, true, "Could not parse incoming data", pkgRobot.Robot{})
			return
		}
		fmt.Println("Received:", rr)
		if err := pkgRoom.ValidateSize(rr.Room.Width, rr.Room.Depth); err != nil {
			fmt.Println(err)

			sendResponse(w, http.StatusBadRequest, true, err.Error(), pkgRobot.Robot{})
			return
		}

		// First we must make sure the provided coordinates are within the existing room limits
		if rr.Robot.Width >= rr.Room.Width {
			msg := fmt.Sprintf("Robot start width position (%d) cannot be equal or larger than room width (%d)",
				rr.Robot.Width, rr.Room.Width)
			fmt.Println(msg)

			sendResponse(w, http.StatusBadRequest, true, fmt.Sprintf(msg), pkgRobot.Robot{})
			return
		}

		if rr.Robot.Depth >= rr.Room.Depth {
			msg := fmt.Sprintf("Robot start depth position (%d) cannot be equal or larger than room depth (%d)",
				rr.Robot.Depth, rr.Room.Depth)

			fmt.Println(msg)

			sendResponse(w, http.StatusBadRequest, true, fmt.Sprintf(msg), pkgRobot.Robot{})
			return
		}

		if err := pkgRobot.ValidateValues(rr.Robot.Width, rr.Robot.Depth, rr.Robot.Orientation); err != nil {
			fmt.Println(err)

			sendResponse(w, http.StatusBadRequest, true, err.Error(), pkgRobot.Robot{})
			return
		}

		instructions := []pkgRobot.Instruction{}
		for _, letter := range rr.Instructions {
			instructions = append(instructions, pkgRobot.Instruction(letter))
		}

		if err := pkgRobot.ValidateInstructions(instructions); err != nil {
			fmt.Println(err)

			sendResponse(w, http.StatusBadRequest, true, err.Error(), pkgRobot.Robot{})
			return
		}

		if err := rr.Robot.Move(instructions, rr.Room.Width, rr.Room.Depth); err != nil {
			fmt.Println(err)

			sendResponse(w, http.StatusBadRequest, true, err.Error(), pkgRobot.Robot{})
			return
		}

		sendResponse(w, http.StatusCreated, false, "", rr.Robot)
	}
}

func sendResponse(w http.ResponseWriter, statusCode int, isError bool, msg string, rob pkgRobot.Robot) {
	er := RobotResponse{
		Error:   isError,
		Message: msg,
		Report:  rob,
	}

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(er); err != nil {
		fmt.Println(err)
		return
	}

	return
}

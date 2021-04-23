package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"

	pkgRobot "github.com/nazarnovak/jayway/backend/pkg/robot"
	pkgRoom "github.com/nazarnovak/jayway/backend/pkg/room"
)

type robotTestCase struct {
	testName           string
	req                RobotRequest
	expectedStatusCode int
	expectedError      bool
	expectedMessage    string
	expectedReport     pkgRobot.Robot
}

var robotTestCases = []robotTestCase{
	{
		testName:           "Success 1",
		req:                newRobotRequest(5, 5, 1, 2, pkgRobot.North, "RFRFFRFRF"),
		expectedStatusCode: http.StatusCreated,
		expectedError:      false,
		expectedMessage:    "",
		expectedReport: pkgRobot.Robot{
			Width:       1,
			Depth:       3,
			Orientation: pkgRobot.North,
		},
	},
	{
		testName:           "Success 2",
		req:                newRobotRequest(5, 5, 0, 0, pkgRobot.East, "RFLFFLRF"),
		expectedStatusCode: http.StatusCreated,
		expectedError:      false,
		expectedMessage:    "",
		expectedReport: pkgRobot.Robot{
			Width:       3,
			Depth:       1,
			Orientation: pkgRobot.East,
		},
	},

}

func newRobotRequest(
	roomWidth,
	roomDepth,
	robotWidth,
	robotDepth int64,
	orientation pkgRobot.Orientation,
	instructions string,
) RobotRequest {
	return RobotRequest{
		Room: pkgRoom.Room{
			Width: roomWidth,
			Depth: roomDepth,
		},
		Robot: pkgRobot.Robot{
			Width:       robotWidth,
			Depth:       robotDepth,
			Orientation: pkgRobot.Orientation(orientation),
		},
		Instructions: instructions,
	}
}

func TestRobotHandler(t *testing.T) {
	mux := web.New()
	mux.Post("/api/robot", RobotHandler())

	for _, testCase := range robotTestCases {
		b, err := json.Marshal(testCase.req)
		if err != nil {
			t.Errorf("Test '%s': Expecting error to be nil, got: %s", testCase.testName, err)
		}

		req, _ := http.NewRequest(http.MethodPost, "/api/robot", bytes.NewReader(b))
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)
		if status := rr.Code; status != testCase.expectedStatusCode {
			t.Errorf("Test '%s': Expecting status: %d got: %d", testCase.testName, http.StatusCreated,
				rr.Code)
		}

		var robresp RobotResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &robresp); err != nil {
			t.Errorf("Test '%s': Unexpected error: %w", testCase.testName, err)
		}

		if testCase.expectedError != robresp.Error {
			t.Errorf("Test '%s': Unexpected error in response: %s", testCase.testName, robresp.Message)
		}

		if testCase.expectedMessage != robresp.Message {
			t.Errorf("Test '%s': Unexpected message in response: %s", testCase.testName, robresp.Message)
		}

		if testCase.expectedReport != robresp.Report {
			t.Errorf("Test '%s':Expected report to be '%v', got - '%v'", testCase.testName, testCase.expectedReport,
				robresp.Report)
		}
	}
}

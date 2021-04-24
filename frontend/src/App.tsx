import React, { useState } from 'react';
import axios, { AxiosError } from 'axios';

import styles from './App.module.css';

import { Explanation } from 'components/explanation';
import { Controls } from 'components/controls';
import { Display } from 'components/display';

const apiSave = '/api/robot';

const allowedOrientations = ["N", "E", "S", "W"];
const instructionsRegex = new RegExp("^[lrf]+$", "i");

function App() {
  const [room, setRoom] = useState({width: 0, depth: 0});
  const [robot, setRobot] = useState({width: -1, depth: -1, orientation: "N"});
  const [instructions, setInstructions] = useState('');
  const [error, setError] = useState('');
  const [movedRobot, setMovedRobot] = useState({width: -1, depth: -1, orientation: "N"});

  const handleRoomWidthChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRoom({ ...room, width: +e.target.value });
  }

  const handleRoomDepthChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRoom({ ...room, depth: +e.target.value });
  }

  const handleRobotWidthChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRobot({ ...robot, width: +e.target.value });
  }

  const handleRobotDepthChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRobot({ ...robot, depth: +e.target.value });
  }

  const handleRobotOrientationChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setRobot({ ...robot, orientation: e.target.value });
  }

  const handleInstructionsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInstructions(e.target.value );
  }

  const handleMove = () => {
    if (room.width < 1) {
      setError('Room width needs to be more than 0');
      return;
    }

    if (room.depth < 1) {
      setError('Room depth needs to be more than 0');
      return;
    }

    if (robot.width < 0) {
      setError('Robot width needs to be 0 or higher');
      return;
    }

    if (robot.width > room.width - 1) {
      setError('Robot width cannot be higher than room width');
      return;
    }

    if (robot.depth < 0) {
      setError('Robot width needs to be 0 or higher');
      return;
    }

    if (robot.depth > room.depth - 1) {
      setError('Robot depth cannot be higher than room depth');
      return;
    }

    if (!allowedOrientations.includes(robot.orientation)) {
      setError('Wrong direction (' + robot.orientation + '). Allowed orientations: ' +
        allowedOrientations);
      return;
    }

    if (instructions === '') {
      setError('Instructions cannot be empty');
      return;
    }

    if (!instructions.match(instructionsRegex)) {
      setError('Instructions can only contain following characters: F,L,R');
      return;
    }

    setError('');

    axios
      .post(apiSave, {room, robot, instructions})
      .then(function (resp) {
        setMovedRobot(resp?.data?.report);
      })
      .catch(function (error: AxiosError) {
        let resp = error.response?.data;

        if (resp?.error) {
          setError(resp?.message);
          return;
        }
      });
  }

  return (
    <div className={styles.mainContainer}>
      <h2>Move a robot in a room coding challenge</h2>
      <div className={styles.sectionsContainer}>
        <div className={styles.controlSection}>
          <Explanation />
          <Controls handleRoomWidthChange={handleRoomWidthChange} handleRoomDepthChange={handleRoomDepthChange}
                   handleRobotWidthChange={handleRobotWidthChange} handleRobotDepthChange={handleRobotDepthChange}
                   handleRobotOrientationChange={handleRobotOrientationChange}
                   handleInstructionsChange={handleInstructionsChange} handleMove={handleMove}
                   error={error}
          />
        </div>
        <div className={styles.displaySection}>
          <Display room={room} robot={robot} movedRobot={movedRobot} />
        </div>
      </div>
    </div>
  );
}

export default App;

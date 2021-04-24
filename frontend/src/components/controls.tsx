import React from "react";

import styles from "./controls.module.css";

type Props = {
  handleRoomWidthChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleRoomDepthChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleRobotWidthChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleRobotDepthChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleRobotOrientationChange: (
    e: React.ChangeEvent<HTMLSelectElement>
  ) => void;
  handleInstructionsChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  error: string;
  handleMove: () => void;
};

export const Controls: React.FunctionComponent<Props> = (props: Props) => {
  return (
    <div className={styles.controls}>
      <h4>Room</h4>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Width</div>
        <input
          type="number"
          name="room-width"
          onChange={props.handleRoomWidthChange}
        />
      </div>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Depth</div>
        <input
          type="number"
          name="room-depth"
          onChange={props.handleRoomDepthChange}
        />
      </div>

      <h4>Robot</h4>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Width</div>
        <input
          type="number"
          name="robot-width"
          onChange={props.handleRobotWidthChange}
        />
      </div>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Depth</div>
        <input
          type="number"
          name="robot-depth"
          onChange={props.handleRobotDepthChange}
        />
      </div>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Orientation</div>
        <select
          name="robot-orientation"
          onChange={props.handleRobotOrientationChange}
        >
          <option value="N">North</option>
          <option value="E">East</option>
          <option value="S">South</option>
          <option value="W">West</option>
        </select>
      </div>

      <h4>Movement instructions (FLR)</h4>
      <div className={styles.controlGroup}>
        <div className={styles.controlLabel}>Instructions</div>
        <input name="instructions" onChange={props.handleInstructionsChange} />
      </div>

      <div className={styles.error}>{props.error}</div>

      <div className={styles.buttonContainer}>
        <button className={styles.buttonAction} onClick={props.handleMove}>
          Move the robot
        </button>
      </div>
    </div>
  );
};

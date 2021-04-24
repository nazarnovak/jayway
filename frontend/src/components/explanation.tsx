import React from "react";

import styles from "./explanation.module.css";

export const Explanation: React.FunctionComponent = () => {
  return (
    <div className={styles.explanation}>
      <div>How it works:</div>
      <ol>
        <li>
          Create room dimensions, by providing the width and depth of the room
        </li>
        <li>
          Provide the starting position of the robot, and which way it faces
        </li>
        <li>
          Provide movement instructions
          <ul>
            <li>F (move forward)</li>
            <li>L (rotate left)</li>
            <li>R (rotate right)</li>
          </ul>
        </li>
      </ol>
    </div>
  );
};

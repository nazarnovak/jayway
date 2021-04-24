import React from 'react';

import styles from './display.module.css';

type Props = {
  room: {
    width: number;
    depth: number;
  };
  robot: {
    width: number;
    depth: number;
    orientation: string;
  };
  movedRobot: {
    width: number;
    depth: number;
    orientation: string;
  };
};

export const Display: React.FunctionComponent<Props> = (props: Props) => {
  let rows = [];

  // Only draw the table, when there's at least 1x1 size
  if (props.room.width < 1 || props.room.depth < 1) {
    return (
      <div className={styles.displayHelper}>
        Please specify room width and depth larger than 0
      </div>
    );
  }

  for (var i = 0; i < props.room.depth; i++) {
    let cols = [];

    for (var j = 0; j < props.room.width; j++) {
      // Places the robot/moved robot to the correct cell
      const robot =
        i === props.robot.depth && j === props.robot.width ? (
          <span className={styles.startRobot}>{props.robot.orientation}</span>
        ) : null;
      const movedRobot =
        i === props.movedRobot.depth && j === props.movedRobot.width ? (
          <span className={styles.movedRobot}>{props.movedRobot.orientation}</span>
        ) : null;

      // TODO: Right now if the moved robot will be in the same position as starting point, it will overlap, so moved
      // robot takes priority.
      cols.push(
        <td key={j}>
          {movedRobot ? movedRobot : robot ? robot : null}
        </td>
      );
    }

    rows.push(
      <tr key={i}>
        {cols}
      </tr>
    );
  }

  return (
    <div>
      <table>
        <tbody>{rows}</tbody>
      </table>
      <div className={styles.displayLegend}>
        <div>
          <span className={styles.startRobot}>N</span> - starting robot position
        </div>
        <div>
          <span className={styles.movedRobot}>N</span> - moved robot position
        </div>
      </div>
    </div>
  );
};

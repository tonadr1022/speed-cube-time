import { CubeSession, Settings } from "../../types/types";
import clsx from "clsx";
import ReactList from "react-list";
import { FaRegCirclePlay } from "react-icons/fa6";
import { FaRegCheckCircle, FaRegEdit } from "react-icons/fa";

type Props = {
  className?: string;
  sessions: CubeSession[];
  settings: Settings;
  setActiveSession: (id: string) => void;
};

const SessionsTable = ({
  className,
  sessions,
  settings,
  setActiveSession,
}: Props) => {
  const renderSessionRow = (i: number, key: number | string) => {
    const active = settings.active_cube_session_id == sessions[i].id;
    return (
      <div
        className={clsx(
          "flex flex-row p-2 rounded-lg font-bold border-2 box-border",
          active ? "border-primary" : "border-transparent",
        )}
        key={key}
      >
        <span className="flex-1">{sessions[i].name}</span>
        <button className="btn btn-xs hover:bg-info ">
          <FaRegEdit />
        </button>
        <button
          onMouseDown={() => {
            if (!active) setActiveSession(sessions[i].id);
          }}
          className={clsx(
            "btn btn-xs",
            active ? "bg-success hover:bg-success" : "hover:bg-success",
          )}
        >
          {active ? <FaRegCheckCircle /> : <FaRegCirclePlay />}
        </button>
      </div>
    );
  };

  return (
    <div className={clsx("rounded-lg bg-base-200", className)}>
      <ReactList
        itemRenderer={renderSessionRow}
        length={sessions.length}
        type="uniform"
      />
    </div>
  );
};
export default SessionsTable;

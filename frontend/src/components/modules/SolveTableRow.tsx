import { formatTime } from "../../util/time";
import clsx from "clsx";
import { Solve } from "../../types/types";
import { FaTrash } from "react-icons/fa6";
interface Props {
  solve: Solve;
  onTogglePlusTwo: (solve: Solve) => void;
  onToggleDnf: (solve: Solve) => void;
  onDelete: (solveId: string) => void;
  solveCount: number;
}

const SolveTableRow = ({
  solve,
  onTogglePlusTwo,
  onToggleDnf,
  onDelete,
  solveCount,
}: Props) => {
  return (
    solve && (
      <div className="w-full flex flex-row py-0.5 items-center" key={solve.id}>
        <div className=" pl-4 pr-8 w-6 font-bold flex items-center">
          {solveCount}.
        </div>
        <div
          className={clsx(
            "px-4 flex-1 font-medium transition flex items-center",
            solve.dnf
              ? "text-error"
              : solve.plus_two
                ? "text-warning"
                : "text-success",
          )}
        >
          {formatTime(solve.duration, 3)}
        </div>
        <div className="flex flex-row">
          <div>
            <button
              className={clsx(
                "btn btn-xs bg-base-300 border-none hover:bg-base-100 transition-none",
                solve.plus_two && "bg-warning text-neutral",
              )}
              onClick={() => onTogglePlusTwo(solve)}
            >
              +2
            </button>
          </div>
          <div>
            <button
              className={clsx(
                "btn btn-xs bg-base-300 border-none hover:bg-base-100 transition-none",
              )}
              onClick={() => onDelete(solve.id)}
            >
              <FaTrash />
            </button>
          </div>
          <div>
            <button
              className={clsx(
                "btn btn-xs bg-base-300 border-none hover:bg-base-100",
                solve.dnf && "bg-error text-neutral",
              )}
              onClick={() => onToggleDnf(solve)}
            >
              DNF
            </button>
          </div>
        </div>
      </div>
    )
  );
};

export default SolveTableRow;

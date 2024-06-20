import { formatTime } from "../../util/time";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";
import { useTimerContext } from "../../hooks/useContext";
interface Props {
  duration: number;
  state: string;
}
const DurationDisplay = ({ duration, state }: Props) => {
  const { cubeType } = useTimerContext();
  const digits = state === "Active" ? 1 : 2;
  return (
    <div>
      <h2
        className={twMerge(
          clsx(
            state === "Ready" && "text-success",
            state === "Stalling" && "text-warning",
          ),
          "text-6xl font-mono font-semibold w-75",
        )}
      >
        {formatTime(duration ? duration : 0, digits)}
      </h2>
      <div>{cubeType}</div>
    </div>
  );
};

export default DurationDisplay;

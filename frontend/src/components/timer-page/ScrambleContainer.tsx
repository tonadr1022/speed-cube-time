import React, { useEffect } from "react";
import { FaArrowRotateRight, FaCopy } from "react-icons/fa6";
import { getScramble } from "../../util/getScramble";
import { useTimerContext } from "../../hooks/useContext";

const ScrambleContainer = React.memo(() => {
  const [resetScramble, setResetScramble] = React.useState(false);
  const { scramble, setScramble, cubeType } = useTimerContext();

  useEffect(() => {
    if (cubeType) {
      const generateScramble = () => {
        const initialScramble = getScramble({
          cubeType,
        });
        setScramble(initialScramble);
      };
      generateScramble();
    }
  }, [setScramble, cubeType, resetScramble]);

  return (
    <div className="flex flex-col items-center jutify-center">
      <h2 className="text-xl">{scramble || ""}</h2>
      <div className="flex">
        <button
          className="btn btn-sm btn-neutral-focus p-1 m-0 rounded-full"
          onClick={() => navigator.clipboard.writeText(scramble)}
        >
          <FaCopy className="w-6 h-6" />
        </button>
        <button
          className="btn btn-sm btn-neutral-focus p-1 m-0 rounded-full"
          onClick={() => setResetScramble(!resetScramble)}
        >
          <FaArrowRotateRight className="w-6 h-6" />
        </button>
      </div>
    </div>
  );
});

export default ScrambleContainer;

ScrambleContainer.displayName = "ScrambleContainer";

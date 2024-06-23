import React, { useEffect, useState } from "react";
import { FaArrowRotateRight, FaCheck, FaCopy } from "react-icons/fa6";
import { getScramble } from "../../util/getScramble";
import { useTimerContext } from "../../hooks/useContext";

const ScrambleContainer = React.memo(() => {
  const [resetScramble, setResetScramble] = useState(false);
  const [clipboardCopied, setClipboardCompied] = useState(false);
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

  const handleCopy = () => {
    navigator.clipboard.writeText(scramble);
    setClipboardCompied(true);
    setTimeout(() => setClipboardCompied(false), 400);
  };

  return (
    <div className="flex flex-col items-center jutify-center">
      <h2 className="text-xl">{scramble || ""}</h2>
      <div
        onTouchStart={(e) => e.stopPropagation()}
        onTouchEnd={(e) => e.stopPropagation()}
        className="flex flex-row gap-2"
      >
        <button
          className="btn btn-sm btn-neutral-focus p-1 m-0 rounded-full"
          onClick={handleCopy}
        >
          {clipboardCopied ? (
            <FaCheck className="w-6 h-6 " />
          ) : (
            <FaCopy className="w-6 h-6" />
          )}
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

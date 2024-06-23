import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useAddSolveMutation, useFetchSettings } from "../hooks/useFetch.ts";
import DurationDisplay from "../components/timer-page/DurationDisplay.tsx";
import { useSpaceBarDown, useSpaceBarUp } from "../hooks/timerHooks.ts";
import { useTimerContext } from "../hooks/useContext.ts";
import { getScramble } from "../util/getScramble.ts";
import ScrambleContainer from "./timer-page/ScrambleContainer.tsx";

type TimerState = "Active" | "Ready" | "Stalling" | "Paused" | "Initial";

const Timer = () => {
  const { scramble, setScramble, keybindsActive, cubeType } = useTimerContext();
  const { data: settings } = useFetchSettings();
  const [timerState, setTimerState] = useState<TimerState>("Initial");
  const [duration, setDuration] = useState<number>(0);
  const [timerTimeoutId, setTimerTimeoutId] = useState<number | null>(null);
  const [timerIntervalId, setTimerIntervalId] = useState<number | null>(null);
  const addSolveMutation = useAddSolveMutation(useQueryClient());

  const updateTimer = (start: number) => {
    const id = setInterval(() => {
      setDuration((Date.now() - start) / 1000);
    }, 10) as unknown as number;
    setTimerIntervalId(id);
  };

  const handleDown = () => {
    if (timerState === "Active") {
      // solve was active and now finished
      if (timerTimeoutId) clearTimeout(timerTimeoutId);
      if (timerIntervalId) clearInterval(timerIntervalId);

      // for (let i = 0; i < 10; i++) {
      addSolveMutation.mutate({
        duration: duration,
        scramble: scramble,
        cube_session_id: settings?.active_cube_session_id || "",
        cube_type: cubeType,
        dnf: false,
        plus_two: false,
        notes: "",
      });
      // }

      setTimerState("Stalling");
      // set new scramble
      setScramble(getScramble({ cubeType }));
    } else if (
      (keybindsActive && timerState === "Initial") ||
      timerState === "Paused"
    ) {
      // timer at 0, reading to turn red before starting
      setDuration(0);
      setTimerState("Stalling");
      setTimerTimeoutId(
        setTimeout(() => {
          setTimerState("Ready");
        }, 300) as unknown as number | null,
      );
    }
  };

  const handleUp = () => {
    if (timerTimeoutId) {
      clearTimeout(timerTimeoutId);
      setTimerTimeoutId(null);
    }
    // If ready to start, start
    if (timerState === "Ready") {
      setTimerState("Active");
      updateTimer(Date.now());
      // Else back to initial state
    } else {
      setTimerState("Initial");
    }
  };

  useSpaceBarDown(handleDown, keybindsActive);
  useSpaceBarUp(handleUp, keybindsActive);
  return (
    <div
      onTouchStart={() => handleDown()}
      onTouchEnd={() => handleUp()}
      className="prevent-select flex flex-col gap-y-4 justify-center items-center text-center flex-1"
    >
      <ScrambleContainer />
      <DurationDisplay duration={duration} state={timerState} />
    </div>
  );
};

export default Timer;

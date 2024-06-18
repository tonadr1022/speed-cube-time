import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useAddSolveMutation, useFetchSettings } from "../hooks/useFetch.tsx";
import DurationDisplay from "../components/timer-page/DurationDisplay.tsx";
import { useSpaceBarDown, useSpaceBarUp } from "../hooks/timerHooks.ts";
import { useTimerContext } from "../hooks/useContext.ts";

type TimerState = "Active" | "Ready" | "Stalling" | "Paused" | "Initial";

const Timer = () => {
  const { scramble, timerCanStart, cubeType } = useTimerContext();
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

      addSolveMutation.mutate({
        duration: duration,
        scramble: scramble,
        cube_session_id: settings?.active_cube_session_id || "",
        cube_type: cubeType,
        dnf: false,
        plus_two: false,
        notes: "notes",
      });

      setTimerState("Stalling");
      // set new scramble
    } else if (
      (timerCanStart && timerState === "Initial") ||
      timerState === "Paused"
    ) {
      // timer at 0, reading to turn red before starting
      setDuration(0);
      setTimerState("Stalling");
      setTimerTimeoutId(
        setTimeout(() => {
          setTimerState("Ready");
        }, 300),
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

  useSpaceBarDown(handleDown, timerCanStart);
  useSpaceBarUp(handleUp);
  return (
    <div
      onTouchStart={() => handleDown()}
      onTouchEnd={() => handleUp()}
      className="flex flex-col justify-center items-center text-center flex-1"
    >
      <DurationDisplay duration={duration} state={timerState} />
    </div>
  );
};

export default Timer;
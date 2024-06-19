import { createContext, useState } from "react";

interface TimerContextType {
  keybindsActive: boolean;
  setKeybindsActive: (state: boolean) => void;
  scramble: string;
  setScramble: (scramble: string) => void;
  cubeType: string;
  setCubeType: (cubeType: string) => void;
}

export const TimerContext = createContext<TimerContextType | null>(null);

export const TimerContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [keybindsActive, setKeybindsActive] = useState<boolean>(true);
  const [scramble, setScramble] = useState<string>("");
  const [cubeType, setCubeType] = useState<string>("333");
  return (
    <TimerContext.Provider
      value={{
        keybindsActive,
        setKeybindsActive,
        scramble,
        setScramble,
        cubeType,
        setCubeType,
      }}
    >
      {children}
    </TimerContext.Provider>
  );
};

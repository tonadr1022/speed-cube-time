import { SetStateAction, createContext, useState } from "react";
import usePersistState from "../hooks/usePersistState";

interface SettingsContextType {
  theme: string;
  setTheme: React.Dispatch<SetStateAction<string>>;
  focusMode: boolean;
  setFocusMode: React.Dispatch<SetStateAction<boolean>>;
  modules: string[];
  setModules: React.Dispatch<SetStateAction<string[]>>;
  moduleCount: number;
  setModuleCount: React.Dispatch<SetStateAction<number>>;
  display3D: boolean;
  setDisplay3D: React.Dispatch<SetStateAction<boolean>>;
  online: boolean;
  setOnline: React.Dispatch<SetStateAction<boolean>>;
}

export const SettingsContext = createContext<SettingsContextType | undefined>(
  undefined,
);

export const SettingsProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  // const [theme, setTheme] = useState<string>("dark");
  const [theme, setTheme] = usePersistState("theme", "light");

  const [focusMode, setFocusMode] = useState<boolean>(false);
  const [modules, setModules] = useState<string[]>([
    "solves",
    "stats",
    "cubeDisplay",
    "timeGraph",
    "none",
  ]);
  const [display3D, setDisplay3D] = useState<boolean>(true);
  const [online, setOnline] = useState<boolean>(true);
  const [moduleCount, setModuleCount] = useState<number>(3);

  return (
    <SettingsContext.Provider
      value={{
        online,
        setOnline,
        theme,
        setTheme,
        focusMode,
        setFocusMode,
        modules,
        setModules,
        display3D,
        setDisplay3D,
        moduleCount,
        setModuleCount,
      }}
    >
      {children}
    </SettingsContext.Provider>
  );
};

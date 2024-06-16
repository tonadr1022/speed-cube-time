import { createContext, useState } from "react";

interface SettingsContextType {
  theme: string;
  setTheme: (theme: string) => void;
  focusMode: boolean;
  setFocusMode: (focus: boolean) => void;
  modules: string[];
  setModules: (modules: string[]) => void;
  moduleCount: number;
  setModuleCount: (moduleCount: number) => void;
  display3D: boolean;
  setDisplay3D: (is3D: boolean) => void;
}

export const SettingsContext = createContext<SettingsContextType | undefined>(
  undefined,
);

export const SettingsProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [theme, setTheme] = useState<string>("");
  const [focusMode, setFocusMode] = useState<boolean>(false);
  const [modules, setModules] = useState<string[]>([
    "solves",
    "stats",
    "cubeDisplay",
    "timeGraph",
    "none",
  ]);
  const [display3D, setDisplay3D] = useState<boolean>(true);
  const [moduleCount, setModuleCount] = useState<number>(3);

  return (
    <SettingsContext.Provider
      value={{
        theme,
        setTheme,
        focusMode: focusMode,
        setFocusMode: setFocusMode,
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

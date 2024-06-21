import { SetStateAction, createContext } from "react";
import usePersistState from "../hooks/usePersistState";
import { useQueryClient } from "@tanstack/react-query";

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
  const [focusMode, setFocusMode] = usePersistState("focus", false);
  const [modules, setModules] = usePersistState("modules", [
    "solves",
    "stats",
    "cubeDisplay",
    "timeGraph",
    "none",
  ]);
  const [display3D, setDisplay3D] = usePersistState("display3D", true);
  const queryClient = useQueryClient();
  const [online, setOnline] = usePersistState("online", true, () => {
    queryClient.invalidateQueries({ queryKey: ["solves"] });
    queryClient.invalidateQueries({ queryKey: ["cubeSessions"] });
    queryClient.invalidateQueries({ queryKey: ["settings"] });
  });
  const [moduleCount, setModuleCount] = usePersistState("moduleCount", 3);

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

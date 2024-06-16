import { createContext, useState } from "react";

interface SettingsContextType {
  theme: string;
  focus: boolean;
  setTheme: (theme: string) => void;
  setFocus: (focus: boolean) => void;
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
  const [focus, setFocus] = useState<boolean>(false);

  return (
    <SettingsContext.Provider value={{ theme, setTheme, focus, setFocus }}>
      {children}
    </SettingsContext.Provider>
  );
};

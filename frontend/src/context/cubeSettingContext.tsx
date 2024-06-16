import { createContext, useEffect, useState } from "react";
import localforage from "localforage";

interface CubeSettingsContextType {
  activeCubeType: string;
  setActiveCubeType: (type: string) => void;
}

export const CubeSettingsContext = createContext<
  CubeSettingsContextType | undefined
>(undefined);

export const SettingsProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [activeCubeType, setActiveCubeType] = useState<string>("333");
  useEffect(() => {
    localforage.getItem<string>("activeCubeType").then((value) => {
      if (value) {
        setActiveCubeType(value);
      }
    });
  }, []);
  const handleSetActiveCubeType = (type: string) => {
    setActiveCubeType(type);
    localforage.setItem("activeCubeType", type);
  };

  return (
    <CubeSettingsContext.Provider
      value={{ activeCubeType, setActiveCubeType: handleSetActiveCubeType }}
    >
      {children}
    </CubeSettingsContext.Provider>
  );
};

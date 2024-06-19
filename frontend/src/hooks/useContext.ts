import { useContext } from "react";
import { AuthContext } from "../context/authContext";
import { SettingsContext } from "../context/settingsContext";
import { CubeSettingsContext } from "../context/cubeSettingContext";
import { TimerContext } from "../context/timerContext";
import { LayoutContext } from "../context/layoutContext";

export const useAuth = () => {
  const auth = useContext(AuthContext);
  if (!auth) {
    throw new Error("Can't use auth context outside provider");
  }
  return auth;
};

export const useTimerContext = () => {
  const timerContext = useContext(TimerContext);
  if (!timerContext) {
    throw new Error("Can't use timer context outside provider");
  }
  return timerContext;
};

export const useSettings = () => {
  const settings = useContext(SettingsContext);
  if (!settings) {
    throw new Error("Can't use settings context outside provider");
  }
  return settings;
};
export const useLayoutContext = () => {
  const settings = useContext(LayoutContext);
  if (!settings) {
    throw new Error("Can't use layout context outside provider");
  }
  return settings;
};

export const useCubeSettings = () => {
  const cubeSettings = useContext(CubeSettingsContext);
  if (!cubeSettings) {
    throw new Error("Can't use cube settings context outside provider");
  }
  return cubeSettings;
};

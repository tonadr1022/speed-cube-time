import { useContext } from "react";
import { AuthContext } from "../context/authContext";
import { SettingsContext } from "../context/settingsContext";
import { CubeSettingsContext } from "../context/cubeSettingContext";
import { TimerContext } from "../context/timerContext";
import { LayoutContext } from "../context/layoutContext";
import { OnlineContext } from "../context/onlineContext";

const useCtx = <T>(ctx: React.Context<T | undefined>): T => {
  const con = useContext(ctx);
  if (!con) {
    throw new Error("Can't use context outside provider");
  }
  return con;
};
export const useAuth = () => {
  return useCtx(AuthContext);
};

export const useTimerContext = () => {
  return useCtx(TimerContext);
};

export const useSettings = () => {
  return useCtx(SettingsContext);
};

export const useOnlineContext = () => {
  return useCtx(OnlineContext);
};
export const useLayoutContext = () => {
  return useCtx(LayoutContext);
};

export const useCubeSettings = () => {
  return useCtx(CubeSettingsContext);
};

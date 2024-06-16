import { useContext } from "react";
import { AuthContext } from "../context/authContext";
import { SettingsContext } from "../context/settingsContext";
import { CubeSettingsContext } from "../context/cubeSettingContext";

export const useAuth = () => {
  const auth = useContext(AuthContext);
  if (!auth) {
    throw new Error("Can't use auth context outside provider");
  }
  return auth;
};

export const useSettings = () => {
  const auth = useContext(SettingsContext);
  if (!auth) {
    throw new Error("Can't use settings context outside provider");
  }
  return auth;
};

export const useCubeSettings = () => {
  const auth = useContext(CubeSettingsContext);
  if (!auth) {
    throw new Error("Can't use cube settings context outside provider");
  }
  return auth;
};

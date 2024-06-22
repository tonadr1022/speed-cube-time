import { useEffect } from "react";
import { useSettings } from "./useContext";

export const useTheme = () => {
  const { theme } = useSettings();
  useEffect(() => {
    document.querySelector("html")!.setAttribute("data-theme", theme);
  }, [theme]);
};

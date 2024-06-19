import { useEffect } from "react";
import { useSettings } from "./useContext";

export const useTheme = () => {
  const { theme } = useSettings();
  useEffect(() => {
    console.log(theme, "th");
    document.querySelector("html")!.setAttribute("data-theme", theme);
  }, [theme]);
};

import { useEffect } from "react";

export const useTheme = () => {
  const theme = "dark";
  useEffect(() => {
    document.querySelector("html")!.setAttribute("data-theme", theme);
  }, [theme]);
};

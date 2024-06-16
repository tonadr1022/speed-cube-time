import { useEffect } from "react";

export const useSpaceBarDown = (
  callback: () => void,
  canStart: boolean,
): void => {
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (canStart && e.code === "Space" && !e.repeat) {
        e.preventDefault();
        callback();
      }
    };
    document.addEventListener("keydown", handleKeyDown);

    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [callback, canStart]);
};

export const useSpaceBarUp = (callback: () => void): void => {
  useEffect(() => {
    const handleKeyUp = (e: KeyboardEvent) => {
      if (e.code === "Space" && !e.repeat) {
        e.preventDefault();
        callback();
      }
    };
    document.addEventListener("keyup", handleKeyUp);
    return () => {
      document.removeEventListener("keyup", handleKeyUp);
    };
  }, [callback]);
};

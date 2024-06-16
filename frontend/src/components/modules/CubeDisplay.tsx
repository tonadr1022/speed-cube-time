import { useEffect, useRef } from "react";
import { ScrambleDisplay } from "scramble-display";
import CubeDisplayToggle from "./CubeDisplayToggle";
import { useSettings, useTimerContext } from "../../hooks/useContext";

type Props = {
  elHeight: number | null;
};

const CubeDisplay = ({ elHeight: elHeight }: Props) => {
  const { scramble, cubeType } = useTimerContext();
  const { display3D, setDisplay3D } = useSettings();
  const scrambleRef = useRef<HTMLDivElement>(null);

  type Visualization = "3D" | "2D";

  useEffect(() => {
    if (cubeType && display3D) {
      const el = new ScrambleDisplay();
      el.event = cubeType;
      el.scramble = scramble;
      if (display3D) {
        el.visualization = "3D" as Visualization;
      } else {
        el.visualization = "2D" as Visualization;
      }
      el.style.width = "100%";
      el.style.height = elHeight
        ? `${elHeight}px`
        : `${Math.round(window.innerHeight * 1 * 0.3)}px`;
      scrambleRef.current?.appendChild(el);
      const newref = scrambleRef.current;
      return () => {
        newref?.removeChild(el);
      };
    }
  }, [scramble, cubeType, display3D, elHeight]);
  // if (loading) return <Loading />;
  return (
    <div className="relative">
      <div ref={scrambleRef}></div>
      <div className="absolute top-1 right-2">
        <CubeDisplayToggle setIs3D={setDisplay3D} is3D={display3D} />
      </div>
    </div>
  );
};

export default CubeDisplay;

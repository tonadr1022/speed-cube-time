import { useEffect, useRef } from "react";
import { ScrambleDisplay } from "scramble-display";
import { useSettings, useTimerContext } from "../../hooks/useContext";
import TextToggle from "../common/TextToggle";

type Props = {
  elHeight: number | null;
};

const CubeDisplay = ({ elHeight: elHeight }: Props) => {
  const { scramble, cubeType } = useTimerContext();
  const { display3D, setDisplay3D } = useSettings();
  const scrambleRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const el = new ScrambleDisplay();
    el.event = cubeType;
    el.scramble = scramble;
    if (display3D) {
      el.visualization = "3D";
    } else {
      el.visualization = "2D";
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
  }, [scramble, cubeType, display3D, elHeight]);
  // if (loading) return <Loading />;
  return (
    <div className="h-full relative">
      <div ref={scrambleRef}></div>
      <div className="absolute top-1 right-2">
        <TextToggle
          title1="2D"
          title2="3D"
          name={`cubedisplay${Math.random()}`}
          on={display3D}
          onChange={(e) =>
            setDisplay3D(e.currentTarget.getAttribute("value") === "3D")
          }
        />
      </div>
    </div>
  );
};

export default CubeDisplay;

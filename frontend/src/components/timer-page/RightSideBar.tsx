import { useEffect, useRef, useState } from "react";

import SolvesOverTime from "../modules/SolvesOverTime";
import ModuleSelect from "../modules/ModuleSelect";
import StatsModule from "../modules/StatsModule";
import clsx from "clsx";
import NoSolves from "../common/NoSolves";
import CubeDisplay from "../modules/CubeDisplay";
import { useSettings } from "../../hooks/useContext";
import {
  useFetchCubeSessionSolves,
  useFetchSettings,
} from "../../hooks/useFetch";

const RightSideBar = () => {
  //  const { moduleOne } = useAppSelector((state) => state.setting);
  const containerRef = useRef<HTMLDivElement>(null);
  const { modules, moduleCount } = useSettings();

  const [elHeight, setElHeight] = useState<number>(
    Math.round(window.innerHeight / moduleCount),
  ); // State to hold the height of el
  // Update elHeight state on window resize
  useEffect(() => {
    if (!containerRef.current) return;
    const handleResize = () => {
      setElHeight(containerRef.current!.offsetHeight / moduleCount);
    };
    // Attach the event listener
    window.addEventListener("resize", handleResize);
    // Clean up the event listener on unmount
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, [containerRef, moduleCount]);

  const { data: settings } = useFetchSettings();
  const { data: solves, isLoading: solvesLoading } = useFetchCubeSessionSolves(
    settings?.active_cube_session_id || "failed",
  );

  const moduleIndices = Array.from(
    { length: moduleCount },
    (_, index) => index,
  );

  if (!solves) return <div></div>;
  return (
    <div
      className="h-full hidden md:flex md:flex-col p-2 box-content bg-base-200 w-64"
      ref={containerRef}
    >
      {moduleIndices.map((i) => {
        return (
          <div
            key={i}
            className={clsx(
              "h-1/3 relative group my-2 rounded-lg",
              modules[i] === "solves" && "overflow-y-auto",
            )}
          >
            {(() => {
              if (
                modules[i] !== "cubeDisplay" &&
                modules[i] !== "stats" &&
                solves.length <= 0
              )
                return <NoSolves />;
              switch (modules[i]) {
                case "timeGraph":
                  return <SolvesOverTime elHeight={elHeight} solves={solves} />;
                case "stats":
                  return <StatsModule solves={solves} />;
                case "solves":
                  return (
                    <div className="h-full overflow-y-auto min-h-full">
                      Solves Table
                    </div>
                  );
                case "cubeDisplay":
                  return <CubeDisplay elHeight={elHeight} />;
                default:
                  return null;
              }
            })()}
            <div className="top-0 left-0 z-50 absolute opacity-0 group-hover:opacity-100 transition-opacity duration-200">
              <ModuleSelect moduleNumber={i} />
            </div>
          </div>
        );
      })}
    </div>
    // </Suspense>
  );
};

export default RightSideBar;

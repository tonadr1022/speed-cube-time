import React, { useEffect, useRef, useState } from "react";

import SolvesOverTime from "../modules/SolvesOverTime";
import ModuleSelect from "../modules/ModuleSelect";
import StatsModule from "../modules/StatsModule";
import NoSolves from "../common/NoSolves";
import CubeDisplay from "../modules/CubeDisplay";
import { useSettings } from "../../hooks/useContext";
import SolveTable from "../modules/SolveTable";
import clsx from "clsx";
import { Solve } from "../../types/types";

const SolveTableMemoized = React.memo(SolveTable);
type Props = {
  solves: Solve[];
};

const RightSideBar = ({ solves }: Props) => {
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

  const moduleIndices = Array.from(
    { length: moduleCount },
    (_, index) => index,
  );

  return (
    <div
      className="h-full bg-base-200 hidden md:flex md:flex-col box-content w-64"
      ref={containerRef}
    >
      {moduleIndices.map((i) => {
        return (
          <div
            key={i}
            className={clsx(
              "relative group rounded-lg overflow-y-auto h-full self-stretch",
              `h-1/${moduleCount}`,
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
                  return <SolveTableMemoized solves={solves} />;
                case "cubeDisplay":
                  return <CubeDisplay elHeight={elHeight} />;
                default:
                  return null;
              }
            })()}
            <div className="top-0 left-0 z-50 absolute opacity-0 group-hover:opacity-100">
              <ModuleSelect moduleNumber={i} />
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default RightSideBar;

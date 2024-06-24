import { calculateStats } from "../../data/calculateStats";
import clsx from "clsx";
import { useMemo } from "react";
import { Solve } from "../../types/types";

type StatProps = {
  value: number | null | undefined;
  name: string;
};

const StatModuleStat = ({ name, value }: StatProps) => {
  return (
    <div className="flex relative bg-base-300 flex-1 rounded-lg text-center justify-center">
      <button
        className={clsx(
          name === "best" && "text-success",
          name === "worst" && "text-error",
          "font-bold text-xs",
        )}
      >
        {value ? value?.toFixed(2) : "-"}
      </button>
      <div className="top-0 left-3 absolute text-sm font-medium">{name}</div>
    </div>
  );
};

type Props = {
  solves: Solve[];
  className?: string;
};

const StatsModule = ({ solves, className }: Props) => {
  const stats = useMemo(() => calculateStats(solves), [solves]);
  return (
    <div className={clsx("flex flex-row", className)}>
      <div className="flex flex-col min-h-full w-full gap-1 px-0.5">
        <StatModuleStat name={"best"} value={stats.best} />
        <StatModuleStat name={"avg"} value={stats.currentAvg} />
        <StatModuleStat name={"ao5"} value={stats.currentAo5} />
        <StatModuleStat name={"pb ao5"} value={stats.bestAo5} />
      </div>
      <div className="flex flex-col min-h-full w-full gap-1">
        <StatModuleStat name={"worst"} value={stats.worst} />
        <StatModuleStat name={"median"} value={stats.median} />
        <StatModuleStat name={"ao12"} value={stats.currentAo12} />
        <StatModuleStat name={"pb ao12"} value={stats.bestAo12} />
      </div>
    </div>
  );
};

export default StatsModule;

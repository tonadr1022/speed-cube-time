import { useMemo } from "react";
import { LineChart, Line, YAxis, ResponsiveContainer } from "recharts";
import NoSolves from "../common/NoSolves";
import { Solve } from "../../types/types";

type Props = { solves: Solve[]; elHeight?: number | null };

const SolvesOverTime = ({ solves, elHeight }: Props) => {
  const reversed = useMemo(() => [...solves].reverse(), [solves]);
  return solves.length ? (
    <ResponsiveContainer
      width={"100%"}
      height={elHeight || "100%"}
      className="bg-base-300 rounded-lg"
    >
      <LineChart data={reversed} className="-ml-4 pt-2">
        <Line
          isAnimationActive={false}
          dot={false}
          dataKey="duration"
          stroke="#36d399"
          strokeWidth={5}
        />
        <YAxis
          tickCount={6}
          padding={{ top: 0, bottom: 0 }}
          style={{ margin: 1 }}
          stroke={"#36d399"}
          axisLine={false}
          tickLine={false}
        />
      </LineChart>
    </ResponsiveContainer>
  ) : (
    <NoSolves />
  );
};

export default SolvesOverTime;

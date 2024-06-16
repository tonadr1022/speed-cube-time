import { Solve } from "../types/types";
import sorted from "sorted-array-functions";

export const calculateStats = (
  solves: Solve[],
): {
  best: number | null | undefined;
  worst: number | null | undefined;
  median: number | null | undefined;
  currentAvg: number | null | undefined;
  bestAo5: number | null | undefined;
  bestAo12: number | null | undefined;
  currentAo5: number | null | undefined;
  currentAo12: number | null | undefined;
} => {
  const durations = solves.map((solve) => solve.duration!);
  const sorted = durations.sort((a, b) => a - b);
  const median = sorted[Math.floor(sorted.length / 2)];
  const best = sorted[0];
  const worst = sorted[sorted.length - 1];

  // const recent = solves.slice(0, 5);

  const currentAvg = getAvg(solves);
  const bestAo5 = getAvgPb(solves, 5)?.best;
  const bestAo12 = getAvgPb(solves, 12)?.best;
  const currentAo5 = getCurrentAvg(solves.slice(0, 5), 5);
  const currentAo12 = getCurrentAvg(solves.slice(0, 12), 12);
  return {
    median,
    best,
    worst,
    currentAvg,
    bestAo5,
    bestAo12,
    currentAo5,
    currentAo12,
  };
};

const getSolveDuration = (solve: Solve | number): number => {
  if (typeof solve === "number") return solve;
  else if (solve.dnf) return -1;
  else if (solve.plus_two) return solve.duration + 2;
  else return solve.duration;
};

const getSortedSolveDurations = (solves: Solve[] | number[]) => {
  if (!solves.length) return [];
  const solveDurations = solves.map((solve) => getSolveDuration(solve));
  const sorted = solveDurations.sort((a, b) => a - b);
  return sorted;
};

const getCurrentAvg = (solves: Solve[] | number[], count: number) => {
  if (
    !solves.length ||
    solves.length < 3 ||
    (count > 0 && solves.length < count)
  ) {
    return null;
  }

  const avg = getAvg(solves.slice(0, count));
  return avg;
};
const getAvg = (solves: Solve[] | number[]) => {
  const list = getSortedSolveDurations(solves);

  const count = list.length;
  if (count < 3) return null;
  let dropCount = Math.ceil(Math.max(1, count * 0.05));
  // < 5 is a special case. We don't drop any from either end
  if (count < 5) {
    dropCount = 0;
  }

  // calc dnf count. dnf's are -1 in list
  let dnfCount = 0;
  for (let i = 0; i < list.length; i++) {
    if (list[i] < 0) dnfCount++;
    else break;
  }

  let total = 0;
  let totalCount = 0;
  for (let i = dropCount; i < list.length - dropCount; i++) {
    const duration = list[i + dnfCount];
    if (duration <= 0) return -1;
    else total += duration;
    totalCount++;
  }
  return total / totalCount;
};

export function getAvgPb(solves: Solve[], count: number) {
  // Return null when solve count is less than desired length.
  if (!solves || !solves.length || solves.length < count || count < 3) {
    return null;
  }

  const firstSolves = solves.slice(0, count);
  const sortedDurations = firstSolves
    .map((solve) => solve.duration)
    .sort((a, b) => a - b);

  let bestList: Solve[] = [...firstSolves];
  let best = getAvg(firstSolves);

  for (let i = 1; i < solves.length - count; i++) {
    const dropSolve = solves[i - 1];
    const addSolve = solves[i + count - 1];

    sorted.remove(sortedDurations, dropSolve.duration);
    sorted.add(sortedDurations, addSolve.duration);

    const avg = getAvg(sortedDurations);

    if (avg && best && avg > 0 && (best <= 0 || avg < best)) {
      best = avg;
      bestList = solves.slice(i, i + count);
    }
  }

  return { best, bestList };
}

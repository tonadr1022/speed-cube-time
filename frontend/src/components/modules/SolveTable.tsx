import SolveTableRow from "./SolveTableRow";
import { useCallback } from "react";
import { Solve } from "../../types/types";
import NoSolves from "../common/NoSolves";
import { useDeleteSolve, useUpdateSolve } from "../../hooks/useFetch";
import { useQueryClient } from "@tanstack/react-query";

type Props = { solves: Solve[] };

const SolveTable = ({ solves }: Props) => {
  const client = useQueryClient();
  const deleteSolve = useDeleteSolve(client);
  const updateSolve = useUpdateSolve(client);
  const onSolveDelete = useCallback(
    (solveId: string) => {
      deleteSolve.mutate(solveId);
    },
    [deleteSolve],
  );
  const onTogglePlusTwo = useCallback(
    (solve: Solve) => {
      updateSolve.mutate({
        id: solve.id,
        solve: { plus_two: !solve.plus_two },
      });
    },
    [updateSolve],
  );
  const onToggleDnf = useCallback(
    (solve: Solve) => {
      updateSolve.mutate({
        id: solve.id,
        solve: { dnf: !solve.dnf },
      });
    },
    [updateSolve],
  );

  const length = solves.length ? solves?.length : 0;
  const counts = Array.from(Array(length).keys(), (i) => length - i);
  return solves.length ? (
    <div className="bg-base-200 h-full w-full p-2 flex flex-col rounded-lg">
      {solves.map((solve, i) => (
        <SolveTableRow
          key={i}
          solveCount={counts[i]}
          onDelete={onSolveDelete}
          solve={solve}
          onTogglePlusTwo={onTogglePlusTwo}
          onToggleDnf={onToggleDnf}
        />
      ))}
    </div>
  ) : (
    <NoSolves />
  );
};

export default SolveTable;

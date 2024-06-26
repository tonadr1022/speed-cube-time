import SolveTableRow from "./SolveTableRow";
import { useCallback } from "react";
import { Solve } from "../../types/types";
import { useDeleteSolve, useUpdateSolve } from "../../hooks/useFetch";
import { useQueryClient } from "@tanstack/react-query";
import ReactList from "react-list";
import clsx from "clsx";

type Props = {
  className?: string;
  solves: Solve[];
};

const SolveTable = ({ solves, className }: Props) => {
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

  const renderSolveRow = (index: number, key: number | string) => {
    return (
      <SolveTableRow
        solve={solves[index]}
        onDelete={onSolveDelete}
        onTogglePlusTwo={onTogglePlusTwo}
        onToggleDnf={onToggleDnf}
        solveCount={counts[index]}
        key={key}
      />
    );
  };

  return (
    <div
      className={clsx(
        "bg-base-200 overflow-y-auto w-full flex flex-col rounded-lg",
        className,
      )}
    >
      <ReactList
        itemRenderer={renderSolveRow}
        length={solves.length}
        type="uniform"
      />
    </div>
  );
};

export default SolveTable;

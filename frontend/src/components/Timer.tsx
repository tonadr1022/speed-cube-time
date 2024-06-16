import { createSolve } from "../api/solves-api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Solve } from "../types/types.ts";
import { useState } from "react";

type Props = {};

const Timer = (props: Props) => {
  const [solve, setSolve] = useState<Solve | null>(null);
  const queryClient = useQueryClient();
  const addSolveMutation = useMutation({
    mutationFn: (solve: Solve) => createSolve(solve),
    // TODO use immmer for large arrays
    onMutate: async (newSolve: Solve) => {
      await queryClient.cancelQueries({ queryKey: ["solves"] });
      const prevSolves = queryClient.getQueryData(["solves"]);
      queryClient.setQueryData(["solves"], (old: Solve[]) => [
        ...old,
        newSolve,
      ]);
      return { prevSolves };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(["solves"], context?.prevSolves);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["solves"] });
    },
  });

  const onSolveAdd = () => {
    if (!solve) return;
    addSolveMutation.mutate(solve);
  };

  return <div>Timer</div>;
};

export default Timer;

import { useMutation, useQuery, QueryClient } from "@tanstack/react-query";
import { fetchUserSettings, updateUserSettings } from "../api/settings-api";
import {
  createCubeSession,
  fetchUserCubeSessions,
} from "../api/cube-session-api";
import {
  createSolve,
  deleteSolve,
  fetchUserSolves,
  updateSolve,
} from "../api/solves-api";
import {
  CubeSession,
  CubeSessionCreatePayload,
  Settings,
  SettingsUpdatePayload,
  Solve,
  SolveCreatePayload,
} from "../types/types";
import { toast } from "react-toastify";
import { useAuth, useOnlineContext } from "./useContext";

const toastServerError = () => {
  toast.error("ServerError");
};

export const useFetchCubeSessions = () => {
  const auth = useAuth();
  const { online } = useOnlineContext();
  return useQuery({
    queryKey: ["cubeSessions"],
    queryFn: () => fetchUserCubeSessions(online, auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export const useFetchAllUserSolves = () => {
  const auth = useAuth();
  const { online } = useOnlineContext();
  return useQuery({
    queryKey: ["solves"],
    queryFn: () => fetchUserSolves(online, auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export const useFetchSettings = () => {
  const auth = useAuth();
  const { online } = useOnlineContext();
  return useQuery({
    queryKey: ["settings"],
    queryFn: () => fetchUserSettings(online, auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export type UpdateSettingsArgs = {
  id: string;
  settings: SettingsUpdatePayload;
};

export const useAddSolveMutation = (queryClient: QueryClient) => {
  const { online } = useOnlineContext();
  return useMutation({
    mutationFn: (solve: SolveCreatePayload) => createSolve(online, solve),
    // TODO use immmer for large arrays
    onMutate: async (newSolve: SolveCreatePayload) => {
      await queryClient.cancelQueries({ queryKey: ["solves"] });
      const prevSolves = queryClient.getQueryData(["solves"]);
      queryClient.setQueryData(["solves"], (old: Solve[]) => [
        newSolve,
        ...old,
      ]);
      return { prevSolves };
    },
    onError: (_, __, context) => {
      toastServerError();
      queryClient.setQueryData(["solves"], context?.prevSolves);
    },
  });
};

export const useUpdateSetings = (queryClient: QueryClient) => {
  const { online } = useOnlineContext();
  return useMutation({
    mutationFn: ({ id, settings }: UpdateSettingsArgs) =>
      updateUserSettings(online, id, settings),
    onMutate: (newSettings: UpdateSettingsArgs) => {
      const prevSettings = queryClient.getQueryData(["settings"]);
      queryClient.setQueryData(["settings"], (old: Settings) => ({
        ...old,
        ...newSettings.settings,
      }));
      return { prevSettings };
    },
    onError: toastServerError,
  });
};

export type UpdateSolveArgs = {
  id: string;
  solve: Partial<Solve>;
};

export const useDeleteSolve = (queryClient: QueryClient) => {
  const { online } = useOnlineContext();
  return useMutation({
    mutationFn: (id: string) => deleteSolve(online, id),
    onMutate: (id: string) => {
      const prevSolves: Solve[] | undefined = queryClient.getQueryData([
        "solves",
      ]);
      queryClient.setQueryData(
        ["solves"],
        prevSolves?.filter((solve) => solve.id !== id),
      );
      return { prevSolves };
    },
  });
};

export const useUpdateSolve = (queryClient: QueryClient) => {
  const { online } = useOnlineContext();
  return useMutation({
    mutationFn: ({ id, solve }: UpdateSolveArgs) =>
      updateSolve(online, id, solve),
    onMutate: (newSolve: UpdateSolveArgs) => {
      // await queryClient.cancelQueries(["solves"]);

      const prevSolves = queryClient.getQueryData<Solve[]>(["solves"]);
      const prevSolve = queryClient.getQueryData<Solve>([
        "solves",
        newSolve.id,
      ]);

      // Optimistically update the list of solves
      if (prevSolves) {
        queryClient.setQueryData(
          ["solves"],
          prevSolves.map((solve) =>
            solve.id === newSolve.id ? { ...solve, ...newSolve.solve } : solve,
          ),
        );
      }

      // Optimistically update the specific solve
      queryClient.setQueryData(["solves", newSolve.id], {
        ...prevSolve,
        ...newSolve.solve,
      });

      return { prevSolves, prevSolve };
    },
  });
};

export const useCreateCubeSession = (queryClient: QueryClient) => {
  const { online } = useOnlineContext();
  return useMutation({
    mutationFn: (session: CubeSessionCreatePayload) =>
      createCubeSession(session, online),
    onMutate: async (newSession: CubeSessionCreatePayload) => {
      await queryClient.cancelQueries({ queryKey: ["cubeSessions"] });
      const prevSessions = queryClient.getQueryData(["cubeSessions"]);
      queryClient.setQueryData(["cubeSessions"], (old: CubeSession[]) => [
        ...old,
        newSession,
      ]);
      return { prevSessions };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(["cubeSessions"], context?.prevSessions);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["cubeSessions"] });
    },
  });
};

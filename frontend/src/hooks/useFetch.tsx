import { useMutation, useQuery, QueryClient } from "@tanstack/react-query";
import { fetchUserSettings, updateUserSettings } from "../api/settings-api";
import { fetchUserCubeSessions } from "../api/cube-session-api";
import {
  createSolve,
  fetchCubeSessionSolves,
  fetchUserSolves,
} from "../api/solves-api";
import {
  Settings,
  SettingsUpdatePayload,
  Solve,
  SolveCreatePayload,
} from "../types/types";
import { toast } from "react-toastify";
import { useAuth } from "./useContext";

const toastServerError = () => {
  toast.error("ServerError");
};
export const useFetchCubeSessions = () => {
  const auth = useAuth();
  return useQuery({
    queryKey: ["cubeSessions"],
    queryFn: () => fetchUserCubeSessions(auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export const useFetchAllUserSolves = () => {
  const auth = useAuth();
  return useQuery({
    queryKey: ["solves"],
    queryFn: () => fetchUserSolves(auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export const useFetchSettings = () => {
  const auth = useAuth();
  return useQuery({
    queryKey: ["settings"],
    queryFn: () => fetchUserSettings(auth.user?.id || ""),
    staleTime: 60 * 1000,
  });
};

export type UpdateSettingsArgs = {
  id: string;
  settings: SettingsUpdatePayload;
};

export const useAddSolveMutation = (queryClient: QueryClient) => {
  return useMutation({
    mutationFn: (solve: SolveCreatePayload) => createSolve(true, solve),
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

export const useFetchCubeSessionSolves = (sessionId: string) => {
  return useQuery({
    queryKey: [`solves${sessionId}`],
    queryFn: () => fetchCubeSessionSolves(sessionId),
    staleTime: 60 * 1000,
  });
};

export const useUpdateSetings = (queryClient: QueryClient) => {
  return useMutation({
    mutationFn: ({ id, settings }: UpdateSettingsArgs) =>
      updateUserSettings(id, settings),
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

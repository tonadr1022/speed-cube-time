import { Solve, SolveCreatePayload, SolveUpdatePayload } from "../types/types";
import axiosInstance from "./api";

export const fetchAllSolves = async (): Promise<Solve[]> => {
  const res = await axiosInstance.get("/solves");
  return res.data;
};

export const fetchCubeSessionSolves = async (
  sessionId: string,
): Promise<Solve[]> => {
  const res = await axiosInstance.get(`/sessions/${sessionId}/solves`);
  return res.data;
};

export const fetchUserSolves = async (userId: string): Promise<Solve[]> => {
  const res = await axiosInstance.get(`/users/${userId}/solves`);
  return res.data;
};

export const deleteSolve = async (id: string) => {
  await axiosInstance.delete(`solves/${id}`);
};

export const updateSolve = async (id: string, update: SolveUpdatePayload) => {
  await axiosInstance.patch(`solves/${id}`, update);
};

export const createSolve = async (
  server: boolean,
  solve: SolveCreatePayload,
) => {
  if (server) {
    const res = await axiosInstance.post("/solves", solve);
    return res.data;
  } else {
    try {
      //
    } catch (e) {
      console.log(e);
      return null;
    }
  }
};

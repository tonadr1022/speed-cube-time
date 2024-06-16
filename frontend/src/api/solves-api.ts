import { Solve, SolveCreatePayload } from "../types/types";
import axiosInstance from "./api";

export const fetchAllSolves = async (): Promise<Solve[]> => {
  const res = await axiosInstance.get("/solves");
  return res.data;
};

export const fetchUserSolves = async (userId: string): Promise<Solve[]> => {
  const res = await axiosInstance.get(`/users/${userId}/solves`);
  return res.data;
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

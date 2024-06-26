import { Solve, SolveCreatePayload, SolveUpdatePayload } from "../types/types";
import axiosInstance from "./api";
import {
  fetchLocalSolves,
  createLocalSolve,
  deleteLocalSolve,
  updateLocalSolve,
} from "../browser_storage/indexedDB";

// export const fetchAllSolves = async (): Promise<Solve[]> => {
//   const res = await fetchLocalSolves();
//   return res;
//   // const res = await axiosInstance.get("/solves");
//   // return res.data;
// };

// export const fetchCubeSessionSolves = async (
//   sessionId: string,
// ): Promise<Solve[]> => {
//   const res = await axiosInstance.get(`/sessions/${sessionId}/solves`);
//   return res.data;
// };

export const fetchUserSolves = async (
  server: boolean,
  userId: string,
): Promise<Solve[]> => {
  if (server) {
    const res = await axiosInstance.get(`/users/${userId}/solves`);
    return res.data;
  } else {
    const res = await fetchLocalSolves();
    return res;
  }
};

export const deleteSolve = async (server: boolean, id: string) => {
  if (server) {
    await axiosInstance.delete(`solves/${id}`);
  } else {
    deleteLocalSolve(id);
  }
};

export const updateSolve = async (
  server: boolean,
  id: string,
  update: SolveUpdatePayload,
) => {
  if (server) {
    await axiosInstance.patch(`solves/${id}`, update);
  } else {
    await updateLocalSolve(id, update);
  }
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
      const res = await createLocalSolve(solve);
      return res;
    } catch (e) {
      console.log("localdb error", e);
      throw e;
    }
  }
};

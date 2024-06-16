import { CubeSession, CubeSessionCreatePayload } from "../types/types";
import axiosInstance from "./api";

export const fetchAllCubeSessions = async () => {
  const res = await axiosInstance.get("/sessions");
  return res.data;
};

export const fetchUserCubeSessions = async (
  userId: string,
): Promise<CubeSession[]> => {
  const res = await axiosInstance.get(`/users/${userId}/sessions`);
  return res.data;
};

export const createCubeSession = async (
  cubeSession: CubeSessionCreatePayload,
  server: boolean,
) => {
  const res = await axiosInstance.post("/sessions", {
    name: cubeSession.name,
    cube_type: cubeSession.cubeType,
  });
  return res.data;
};

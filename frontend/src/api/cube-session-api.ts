import {
  createLocalCubeSession,
  fetchLocalCubeSessions,
} from "../browser_storage/indexedDB";
import { CubeSession, CubeSessionCreatePayload } from "../types/types";
import axiosInstance from "./api";

export const fetchAllCubeSessions = async () => {
  const res = await axiosInstance.get("/sessions");
  return res.data;
};

export const fetchUserCubeSessions = async (
  server: boolean,
  userId: string,
): Promise<CubeSession[]> => {
  if (server) {
    const res = await axiosInstance.get(`/users/${userId}/sessions`);
    return res.data;
  } else {
    const res = await fetchLocalCubeSessions();
    if (!res.length) {
      await createLocalCubeSession({
        name: "Default",
        cube_type: "333",
      });
      return fetchUserCubeSessions(false, "");
    }
    return res;
  }
};

export const createCubeSession = async (
  cubeSession: CubeSessionCreatePayload,
  server: boolean,
) => {
  if (server) {
    const res = await axiosInstance.post("/sessions", {
      name: cubeSession.name,
      cube_type: cubeSession.cube_type,
    });
    return res.data;
  } else {
    const res = await createLocalCubeSession(cubeSession);
    console.log(res);
    return res;
  }
};

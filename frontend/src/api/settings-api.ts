import {
  fetchLocalCubeSessions,
  fetchLocalSettings,
  updateLocalSettings,
} from "../browser_storage/indexedDB";
import { Settings } from "../types/types";
import axiosInstance from "./api";
import { v4 as uuid } from "uuid";

export const fetchUserSettings = async (
  server: boolean,
  userId: string,
): Promise<Settings> => {
  if (server) {
    const res = await axiosInstance.get(`/users/${userId}/settings`);
    return res.data;
  } else {
    const res = await fetchLocalSettings();
    if (!res) {
      const cubeSessions = await fetchLocalCubeSessions();
      const initialSettings: Settings = {
        id: uuid(),
        theme: "dark",
        active_cube_session_id: cubeSessions[0].id,
        created_at: new Date(),
        updated_at: new Date(),
      };
      await updateLocalSettings(initialSettings);
      console.log("mc");
      return new Promise<Settings>(() => {
        return initialSettings;
      });
    }
    return res;
  }
};

export const updateUserSettings = async (
  server: boolean,
  id: string,
  settings: Partial<Settings>,
) => {
  if (server) {
    const res = await axiosInstance.patch(`/settings/${id}`, settings);
    return res.data;
  } else {
    await updateLocalSettings(settings);
  }
};

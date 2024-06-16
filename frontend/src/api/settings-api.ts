import { Settings, SettingsUpdatePayload } from "../types/types";
import axiosInstance from "./api";

export const fetchUserSettings = async (userId: string): Promise<Settings> => {
  const res = await axiosInstance.get(`/users/${userId}/settings`);
  console.log(res.data);
  return res.data;
};

export const updateUserSettings = async (
  id: string,
  settings: SettingsUpdatePayload,
) => {
  const res = await axiosInstance.patch(`/settings/${id}`, settings);
  return res.data;
};

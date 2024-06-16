import axiosInstance from "./api";

export type RegisterResponse = {
  token: string;
};

export const registerUser = async (
  username: string,
  password: string,
): Promise<RegisterResponse> => {
  const res = await axiosInstance.post("/register", { username, password });
  return res.data;
};

export type LoginResponse = {
  token: string;
};

export const loginUser = async (
  username: string,
  password: string,
): Promise<LoginResponse> => {
  const res = await axiosInstance.post("/login", { username, password });
  return res.data;
};

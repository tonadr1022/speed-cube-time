import axios from "axios";

export const baseURL = import.meta.env.VITE_BASE_URL;
const axiosInstance = axios.create({
  baseURL: baseURL,
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
    Authorization: localStorage.getItem("token")
      ? `Bearer ${localStorage.getItem("token")}`
      : null,
  },
});

export default axiosInstance;

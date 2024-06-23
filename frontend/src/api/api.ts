import axios from "axios";

// const baseURL = "http://localhost";
const baseURL = "https://speed-cube-time-e63478004c61.herokuapp.com";
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

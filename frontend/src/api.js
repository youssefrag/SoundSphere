import axios from "axios";

export default axios.create({
  // baseURL: "http://localhost:8080",
  baseURL: import.meta.env.VITE_APP_API_URL,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

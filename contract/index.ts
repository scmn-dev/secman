const axios = require("axios");
import { API_URL } from "../constants";

export const API = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json; charset=utf-8",
    Accept: "application/json, text/plain, */*",
  },
  withCredentials: true,
});

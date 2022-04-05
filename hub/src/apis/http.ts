import Axios from "axios";

export let baseURL = "https://api.secman.dev";

const client = Axios.create({
  baseURL,
  headers: {
    "Content-Type": "application/json; charset=utf-8",
    Accept: "application/json, text/plain, */*",
  },
  withCredentials: true,
});

export default class HTTPClient {
  static async head(path: any) {
    return client.head(path);
  }

  static async get(path: any, params = {}, headers = {}) {
    return client.get(path, {
      params,
      headers,
    });
  }

  static async post(path: any, data = {}, headers = {}, onUploadProgress: any) {
    return client.post(path, data, {
      headers,
      onUploadProgress,
    });
  }

  static async put(path: any, data = {}, headers = {}) {
    return client.put(path, data, {
      headers,
    });
  }

  static async delete(path: any, data = {}, headers = {}) {
    return client.delete(path, {
      data,
      headers,
    });
  }

  static setHeader(key: any, value: any) {
    client.defaults.headers.common[key] = value;
  }

  static setBaseURL() {
    client.defaults.baseURL = baseURL;
  }
}

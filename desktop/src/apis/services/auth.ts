import HTTPClient from "@/apis/http";

export default class AuthService {
  static async Login(payload: any) {
    return HTTPClient.post(`/auth/signin`, payload, {}, {});
  }

  static async Logout(payload: any) {
    return HTTPClient.post(`/auth/signout`, payload, {}, {});
  }

  static async Check(payload: any) {
    return HTTPClient.post("/auth/check", payload, {}, {});
  }

  static async Refresh(payload: any) {
    return HTTPClient.post(`/auth/refresh`, payload, {}, {});
  }
}

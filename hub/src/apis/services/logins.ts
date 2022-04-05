import HTTPClient from "@/apis/http";

export default class LoginsService {
  static async FetchAll(query: any) {
    return HTTPClient.get(`/api/logins`, query);
  }

  static async Get(id: any) {
    return HTTPClient.get(`/api/logins/${id}`);
  }

  static async Create(payload: any) {
    return HTTPClient.post(`/api/logins`, payload, {}, {});
  }

  static async Update(id: any, payload: any) {
    return HTTPClient.put(`/api/logins/${id}`, payload);
  }

  static async Delete(id: any) {
    return HTTPClient.delete(`/api/logins/${id}`);
  }
}

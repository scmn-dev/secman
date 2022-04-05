import HTTPClient from "@/apis/http";

export default class ServersService {
  static async FetchAll(query: any) {
    return HTTPClient.get(`/api/servers`, query);
  }

  static async Get(id: any) {
    return HTTPClient.get(`/api/servers/${id}`);
  }

  static async Create(payload: any) {
    return HTTPClient.post(`/api/servers`, payload, {}, {});
  }

  static async Update(id: any, payload: any) {
    return HTTPClient.put(`/api/servers/${id}`, payload);
  }

  static async Delete(id: any) {
    return HTTPClient.delete(`/api/servers/${id}`);
  }
}

import HTTPClient from "@/apis/http";

export default class EmailsService {
  static async FetchAll(query: any) {
    return HTTPClient.get(`/api/emails`, query);
  }

  static async Get(id: any) {
    return HTTPClient.get(`/api/emails/${id}`);
  }

  static async Create(payload: any) {
    return HTTPClient.post(`/api/emails`, payload, {}, {});
  }

  static async Update(id: any, payload: any) {
    return HTTPClient.put(`/api/emails/${id}`, payload);
  }

  static async Delete(id: any) {
    return HTTPClient.delete(`/api/emails/${id}`);
  }
}

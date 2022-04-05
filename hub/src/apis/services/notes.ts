import HTTPClient from "@/apis/http";

export default class NotesService {
  static async FetchAll(query: any) {
    return HTTPClient.get(`/api/notes`, query);
  }

  static async Get(id: any) {
    return HTTPClient.get(`/api/notes/${id}`);
  }

  static async Create(payload: any) {
    return HTTPClient.post(`/api/notes`, payload, {}, {});
  }

  static async Update(id: any, payload: any) {
    return HTTPClient.put(`/api/notes/${id}`, payload);
  }

  static async Delete(id: any) {
    return HTTPClient.delete(`/api/notes/${id}`);
  }
}

import HTTPClient from "@/apis/http";

export default class CreditCardsService {
  static async FetchAll(query: any) {
    return HTTPClient.get(`/api/credit-cards`, query);
  }

  static async Get(id: any) {
    return HTTPClient.get(`/api/credit-cards/${id}`);
  }

  static async Create(payload: any) {
    return HTTPClient.post(`/api/credit-cards`, payload, {}, {});
  }

  static async Update(id: any, payload: any) {
    return HTTPClient.put(`/api/credit-cards/${id}`, payload);
  }

  static async Delete(id: any) {
    return HTTPClient.delete(`/api/credit-cards/${id}`);
  }
}

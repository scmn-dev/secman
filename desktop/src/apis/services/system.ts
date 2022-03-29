import HTTPClient from "@/apis/http";

export default class SystemService {
  static async GeneratePassword() {
    return HTTPClient.post(`/api/system/generate-password`, "", {}, {});
  }

  static async Import(data: any) {
    return HTTPClient.post(`/api/system/import`, data, {}, {});
  }

  static async Export(data: any) {
    return HTTPClient.post(`/api/system/export`, data, {}, {});
  }
}

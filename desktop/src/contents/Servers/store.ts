import ServersService from "@/apis/services/servers";
import { CryptoTools } from "@secman/crypto";

const EncryptedFields = [
  "ip",
  "username",
  "password",
  "hosting_username",
  "hosting_password",
  "admin_username",
  "admin_password",
  "extra",
];

export default {
  namespaced: true,

  state() {
    return {
      ItemList: [],
      Detail: {},
    };
  },

  actions: {
    async FetchAll({ state }: any, query: any) {
      const { data } = await ServersService.FetchAll(query);

      const itemList = JSON.parse(CryptoTools.aesDecrypt(data.data));

      itemList.forEach((element: any) => {
        CryptoTools.decryptFields(element, EncryptedFields);
      });

      state.ItemList = itemList;
    },

    Delete(_: any, id: any) {
      return ServersService.Delete(id);
    },

    Create(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return ServersService.Create(payload);
    },

    Update(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return ServersService.Update(data.id, payload);
    },
  },
};

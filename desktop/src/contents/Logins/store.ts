import LoginsService from "@/apis/services/logins";
import { CryptoTools } from "@secman/crypto";

const EncryptedFields = ["username", "password", "extra"];

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
      const { data } = await LoginsService.FetchAll(query);

      const itemList = JSON.parse(CryptoTools.aesDecrypt(data.data));

      itemList.forEach((element: any) => {
        CryptoTools.decryptFields(element, EncryptedFields);
      });

      state.ItemList = itemList;
    },

    Delete(_: any, id: any) {
      return LoginsService.Delete(id);
    },

    Create(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return LoginsService.Create(payload);
    },

    Update(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return LoginsService.Update(data.id, payload);
    },
  },
};

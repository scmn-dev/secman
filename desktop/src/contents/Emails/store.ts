import EmailsService from "@/apis/services/emails";
import { CryptoTools } from "@secman/crypto";

const EncryptedFields = ["email", "password"];

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
      const { data } = await EmailsService.FetchAll(query);

      const itemList = JSON.parse(CryptoTools.aesDecrypt(data.data));

      itemList.forEach((element: any) => {
        CryptoTools.decryptFields(element, EncryptedFields);
      });

      state.ItemList = itemList;
    },

    Delete(_: any, id: any) {
      return EmailsService.Delete(id);
    },

    Create(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return EmailsService.Create(payload);
    },

    Update(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return EmailsService.Update(data.id, payload);
    },
  },
};

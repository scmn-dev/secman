import CreditCardsService from "@/apis/services/cc";
import { CryptoTools } from "@secman/crypto";

const EncryptedFields = [
  "type",
  "number",
  "expiry_date",
  "cardholder_name",
  "verification_number",
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
      const { data } = await CreditCardsService.FetchAll(query);

      const itemList = JSON.parse(CryptoTools.aesDecrypt(data.data));

      itemList.forEach((element: any) => {
        CryptoTools.decryptFields(element, EncryptedFields);
      });

      state.ItemList = itemList;
    },

    Delete(_: any, id: any) {
      return CreditCardsService.Delete(id);
    },

    Create(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return CreditCardsService.Create(payload);
    },

    Update(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return CreditCardsService.Update(data.id, payload);
    },
  },
};

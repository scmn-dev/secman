import NotesService from "@/apis/services/notes";
import { CryptoTools } from "@secman/crypto";

const EncryptedFields = ["note"];

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
      const { data } = await NotesService.FetchAll(query);

      const itemList = JSON.parse(CryptoTools.aesDecrypt(data.data));

      itemList.forEach((element: any) => {
        CryptoTools.decryptFields(element, EncryptedFields);
      });

      state.ItemList = itemList;
    },

    Delete(_: any, id: any) {
      return NotesService.Delete(id);
    },

    Create(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return NotesService.Create(payload);
    },

    Update(_: any, data: any) {
      const payload = CryptoTools.encryptPayload(data, EncryptedFields);
      return NotesService.Update(data.id, payload);
    },
  },
};

import Vue from "vue";
import Vuex from "vuex";
import { CryptoTools } from "@secman/crypto";
import AuthService from "@/apis/services/auth";
import Storage from "@/tools/storage";
import Logins from "@/contents/Logins/store";
import CreditCards from "@/contents/CreditCards/store";
import Emails from "@/contents/Emails/store";
import Notes from "@/contents/Notes/store";
import Servers from "@/contents/Servers/store";
import HTTPClient, { baseURL } from "@/apis/http";

Vue.use(Vuex);

export default new Vuex.Store({
  state() {
    return {
      access_token: "",
      refresh_token: "",
      transmission_key: "",
      master_hash: "",
      searchQuery: "",

      user: {},
    };
  },
  getters: {},

  actions: {
    async init({ state }: any) {
      CryptoTools.encryptKey = await Storage.getItem("master_hash");
      CryptoTools.transmissionKey = await Storage.getItem("transmission_key");

      state.access_token = await Storage.getItem("access_token");
      state.refresh_token = await Storage.getItem("refresh_token");
      state.transmission_key = await Storage.getItem("transmission_key");
      state.master_hash = await Storage.getItem("master_hash");
      state.user = await Storage.getItem("user");
    },

    async RefreshToken({ state }: any, payload) {
      var token = await Storage.getItem("refresh_token");
      const { data } = await AuthService.Refresh({ refresh_token: token });

      state.access_token = data.access_token;
      state.refresh_token = data.refresh_token;
      state.transmission_key = data.transmission_key.substr(0, 32);
      CryptoTools.transmissionKey = state.transmission_key;

      await Promise.all([
        Storage.setItem("access_token", data.access_token),
        Storage.setItem("refresh_token", data.refresh_token),
        Storage.setItem("transmission_key", state.transmission_key),
      ]);

      HTTPClient.setHeader("Authorization", `Bearer ${state.access_token}`);
    },

    async Login({ state }: any, payload) {
      payload.master_password = CryptoTools.sha256Encrypt(
        payload.master_password
      );

      const { data } = await AuthService.Login(payload);

      state.access_token = data.access_token;
      state.refresh_token = data.refresh_token;
      state.transmission_key = data.transmission_key.substr(0, 32);
      state.master_hash = CryptoTools.pbkdf2Encrypt(
        data.secret,
        payload.master_password
      );
      CryptoTools.encryptKey = state.master_hash;
      CryptoTools.transmissionKey = state.transmission_key;
      state.user = data;

      await Promise.all([
        Storage.setItem("email", payload.email),
        Storage.setItem("server", baseURL),
        Storage.setItem("access_token", data.access_token),
        Storage.setItem("refresh_token", data.refresh_token),
        Storage.setItem("user", data),
        Storage.setItem("master_hash", state.master_hash),
        Storage.setItem("transmission_key", state.transmission_key),
      ]);

      HTTPClient.setHeader("Authorization", `Bearer ${state.access_token}`);
    },

    async Logout({ state }: any) {
      const email = await Storage.getItem("email");
      const server = baseURL;
      await Storage.clear();

      state.access_token = null;
      state.refresh_token = null;
      state.transmission_key = null;
      state.master_hash = null;
      state.user = null;

      await Storage.setItem("email", email);
      await Storage.setItem("server", server);
    },
    async loadStore({ state }: any) {
      state.user = await Storage.getItem("user");
    },
  },
  mutations: {
    onInputSearchQuery(state: any, event) {
      state.searchQuery = event.target.value;
    },
  },

  modules: {
    Logins,
    CreditCards,
    Emails,
    Notes,
    Servers,
  },
});

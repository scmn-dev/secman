import Vue from "vue";

import "./styles/app.scss";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import i18n from "./i18n";
import "./registerServiceWorker";
import "./config";
import { waiter } from "./config";

store.dispatch("Logout");

Vue.config.productionTip = false;
Vue.config.silent = true;

new Vue({
  router,
  store,
  i18n,
  wait: waiter,
  render: (h) => h(App),
}).$mount("#app");

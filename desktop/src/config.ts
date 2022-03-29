import Vue from "vue";
import i18n from "@/i18n";
import store from "@/store";
import router from "@/router";
import * as Waiters from "@/tools/waiters";
import storage from "@/tools/storage";
import * as Contant from "@/tools/constants";
import VueClipboard from "vue-clipboard2";
import VuePerfectScrollbar from "vue2-perfect-scrollbar";
import "vue2-perfect-scrollbar/dist/vue2-perfect-scrollbar.css";
import VeeValidate from "vee-validate";
import VueDebounce from "vue-debounce";
import VTooltip from "v-tooltip";
import VueWait from "vue-wait";
import Notifications from "vue-notification";

Vue.prototype.$waiters = Waiters;
Vue.prototype.$storage = storage;
Vue.prototype.$C = Contant;
Vue.use(VueClipboard);
Vue.use(VuePerfectScrollbar);
Vue.use(VeeValidate, { events: "input|blur" });
Vue.use(VueDebounce, { defaultTime: "700ms" });
Vue.use(VTooltip);
Vue.use(VueWait);
Vue.use(Notifications, { duration: 2500 });

export const waiter = new VueWait({
  registerComponent: false,
  registerDirective: false,
});

const requireComponent = require.context("./components", true, /\.(vue)$/);
requireComponent.keys().forEach((fileName) => {
  const componentConfig = requireComponent(fileName);
  Vue.component(componentConfig.default.name, componentConfig.default);
});

Vue.prototype.$notifyError = (text: any) =>
  Vue.prototype.$notify({ type: "error", text });
Vue.prototype.$notifyWarn = (text: any) =>
  Vue.prototype.$notify({ type: "warn", text });
Vue.prototype.$notifySuccess = (text: any) =>
  Vue.prototype.$notify({ type: "success", text });

Vue.prototype.$request = async (
  callback: any,
  waitKey: any,
  errorCallback: any = null
) => {
  try {
    waiter.start(waitKey);
    await callback();
  } catch (error: any) {
    console.log(error);

    if (error.response) {
      if (error.response.status === 401 && !router.app.$route.meta?.auth) {
        store.dispatch("Logout");
        return router.push({ name: "Login" });
      }

      if (errorCallback) {
        errorCallback(error);
      } else if (error.response.status >= 500) {
        Vue.prototype.$notifyError(i18n.t("API500ErrorMessage"));
      } else if (error.response.data.Message && error.response.status != 401) {
        Vue.prototype.$notifyError(error.response.data.Message);
      }
    } else {
      Vue.prototype.$notifyError("Network Error !");
    }
  } finally {
    waiter.end(waitKey);
  }
};

import store from "@/store";

export default () => {
  store.commit("onInputSearchQuery", { target: { value: "" } });
};

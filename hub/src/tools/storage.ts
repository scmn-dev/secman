import LocalForage from "localforage";

export default LocalForage.createInstance({
  driver: LocalForage.INDEXEDDB,
  name: "SMD Storage",
  storeName: "login_data",
});

import Storage from "@/tools/storage";

export default async (to: any, _: any, next: any) => {
  const isAuthPage = to.matched.some((record: any) => record.meta.auth);

  const user = await Storage.getItem("user");

  if (user !== null) {
    if (isAuthPage) {
      return next({ name: "Home" });
    }
  } else {
    if (!isAuthPage) {
      return next({ name: "Login" });
    }
  }

  next();
};

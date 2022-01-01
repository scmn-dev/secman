import axios from "axios";

export const GetLatest = async (repo: string) => {
  let url = "https://api.secman.dev/latest";

  if (repo === "core") {
    url = "https://api.secman.dev/latest-core";
  }

  try {
    const res = await axios.get(url);
    return res.data;
  } catch (error) {
    console.error(error);
  }
};

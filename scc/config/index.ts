import { DOT_SECMAN_PATH, SECMAN_CONFIG_PATH } from "../constants";
import { writeJsonFile as writeJSON } from "../tools/json/write";
import fs from "fs";
import path from "path";

export default async function writeConfigFile(
  username: any,
  user_email: any,
  access_token: any,
  refresh_token: any,
  transmission_key: any,
  master_password_hash: any,
  secret: any
) {
  if (!fs.existsSync(DOT_SECMAN_PATH)) {
    fileIsNotFound();
  }

  // write config file
  await writeJSON(
    SECMAN_CONFIG_PATH,
    {
      config: {
        name: username,
        secret: secret,
        user: user_email,
      },
      data: {
        access_token: access_token,
        master_password_hash: master_password_hash,
        refresh_token: refresh_token,
        transmission_key: transmission_key,
      },
    },
    {}
  );
}

export function readConfig(obj: any) {
  if (!fs.existsSync(SECMAN_CONFIG_PATH)) {
    fileIsNotFound();
  } else {
    let rawData: any = fs.readFileSync(path.resolve(SECMAN_CONFIG_PATH));

    let data: any = JSON.parse(rawData)["config"][obj];

    return data;
  }
}

export function readData(obj: any) {
  if (!fs.existsSync(SECMAN_CONFIG_PATH)) {
    fileIsNotFound();
  } else {
    let rawData: any = fs.readFileSync(path.resolve(SECMAN_CONFIG_PATH));

    let data: any = JSON.parse(rawData)["data"][obj];

    return data;
  }
}

const fileIsNotFound = () => {
  console.log("~/.secman/secman.json does not exist, run `secman init`");

  process.exit(0);
};

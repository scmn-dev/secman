import sh from "shelljs";
const powershell = require("powershell");
import {
  DOT_SECMAN_PATH,
  SECMAN_CONFIG_PATH,
  SECMAN_DATA_PATH,
  SECMAN_SETTINGS_PATH
} from "../../constants";
import { writeJsonFile as writeJSON } from "../../tools/json/write";
import { platform } from "os";
import fs from "fs";
import path from "path";
import { command, error, warning } from "../../design/layout";
import { logSymbols } from "../../design/log";

export default async function writeConfigFile(
  username: any,
  user_email: any,
  access_token: any,
  refresh_token: any,
  transmission_key: any,
  master_password_hash: any,
  secret: any
) {
  if (platform() === "win32") {
    if (!fs.existsSync(DOT_SECMAN_PATH)) {
      const ps = new powershell(`
        New-Item -ItemType "directory" -Path "${DOT_SECMAN_PATH}"
        New-Item ${SECMAN_CONFIG_PATH}
        New-Item ${SECMAN_DATA_PATH}
      `);

      ps.on("output", (data: any) => {
        console.log(data);
      });
    }
  } else {
    if (!fs.existsSync(DOT_SECMAN_PATH)) {
      fs.mkdirSync(DOT_SECMAN_PATH, { recursive: true });
    }

    if (!fs.existsSync(SECMAN_CONFIG_PATH)) {
      sh.touch(SECMAN_CONFIG_PATH);
      writeCFile();
    }

    if (!fs.existsSync(SECMAN_DATA_PATH)) {
      sh.touch(SECMAN_DATA_PATH);
      writeDFile();
    }
  }

  // write config file
  await writeJSON(
    SECMAN_CONFIG_PATH,
    { name: username, user: user_email, secret: secret },
    {}
  );

  writeDataFile(
    access_token,
    refresh_token,
    transmission_key,
    master_password_hash
  );
}

export async function writeDataFile(
  access_token: any,
  refresh_token: any,
  transmission_key: any,
  master_password_hash: any
) {
  await writeJSON(
    SECMAN_DATA_PATH,
    {
      data: {
        access_token: access_token,
        refresh_token: refresh_token,
        transmission_key: transmission_key,
        master_password_hash: master_password_hash
      }
    },
    {}
  );
}

export async function writeCFile() {
  await writeJSON(SECMAN_CONFIG_PATH, {}, {});
}

export async function writeDFile() {
  await writeJSON(SECMAN_DATA_PATH, {}, {});
}

export async function writeSettingFile() {
  await writeJSON(
    SECMAN_SETTINGS_PATH,
    {
      editor: "secman_editor"
    },
    {}
  );
}

export function readConfigFile(obj: any) {
  if (!fs.existsSync(SECMAN_CONFIG_PATH)) {
    fileIsNotFound("config");
  } else {
    let rawData: any = fs.readFileSync(path.resolve(SECMAN_CONFIG_PATH));

    let data: any = JSON.parse(rawData)[obj];

    return data;
  }
}

export function readDataFile(obj: any) {
  if (!fs.existsSync(SECMAN_DATA_PATH)) {
    fileIsNotFound("data");
  } else {
    try {
      let rawData: any = fs.readFileSync(path.resolve(SECMAN_DATA_PATH));

      let data: any = JSON.parse(rawData).data[obj];

      return data;
    } catch {
      console.log(
        error(
          `can't find your auth tokens, to authenticate run ${command(
            "`secman auth`",
            true
          )}.`
        )
      );
      process.exit(0);
    }
  }
}

export function readSettingsFile(obj: any) {
  if (!fs.existsSync(SECMAN_SETTINGS_PATH)) {
    fileIsNotFound("settings");
  } else {
    let rawData: any = fs.readFileSync(path.resolve(SECMAN_SETTINGS_PATH));

    let data: any = JSON.parse(rawData)[obj];

    return data;
  }
}

const fileIsNotFound = (fileName: string) => {
  console.log(
    logSymbols.warning +
      warning(" ~/.secman/" + fileName + ".json ") +
      `does not exist, run ${command("secman init")}`
  );

  process.exit(0);
};

import * as sh from "shelljs";
const powershell = require("powershell");

import {
  DOT_SECMAN_PATH,
  SECMAN_CONFIG_PATH,
  SECMAN_DATA_PATH,
  SECMAN_SETTINGS_PATH,
} from "../../constants";
import { writeJsonFile as writeJSON } from "../../tools/json/write";
import { homedir, platform } from "os";
import fs from "fs";
import path from "path";
import chalk from "chalk";

const secman_dir = path.join(homedir(), DOT_SECMAN_PATH);
const sm_config = path.join(homedir(), SECMAN_CONFIG_PATH);
const sm_data = path.join(homedir(), SECMAN_DATA_PATH);
const sm_setting = path.join(homedir(), SECMAN_SETTINGS_PATH);

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
    if (!fs.existsSync(secman_dir)) {
      const ps = new powershell(`
        New-Item -ItemType "directory" -Path "${secman_dir}"
        New-Item ${sm_config}
        New-Item ${sm_data}
      `);

      ps.on("output", (data: any) => {
        console.log(data);
      });
    }
  } else {
    // sh.mkdir(secman_dir);
    if (!fs.existsSync(secman_dir)) {
      fs.mkdirSync(secman_dir, { recursive: true });
    }

    if (!fs.existsSync(sm_config)) {
      sh.touch(sm_config);
    }

    if (!fs.existsSync(sm_data)) {
      sh.touch(sm_data);
    }
  }

  // write config file
  await writeJSON(
    sm_config,
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
    sm_data,
    {
      data: {
        access_token: access_token,
        refresh_token: refresh_token,
        transmission_key: transmission_key,
        master_password_hash: master_password_hash,
      },
    },
    {}
  );
}

export async function writeSettingFile() {
  await writeJSON(
    sm_setting,
    {
      read_output: "table",
      disable_version_check: false,
      editor: "secman_editor",
    },
    {}
  );
}

export function readConfigFile(obj: any) {
  let rawdata: any = fs.readFileSync(path.resolve(sm_config));

  let data: any = JSON.parse(rawdata)[obj];

  return data;
}

export function readDataFile(obj: any) {
  let rawdata: any = fs.readFileSync(path.resolve(sm_data));

  let data: any = JSON.parse(rawdata).data[obj];

  return data;
}

export function readSettingsFile(obj: any) {
  // check if file exists
  if (!fs.existsSync(sm_setting)) {
    chalk.yellow.bold(
      `~/.secman/settings.json does not exist, run ${chalk.grey.bold(
        "secman init"
      )}.`
    );
    return;
  } else {
    let rawdata: any = fs.readFileSync(path.resolve(sm_setting));

    let data: any = JSON.parse(rawdata)[obj];

    return data;
  }
}

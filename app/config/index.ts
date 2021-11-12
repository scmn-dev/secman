import sh from "shelljs";
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
import process from "process";

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
    if (!fs.existsSync(secman_dir)) {
      fs.mkdirSync(secman_dir, { recursive: true });
    }

    if (!fs.existsSync(sm_config)) {
      sh.touch(sm_config);
      writeCFile();
    }

    if (!fs.existsSync(sm_data)) {
      sh.touch(sm_data);
      writeDFile();
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

export async function writeCFile() {
  await writeJSON(sm_config, {}, {});
}

export async function writeDFile() {
  await writeJSON(sm_data, {}, {});
}

export async function writeSettingFile() {
  await writeJSON(
    sm_setting,
    {
      editor: "secman_editor",
    },
    {}
  );
}

export function readConfigFile(obj: any) {
  if (!fs.existsSync(sm_config)) {
    fileIsNotFound("config");
  } else {
    let rawData: any = fs.readFileSync(path.resolve(sm_config));

    let data: any = JSON.parse(rawData)[obj];

    return data;
  }
}

export function readDataFile(obj: any) {
  if (!fs.existsSync(sm_data)) {
    fileIsNotFound("data");
  } else {
    try {
      let rawData: any = fs.readFileSync(path.resolve(sm_data));

      let data: any = JSON.parse(rawData).data[obj];

      return data;
    } catch {
      console.log(
        chalk.red(
          `can't find your auth tokens, to authenticate run ${chalk.grey.bold(
            "`secman auth`"
          )}.`
        )
      );
      process.exit(0);
    }
  }
}

export function readSettingsFile(obj: any) {
  if (!fs.existsSync(sm_setting)) {
    fileIsNotFound("settings");
  } else {
    let rawData: any = fs.readFileSync(path.resolve(sm_setting));

    let data: any = JSON.parse(rawData)[obj];

    return data;
  }
}

const fileIsNotFound = (fileName: string) => {
  console.log(
    `${chalk.yellow.bold(
      "~/.secman/" + fileName + ".json"
    )} does not exist, run ${chalk.grey.bold("secman init")}.`
  );

  process.exit(0);
};

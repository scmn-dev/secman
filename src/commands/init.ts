import { Command, flags } from "@oclif/command";
const chalk = require("chalk");
import * as sh from "shelljs";
const powershell = require("powershell");
import { spinner } from "@secman/spinner";
import { homedir, platform } from "os";
import fs from "fs";
import path from "path";
import {
  DOT_SECMAN_PATH,
  SECMAN_CONFIG_PATH,
  SECMAN_DATA_PATH,
  SECMAN_EDITOR_PATH,
  SECMAN_SETTINGS_PATH,
  SECMAN_SETTINGS_URL,
} from "../../constants";
import { InstallEditor } from "../../tools/install-editor";
import { writeSettingFile } from "../../app/config";

const secman_dir = path.join(homedir(), DOT_SECMAN_PATH);
const sm_config = path.join(homedir(), SECMAN_CONFIG_PATH);
const sm_data = path.join(homedir(), SECMAN_DATA_PATH);
const sm_setting = path.join(homedir(), SECMAN_SETTINGS_PATH);
const sm_editor = path.join(homedir(), SECMAN_EDITOR_PATH);

export default class Init extends Command {
  static description = "Initialize ~/.secman .";

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Init);

    const initSpinner = spinner("💿 Initializing").start();

    if (platform() === "win32") {
      const ps = new powershell(
        `
        if (Test-Path -Path ~/.secman) {
          Write-Host "~/.secman already exists"
        } else {
          New-Item -ItemType "directory" -Path "${secman_dir}"
          New-Item ${sm_config}
          New-Item ${sm_data}
          // iwr -useb ${SECMAN_SETTINGS_URL} -o ${sm_setting}
        }
      `
      );

      initSpinner.stop();

      InstallEditor();
      writeSettingFile();

      ps.on("output", (data: any) => {
        console.log(data);
      });

      initSpinner.succeed(chalk.green("💿 Initialization complete"));
      console.log(
        chalk.bold(`run ${chalk.grey("`secman auth`")} to authenticate`)
      );
    } else {
      if (sh.test("-e", secman_dir)) {
        initSpinner.fail(chalk.red("💿 ~/.secman already exists"));
      } else {
        if (!fs.existsSync(secman_dir)) {
          fs.mkdirSync(secman_dir, { recursive: true });
        }

        if (!fs.existsSync(sm_config)) {
          sh.touch(sm_config);
        }

        if (!fs.existsSync(sm_data)) {
          sh.touch(sm_data);
        }

        if (!fs.existsSync(sm_setting)) {
          // sh.exec(`curl -s ${SECMAN_SETTINGS_URL} > ${sm_setting}`);
          sh.touch(sm_setting);
          writeSettingFile();
        }

        if (!fs.existsSync(sm_editor)) {
          InstallEditor();
        }

        if (fs.existsSync(secman_dir)) {
          initSpinner.succeed(chalk.green("💿 Initialization complete"));
          console.log(
            chalk.bold(`run ${chalk.grey("`secman auth`")} to authenticate`)
          );
        }
      }
    }
  }
}

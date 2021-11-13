import { Command, flags } from "@oclif/command";
import * as sh from "shelljs";
const powershell = require("powershell");
import { spinner } from "@secman/spinner";
import { platform } from "os";
import fs from "fs";
import {
  DOT_SECMAN_PATH,
  SECMAN_CONFIG_PATH,
  SECMAN_DATA_PATH,
  SECMAN_EDITOR_PATH,
  SECMAN_SETTINGS_PATH,
} from "../../constants";
import { InstallEditor } from "../../tools/install-editor";
import { writeCFile, writeDFile, writeSettingFile } from "../../app/config";
import { bold, command, error, success } from "../../design/layout";

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
          New-Item -ItemType "directory" -Path "${DOT_SECMAN_PATH}"
          New-Item ${SECMAN_CONFIG_PATH}
          New-Item ${SECMAN_DATA_PATH}
        }
      `
      );

      initSpinner.stop();

      InstallEditor();
      writeSettingFile();

      ps.on("output", (data: any) => {
        console.log(data);
      });

      initSpinner.succeed(success("💿 Initialization complete"));
      console.log(bold(`run ${command("`secman auth`")} to authenticate`));
    } else {
      if (sh.test("-e", DOT_SECMAN_PATH)) {
        initSpinner.fail(error("💿 ~/.secman already exists"));
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

        if (!fs.existsSync(SECMAN_SETTINGS_PATH)) {
          sh.touch(SECMAN_SETTINGS_PATH);
          writeSettingFile();
        }

        if (!fs.existsSync(SECMAN_EDITOR_PATH)) {
          InstallEditor();
        }

        if (fs.existsSync(DOT_SECMAN_PATH)) {
          initSpinner.succeed(success("💿 Initialization complete"));

          console.log(bold(`run ${command("`secman auth`")} to authenticate`));
        }
      }
    }
  }
}

import { Command, flags } from "@oclif/command";
import * as sh from "shelljs";
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
import { InstallEditor } from "../../tools/install_editor";
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
      fs.mkdirSync(DOT_SECMAN_PATH, { recursive: true });

      fs.writeFileSync(SECMAN_CONFIG_PATH, "");
      writeCFile();
      fs.writeFileSync(SECMAN_DATA_PATH, "");
      writeDFile();
      fs.writeFileSync(SECMAN_SETTINGS_PATH, "");

      initSpinner.stop();

      InstallEditor();
      writeSettingFile();

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

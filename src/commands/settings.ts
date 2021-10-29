import { Command, flags } from "@oclif/command";
import * as chalk from "chalk";
import * as sh from "shelljs";
import { platform } from "os";
import { spnr as spinner } from "@secman/spinner";
import { InstallEditor } from "../../tools/install-editor";
const prompts = require("prompts");
prompts.override(require("yargs").argv);
const powershell = require("powershell");

export default class Settings extends Command {
  static description = "Settings for the CLI.";

  static aliases = ["configs"];

  static flags = {
    help: flags.help({ char: "h" }),
    "editor-install": flags.boolean({
      char: "e",
      description: "Install editor",
    }),
    docs: flags.boolean({
      char: "d",
      description: "Read docs about secman settings",
    }),
  };

  async run() {
    const { flags } = this.parse(Settings);

    const opening = spinner("📜 Opening settings file...").start();

    if (flags.docs) {
      console.log(`
${chalk.bold("secman settings")}
----------------
${chalk.bold(
  "read_output"
)}: The output of reading password, values [ "table", "raw" ]
----------------
${chalk.bold("disable_version_check")}: Disable version check
`);
    } else if (flags["editor-install"]) {
      InstallEditor();
    } else {
      const editor = sh.find("~/.secman/editor");

      if (editor.length === 0) {
        const qe = await prompts({
          type: "toggle",
          name: "value",
          message: "The secman editor is not found. Do you want to install it?",
          active: "yes",
          inactive: "no",
        });

        if (qe.value) {
          InstallEditor();
        } else {
          this.exit(0);
        }
      } else {
        if (platform() === "win32") {
          const ps = new powershell(
            "$HOME/.secman/editor.exe $HOME/.secman/settings.json"
          );

          ps.on("output", (data: any) => {
            console.log(data);
          });

          opening.stop();
        } else {
          sh.exec("~/.secman/editor ~/.secman/settings.json");
          opening.stop();
        }
      }
    }

    // switch (true) {
    //   case readSettingsFile("editor") === "secman_editor":
    //     opening.start();

    //     if (platform() === "win32") {
    //       const ps = new powershell(
    //         "$HOME/.secman/editor.exe $HOME/.secman/settings.json"
    //       );

    //       ps.on("output", (data: any) => {
    //         console.log(data);
    //       });

    //       opening.stop();
    //     } else {
    //       const editor = sh.find("~/.secman/editor");

    //       if (editor.length === 0) {
    //         opening.stop();
    //         const qe = await prompts({
    //           type: "toggle",
    //           name: "value",
    //           message:
    //             "The secman editor is not found. Do you want to install it?",
    //           active: "yes",
    //           inactive: "no",
    //         });

    //         if (qe.value) {
    //           InstallEditor();
    //         } else {
    //           this.exit(0);
    //         }
    //       } else {
    //         sh.exec("~/.secman/editor ~/.secman/settings.json");
    //         opening.stop();
    //       }
    //     }

    //     break;
    // }
  }
}

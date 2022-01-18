import { Command, flags } from "@oclif/command";
import { spinner } from "@secman/spinner";
import { platform } from "os";
import * as sh from "shelljs";
const powershell = require("powershell");
import { GetLatest } from "../../api/latest";
import { bold } from "../../design/layout";

export default class Update extends Command {
  static description = "Update the secman CLI.";

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Update);

    const spnr = spinner("📦 Checking for updates...");
    const currentVersion = "v" + this.config.version;
    const latestVersion = await GetLatest("secman")
    const successMsg =
      "Secman CLI upgraded to " + latestVersion + " successfully";

    if (currentVersion === latestVersion) {
      spnr.succeed("already on latest version: " + bold(currentVersion));
    } else if (currentVersion !== latestVersion) {
      spnr.stop();

      const upgradingSpinner = spinner(
        "🚧 Upgrading Secman CLI from " +
          bold(currentVersion) +
          " to " +
          latestVersion +
          "\n"
      ).start();

      if (platform() === "win32") {
        const ps = new powershell("npm upgrade -g secman");

        ps.on("output", (data: any) => {
          console.log(data);
        });

        upgradingSpinner.succeed(successMsg);
      } else {
        sh.exec(`
          if [ "$SM_PROVIDER" = "script" ]; then
            sudo apt-get update && sudo apt-get update -y secman
          elif [ "$SM_PROVIDER" != "script" ]; then
            echo "if you are using a non-script provider, you must manually update the CLI"
            echo "like if you installed it with homebrew, run 'brew upgrade secman'"
          else
            npm upgrade -g secman
          fi
        `);

        upgradingSpinner.succeed(successMsg);
      }
    }
  }
}

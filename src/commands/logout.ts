import { Command, flags } from "@oclif/command";
import { spinner } from "@secman/spinner";
import writeConfigFile from "../../app/config";

export default class Logout extends Command {
  static description = "Logout of the current user account";

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Logout);

    const logoutSpinner = spinner("Logging out...");

    writeConfigFile(null, null, null, null, null, null, null);

    logoutSpinner.succeed("Logged out successfully");
  }
}

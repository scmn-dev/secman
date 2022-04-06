import { Command, flags } from "@oclif/command";
import { spinner } from "@secman/spinner";
import writeConfigFile, { readConfig } from "../../config";
import { API } from "../../contract";

export default class Logout extends Command {
  static description = "Logout of the current user account.";

  static aliases = ["signout"];

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Logout);

    const logoutSpinner = spinner("Logging out...");

    const data = {
      email: readConfig("email"),
    };

    await API.post("/auth/signout", data)
      .then(() => {
        writeConfigFile("", "", "", "", "", "", "");

        logoutSpinner.succeed("Logged out successfully");
      })
      .catch((err: any) => {
        logoutSpinner.fail("Failed to log out");
        console.log(err);
      });
  }
}

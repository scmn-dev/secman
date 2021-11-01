import { Command, flags } from "@oclif/command";
import { spinner } from "@secman/spinner";
import { readConfigFile, writeDataFile } from "../../app/config";
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
      email: readConfigFile("email"),
    };

    await API.post("/auth/signout", data)
      .then(() => {
        writeDataFile(null, null, null, null);

        logoutSpinner.succeed("Logged out successfully");
      })
      .catch((err: any) => {
        logoutSpinner.fail("Failed to log out");
        console.log(err);
      });
  }
}

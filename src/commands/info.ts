import { Command, flags } from "@oclif/command";
import { GetLatestGHRelease } from "../../api/github/api";
import { readConfigFile } from "../../app/config";

export default class Info extends Command {
  static description = "Information about the secman CLI.";

  static aliases = ["data"];

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Info);

    const smca_version = await GetLatestGHRelease("core");
    const name = readConfigFile("name") ?? "No User";

    console.log(`Secman CLI

Version: v${this.config.version}
Secman Core Version: ${smca_version}
Current User: ${name}`);
  }
}

import { Command, flags } from "@oclif/command";
import { Octokit } from "octokit";
import { readConfigFile } from "../../app/config";
import { GH_TOKEN } from "../../constants";

const octokit = new Octokit({
  auth: GH_TOKEN,
});

export default class Info extends Command {
  static description = "Info about the secman CLI.";

  static aliases = ["data"];

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Info);

    const smca_version = await octokit.rest.repos
      .listReleases({
        owner: "scmn-dev",
        repo: "core",
      })
      .then((res) => {
        return res.data[0].tag_name;
      });

    console.log(`Secman CLI

Version: v${this.config.version}
Secman Core Version: ${smca_version}
Current User: ${readConfigFile("name")}`);
  }
}

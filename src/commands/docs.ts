import { Command, flags } from "@oclif/command";
import { spinner } from "@secman/spinner";
import { cli } from "cli-ux";

export default class Docs extends Command {
  static description = "Open Secman Documentation in default browser.";

  static aliases = ["doc"];

  static flags = {
    help: flags.help({ char: "h" }),
  };

  async run() {
    const { flags } = this.parse(Docs);

    const opening = spinner("📚 Opening docs...").start();

    await cli.open("https://secman.dev/docs");

    opening.stop();
  }
}

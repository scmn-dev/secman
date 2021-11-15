import { Command, flags } from "@oclif/command";
import bcrypt from "bcrypt";
import { bold } from "../../design/layout";

export default class Bcrypt extends Command {
  static description = `Encrypt data using ${bold("bcrypt")}.`;

  static flags = {
    help: flags.help({ char: "h" }),
    length: flags.integer({
      char: "l",
      description: "Length of the output",
      default: 10,
    }),
  };

  static args = [{ name: "STRING" }];

  async run() {
    const { args, flags } = this.parse(Bcrypt);

    if (args.STRING) {
      let hash = await bcrypt.hash(args.STRING, flags.length);

      this.log(`String: ${args.STRING}
Hash: ${hash}`);
    } else {
      this.error("No string provided");
    }
  }
}

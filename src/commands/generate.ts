import { Command, flags } from "@oclif/command";
import { COMPLEXIES } from "../../constants";
import { GenerateExamples } from "../../contents/examples/generate";

export default class Generate extends Command {
  static description = "Generate a secure password.";

  static examples = GenerateExamples;

  static aliases = ["gen"];

  static flags = {
    help: flags.help({ char: "h" }),
    length: flags.integer({
      char: "l",
      description: "length of the generated password",
      default: 10,
    }),
    numbers: flags.boolean({ char: "n", description: "include numbers" }),
    symbols: flags.boolean({ char: "s", description: "include symbols" }),
    "capital-letters": flags.boolean({
      char: "c",
      description: "include capital letters",
    }),
  };

  async run() {
    const { flags } = this.parse(Generate);

    COMPLEXIES[1].checked = flags.numbers ?? false;
    COMPLEXIES[2].checked = flags.symbols ?? false;
    COMPLEXIES[3].checked = flags["capital-letters"] ?? false;

    let generatedPassword = "";

    const generator = COMPLEXIES.filter((item) => item.checked).reduce(
      (acc, current) => {
        return acc + current.value;
      },
      ""
    );

    for (let i = 0; i < flags.length; i++) {
      generatedPassword += generator.charAt(
        Math.floor(Math.random() * generator.length)
      );
    }

    console.log(generatedPassword);
  }
}

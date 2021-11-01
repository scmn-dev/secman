const chalk = require("chalk");

export async function root() {
  console.log(chalk.grey.bold("\nEXAMPLES\n"));
  console.log(chalk.grey("  -") + ` Initialize ${chalk.gray("~/.secman")}\n`);
  console.log(chalk.cyan("    $ secman init\n"));
  console.log(chalk.grey("  -") + " Authorize With Secman\n");
  console.log(chalk.cyan("    $ secman auth\n"));
  console.log(chalk.grey("  -") + " Insert a New Password\n");
  console.log(chalk.cyan("    $ secman insert --[PASSWORD_TYPE]\n"));
  console.log(chalk.grey("  -") + " List Passwords\n");
  console.log(chalk.cyan("    $ secman list\n"));
  console.log(chalk.grey("  -") + " Read The Password\n");
  console.log(
    chalk.cyan("    $ secman read --[PASSWORD_TYPE] <PASSWORD_NAME>\n")
  );
  console.log(chalk.grey("  -") + " Edit Password\n");
  console.log(
    chalk.cyan("    $ secman edit --[PASSWORD_TYPE] <PASSWORD_NAME>\n")
  );
  console.log(chalk.grey("  -") + " Generate\n");
  console.log(chalk.cyan("    $ secman generate\n"));
  console.log(chalk.grey("  -") + " Edit Settings\n");
  console.log(chalk.cyan("    $ secman settings\n"));
  learnMore();
  console.log(chalk.grey.bold("TELL US\n"));
  console.log(
    `  Open an issue at ${chalk.bold(
      "https://github.com/scmn-dev/secman/issues"
    )}\n`
  );
}

export function learnMore(command: any = "") {
  command = command || "";

  console.log(chalk.grey.bold("LEARN MORE"));
  console.log(`
  Use ${chalk.grey(
    "`secman " + (command || "<COMMAND>") + " --help`"
  )} for more information.
  Read 📚 at https://secman.dev/docs/cli${command ? `/${command}` : ""}
    `);
}

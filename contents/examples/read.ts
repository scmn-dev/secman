import chalk from "chalk";

export const ReadExamples = [
  `Read a Password\n\n ${chalk.cyan("$ secman read --logins | -l Expo")}\n`,
  `- Show The Password\n\n ${chalk.cyan(
    "$ secman read --emails | -e Gmail --show-password | -p"
  )}`,
];

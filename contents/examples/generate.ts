import chalk from "chalk";

export const GenerateExamples = [
  `Generate a Password\n\n ${chalk.cyan("$ secman generate")}\n`,
  `- Spicefy The Password Length\n\n ${chalk.cyan(
    "$ secman generate --length=13 | -l=13"
  )}\n`,
  `- Include Numbers\n\n ${chalk.cyan("$ secman generate --numbers | -n")}\n`,
  `- Include Symbols\n\n ${chalk.cyan("$ secman generate --symbols | -s")}\n`,
  `- Include Capital Letters\n\n ${chalk.cyan(
    "$ secman generate --capital-letters | -c"
  )}\n`,
];

import { withSecondary } from "../../design/layout";

export const ReadExamples = [
  `Read a Password\n\n ${withSecondary("$ secman read --logins | -l Expo")}\n`,
  `- Show The Password\n\n ${withSecondary(
    "$ secman read --emails | -e Gmail --show-password | -p"
  )}`,
];

import { withSecondary } from "../../design/layout";

export const GenerateExamples = [
  `Generate a Password\n\n ${withSecondary("$ secman generate")}\n`,
  `- Spicefy The Password Length\n\n ${withSecondary(
    "$ secman generate --length 13 | -l 13"
  )}\n`,
  `- Include Numbers\n\n ${withSecondary("$ secman generate --numbers | -n")}\n`,
  `- Include Symbols\n\n ${withSecondary("$ secman generate --symbols | -s")}\n`,
  `- Include Capital Letters\n\n ${withSecondary(
    "$ secman generate --capital-letters | -c"
  )}\n`,
];

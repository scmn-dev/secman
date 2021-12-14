import { withSecondary } from "../../design/layout";

export const ListExamples = [
  `List Passwords\n\n ${withSecondary("$ secman list")}\n`,
  `- List Password from spicefic password type\n\n ${withSecondary(
    "$ secman . --servers | -s"
  )}`,
];

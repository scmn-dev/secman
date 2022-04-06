import { withSecondary } from "../../design/layout";

export const AuthExamples = [
  `Authenticate With Secman\n\n ${withSecondary("$ secman auth")}\n`,
  `- Authinticate With Email and Master Password\n\n ${withSecondary(
    "$ secman auth --user EMAIL --master-password MASTER_PASSWORD"
  )}\n`,
  `- Read Master Password from stdin\n\n ${withSecondary(
    "$ echo $PASSWORD | secman auth --user EMAIL --password-stdin"
  )}\n`,
];

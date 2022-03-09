import { bold, command as cmd, withSecondary } from "../../design/layout";

export async function root() {
  console.log(cmd("\nEXAMPLES\n", true));
  console.log(cmd("  -") + ` Initialize ${cmd("~/.secman")}\n`);
  console.log(withSecondary("    $ secman init\n"));
  console.log(cmd("  -") + " Authorize With Secman\n");
  console.log(withSecondary("    $ secman auth\n"));
  console.log(cmd("  -") + " Insert a New Password\n");
  console.log(withSecondary("    $ secman insert --[PASSWORD_TYPE]\n"));
  console.log(cmd("  -") + " List Passwords\n");
  console.log(withSecondary("    $ secman .\n"));
  console.log(cmd("  -") + " Read The Password\n");
  console.log(
    withSecondary("    $ secman read --[PASSWORD_TYPE] <PASSWORD_NAME>\n")
  );
  console.log(cmd("  -") + " Edit Password\n");
  console.log(
    withSecondary("    $ secman edit --[PASSWORD_TYPE] <PASSWORD_NAME>\n")
  );
  console.log(cmd("  -") + " Generate\n");
  console.log(withSecondary("    $ secman generate\n"));
  console.log(cmd("  -") + " Edit Settings\n");
  console.log(withSecondary("    $ secman settings\n"));
  console.log(cmd("TELL US\n", true));
  console.log(
    `  Open an issue at ${bold("https://github.com/scmn-dev/secman/issues")}\n`
  );
}

export function end() {}

const { Command, flags } = require("@oclif/command");
const PowerShell = require("powershell");
const { LOC, CDCmd, lastCheck, mainCode } = require("./shared");

class UninstallCommand extends Command {
  async run() {
    const { flags } = this.parse(UninstallCommand);
    const deleteDataCmd = `
      ${LOC}

      ${mainCode}

      ${CDCmd}

      ${lastCheck}
    `;

    const mainCmd = `
      ${LOC}

      ${mainCode}

      $ClearData = Read-Host -Prompt "Clear all data?/n[y/N]"

      switch($ClearData) {
        {($_ -eq "y") -or ($_ -eq "yes") -or ($_ -eq "Y") -or ($_ -eq "Yes")} {
          ${CDCmd}
        }

        {($_ -eq "n") -or ($_ -eq "no") -or ($_ -eq "N") -or ($_ -eq "No") -or ($_ -eq "")} {
          Write-Host "Ok" -ForegroundColor DarkGreen
        }

        default {
          Write-Host "wrong input"
        }
      }

      ${lastCheck}
    `;

    if (flags["delete-data"]) {
      let ps = new PowerShell(deleteDataCmd);

      ps.on("output", (out) => {
        console.log(out);
      });
    } else {
      let ps = new PowerShell(mainCmd);

      ps.on("output", (out) => {
        console.log(out);
      });
    }
  }
}

UninstallCommand.description = `Uninstall your secman`;

UninstallCommand.flags = {
  "delete-data": flags.boolean({
    char: "d",
    description: "delete data (~/.secman)",
  }),
};

module.exports = UninstallCommand;

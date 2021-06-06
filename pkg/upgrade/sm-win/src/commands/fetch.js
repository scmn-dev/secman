const { Command } = require("@oclif/command");
const PowerShell = require("powershell");
const { Check } = require("../shared");

class FetchCommand extends Command {
  async run() {
    let ps = new PowerShell(`
      ${Check}

      if ($l -ne $c) {
        $nr = "there's a new release of secman is avalaible:"
        $up = "to upgrade run "
        $smu = "sm-win start"

        Write-Host ""
        Write-Host -NoNewline $nr -ForegroundColor DarkYellow
        Write-Host "$c → $l" -ForegroundColor DarkCyan
        Write-Host -NoNewline $up -ForegroundColor DarkYellow
        Write-Host $smu -ForegroundColor DarkCyan
      }
    `);

    ps.on("output", (data) => {
      console.log(data);
    });
  }
}

FetchCommand.description = `Fetch if there's a new release`;

FetchCommand.flags = {};

module.exports = FetchCommand;

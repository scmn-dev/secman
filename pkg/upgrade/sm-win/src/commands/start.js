const { Command } = require("@oclif/command");
const PowerShell = require("powershell");
const { LOC, Check } = require("../shared");

class StartCommand extends Command {
  async run() {
    let ps = new PowerShell(`
      ${LOC}

      ${Check}

      if ($l -ne $c) {
        Remove-Item -Force -Recurse $loc
  
        iwr -useb https://win-upg.secman.dev | iex

        Write-Host "secman was upgraded successfully 🎊"
      } else {
        Write-Host "secman is already up-to-date and it's the latest release $l"
      }
    `);

    ps.on("output", (data) => {
      console.log(data);
    });
  }
}

StartCommand.description = `Start Upgrade secman`;

StartCommand.flags = {};

module.exports = StartCommand;

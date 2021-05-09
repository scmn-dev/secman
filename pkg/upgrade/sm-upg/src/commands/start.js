const { Command } = require("@oclif/command");
const PowerShell = require("powershell");
const { LOC } = require("./shared");

class StartCommand extends Command {
  async run() {
    let ps = new PowerShell(`
      ${LOC}
      Remove-Item $loc

      iwr -useb https://get.secman.dev/install.ps1 | iex
    `);

    ps.on("output", (data) => {
      console.log(data);
    });
  }
}

StartCommand.description = `Start Upgrade secman`;

StartCommand.flags = {};

module.exports = StartCommand;

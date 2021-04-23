const { Command, flags } = require("@oclif/command");
const PowerShell = require("powershell");

class StartCommand extends Command {
  async run() {
    let ps = new PowerShell("& $HOME/sm/vx.ps1 --upg");

    ps.on("output", (data) => {
      console.log(data);
    });
  }
}

StartCommand.description = `Start Upgrade secman`;

StartCommand.flags = {};

module.exports = StartCommand;

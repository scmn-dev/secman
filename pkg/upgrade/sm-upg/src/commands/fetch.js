const { Command } = require("@oclif/command");
const PowerShell = require("powershell");

class FetchCommand extends Command {
  async run() {
    let ps = new PowerShell("& $HOME/sm/vx.ps1 --sm");

    ps.on("output", (data) => {
      console.log(data);
    });
  }
}

FetchCommand.description = `Fetch if there's a new release`;

FetchCommand.flags = {};

module.exports = FetchCommand;

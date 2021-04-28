const { Command, flags } = require("@oclif/command");
const PowerShell = require("powershell");

class UninstallCommand extends Command {
  async run() {
    const { flags } = this.parse(UninstallCommand);
    const coreCmd = "& $HOME/sm/uninstall.ps1";

    if (flags["delete-data"]) {
      let ps = new PowerShell(`${coreCmd} -d`);

      ps.on("output", (out) => {
        console.log(out);
      });
    } else {
      let ps = new PowerShell(coreCmd);

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

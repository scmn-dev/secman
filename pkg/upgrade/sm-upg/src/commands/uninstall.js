const { Command, flags } = require("@oclif/command");
const PowerShell = require("powershell");
const { LOC, MainCode, LastCheck, CDCmd } = require("../shared");

class UninstallCommand extends Command {
  async run() {
    const { flags } = this.parse(UninstallCommand);
    const deleteDataCmd = `
      ${LOC}

      ${MainCode}

      ${CDCmd}

      ${LastCheck}
    `;

    const mainCmd = `
      ${LOC}

      ${MainCode}

      ${LastCheck}
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

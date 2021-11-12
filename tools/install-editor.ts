import * as sh from "shelljs";
const powershell = require("powershell");
import { platform } from "os";

export const InstallEditor = () => {
  if (platform() === "win32") {
    const ps = new powershell("iwr -useb https://win-editor.secman.dev | iex");

    ps.on("output", (data: any) => {
      console.log(data);
    });
  } else {
    sh.exec("curl -fsSL https://unix-editor.secman.dev | bash", {
      silent: true,
    }).stdout;
  }
};

import * as process from "process";

export function isUnicodeSupported() {
  if (process.platform !== "win32") {
    return process.env.TERM !== "linux";
  }

  return (
    Boolean(process.env.CI) ||
    Boolean(process.env.WT_SESSION) ||
    process.env.ConEmuTask === "{cmd::Cmder}" ||
    process.env.TERM_PROGRAM === "vscode" ||
    process.env.TERM === "xterm-256color" ||
    process.env.TERM === "alacritty"
  );
}

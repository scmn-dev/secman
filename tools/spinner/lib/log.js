const { isUnicodeSupported } = require("./unicode");
const chalk = require("chalk");

const main = {
  info: chalk.blue("ℹ"),
  success: chalk.green("✔"),
  warning: chalk.yellow("⚠"),
  error: chalk.red("✖"),
};

const fallback = {
  info: chalk.blue("i"),
  success: chalk.green("√"),
  warning: chalk.yellow("‼"),
  error: chalk.red("×"),
};

const logSymbols = isUnicodeSupported() ? main : fallback;

module.exports.logSymbols = logSymbols;

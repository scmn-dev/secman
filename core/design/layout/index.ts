import chalk from "chalk";
import { PRIMARY_COLOR } from "../../constants";

export const withPrimary = (text: string) =>
  chalk.hex(PRIMARY_COLOR).bold(text);

export const withSecondary = (text: string) => chalk.cyan(text);

export const bold = (text: string) => chalk.bold(text);

export const underline = (text: string) => chalk.underline(text);

export const dim = (text: string) => {
  if (process.env.ConEmuANSI === "ON") {
    return command(text);
  } else {
    return chalk.dim(text);
  }
};

export const error = (text: string, bold: boolean = false) => {
  if (bold) {
    return chalk.red.bold(text);
  } else {
    return chalk.red(text);
  }
};

export const success = (text: string, bold: boolean = false) => {
  if (bold) {
    return chalk.green.bold(text);
  } else {
    return chalk.green(text);
  }
};

export const command = (text: string, bold: boolean = false) => {
  if (bold) {
    return chalk.gray.bold(text);
  } else {
    return chalk.gray(text);
  }
};

export const warning = (text: string, bold: boolean = false) => {
  if (bold) {
    return chalk.yellow.bold(text);
  } else {
    return chalk.yellow(text);
  }
};

export const blue = (text: string) => chalk.blue(text);

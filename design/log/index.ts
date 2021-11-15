const { isUnicodeSupported } = require("../../tools/unicode");
import { blue, error, success, warning } from "../layout";

const main = {
  info: blue("ℹ"),
  success: success("✔"),
  warning: warning("⚠"),
  error: error("✖"),
};

const fallback = {
  info: blue("i"),
  success: success("√"),
  warning: warning("‼"),
  error: error("×"),
};

export const logSymbols = isUnicodeSupported() ? main : fallback;

import { platform } from "os";
import env from "dotenv";

env.config();

export const PRIMARY_COLOR: any = "#1163E6";
export const GH_TOKEN: any = process.env.GH_TOKEN;
export const API_URL: any = "https://api.secman.dev";
export const CHARS = {
  alphabet: "abcdefghijklmnopqrstuvwxyz",
  numeric: "0123456789",
  special: "_-+=)/(*&^%$#@%!?~",
};
export const DOT_SECMAN_PATH = `${process.env.HOME}/.secman`;
export const SECMAN_CONFIG_PATH: any = `/${DOT_SECMAN_PATH}/config.json`;
export const SECMAN_DATA_PATH: any = `/${DOT_SECMAN_PATH}/data.json`;
export const SECMAN_SETTINGS_PATH: any = `/${DOT_SECMAN_PATH}/settings.json`;
export const SECMAN_EDITOR_PATH: any =
  platform() === "win32"
    ? `/${DOT_SECMAN_PATH}/editor.exe`
    : `/${DOT_SECMAN_PATH}/editor`;
export let COMPLEXIES = [
  { name: "abc", value: CHARS.alphabet, checked: true, visible: false },
  { name: "Numbers", value: CHARS.numeric, checked: true },
  { name: "Symbols", value: CHARS.special, checked: false },
  {
    name: "Capital Letters",
    value: CHARS.alphabet.toUpperCase(),
    checked: true,
  },
];
export const LOGINS_ENCRYPTED_FIELDS = ["username", "password", "extra"];
export const CCS_ENCRYPTED_FIELDS = [
  "type",
  "number",
  "expiry_date",
  "cardholder_name",
  "verification_number",
];
export const EMAILS_ENCRYPTED_FIELDS = ["email", "password"];
export const NOTES_ENCRYPTED_FIELDS = ["note"];
export const SERVERS_ENCRYPTED_FIELDS = [
  "ip",
  "username",
  "password",
  "hosting_username",
  "hosting_password",
  "admin_username",
  "admin_password",
  "extra",
];
export const TABLE_DESIGN = {
  border: {
    topBody: `─`,
    topJoin: `┬`,
    topLeft: `╭`,
    topRight: `╮`,

    bottomBody: `─`,
    bottomJoin: `┴`,
    bottomLeft: `╰`,
    bottomRight: `╯`,

    bodyLeft: `│`,
    bodyRight: `│`,
    bodyJoin: `│`,

    joinBody: `─`,
    joinLeft: `├`,
    joinRight: `┤`,
    joinJoin: `┼`,
  },
};

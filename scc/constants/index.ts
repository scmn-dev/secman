import path from "path";

export const PRIMARY_COLOR: any = "#1163E6";
export const API_URL: any = "https://api.secman.dev";
export const HOMEDIR: any = process.env.HOME || process.env.USERPROFILE;
export const CHARS = {
  alphabet: "abcdefghijklmnopqrstuvwxyz",
  numeric: "0123456789",
  special: "_-+=)/(*&^%$#@%!?~",
};
export const DOT_SECMAN_PATH = path.join(HOMEDIR, ".secman");
export const SECMAN_CONFIG_PATH: any = path.join(
  DOT_SECMAN_PATH,
  "secman.json"
);
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

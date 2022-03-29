import { CryptoTools } from "./main";

const string = "Hi Secman";

const hash = CryptoTools.sha256Encrypt(string);

console.log(`String: ${string}\nHash: ${hash}`);

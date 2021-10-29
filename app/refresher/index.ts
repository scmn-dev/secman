import { readConfigFile, writeDataFile } from "../config";
import { spnr as spinner } from "@secman/spinner";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import * as cryptojs from "crypto-js";
import chalk from "chalk";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export async function refresh() {
  let password = await prompts({
    type: "password",
    name: "mp",
    message:
      "Your Authentication is Expired. Please enter your Master Password",
    validate: (value: string) => {
      if (value.length > 0) {
        return true;
      } else {
        return "Please enter your master password";
      }
    },
  });

  let master_password = password.mp;

  let hash = CryptoTools.sha256Encrypt(master_password);
  let pswd = hash.toString();

  let data = JSON.stringify({
    email: readConfigFile("user"),
    master_password: pswd,
  });

  await API.post("/auth/signin", data)
    .then(async function (res: any) {
      let { access_token, refresh_token, transmission_key, secret } = res.data;

      const mp = CryptoTools.sha256Encrypt(master_password);
      const master_password_hash = CryptoTools.pbkdf2Encrypt(secret, mp);

      const refreshSpinner = spinner("🔑 Refreshing token...").start();

      writeDataFile(
        access_token,
        refresh_token,
        transmission_key,
        master_password_hash
      );

      refreshSpinner.stop();

      refreshSpinner.succeed("🔗 Refreshed");
      console.log(chalk.bold("run the command again"));
    })
    .catch(function (err: any) {
      console.log(err);
    });
}

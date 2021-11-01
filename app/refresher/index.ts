import { readConfigFile, writeDataFile } from "../config";
import { spinner } from "@secman/spinner";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import chalk from "chalk";
import { PRIMARY_COLOR } from "../../constants";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export async function refresh(cmd: any) {
  const password = await prompts({
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

  const master_password = password.mp;

  const pswd = CryptoTools.sha256Encrypt(master_password).toString();

  const data = JSON.stringify({
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
      console.log(
        chalk.bold(
          `run ${chalk.hex(PRIMARY_COLOR).bold("secman " + cmd)} command again`
        )
      );
    })
    .catch(function (err: any) {
      console.log(err);
    });
}

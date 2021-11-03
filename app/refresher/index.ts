import { readConfigFile, readDataFile, writeDataFile } from "../config";
import { spinner } from "@secman/spinner";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import chalk from "chalk";
import { PRIMARY_COLOR } from "../../constants";
import { stdout } from "process";
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

  let master_password = password.mp;
  master_password = CryptoTools.encrypt(master_password);

  let data: any;
  const func = (res: any) => {
    let { access_token, refresh_token, transmission_key, secret } = res.data;

    const master_password_hash = CryptoTools.pbkdf2Encrypt(
      secret,
      master_password
    );

    const refreshSpinner = spinner("🔑 Refreshing token...").start();

    writeDataFile(
      access_token,
      refresh_token,
      transmission_key,
      master_password_hash
    ).then(async () => {
      refreshSpinner.succeed("🔗 Refreshed");

      stdout.write(
        chalk.bold(
          `run ${chalk.hex(PRIMARY_COLOR).bold("secman " + cmd)} command again`
        )
      );
    });
  };

  await API.post("/auth/refresh", data)
    .then(async (res: any) => {
      data = JSON.stringify({
        refresh_token: readDataFile("refresh_token"),
      });

      func(res);
    })
    .catch(async (err: any) => {
      if (err.response.status === 401) {
        data = JSON.stringify({
          email: readConfigFile("user"),
          master_password: master_password,
        });

        await API.post("/auth/signin", data)
          .then(async (res: any) => {
            func(res);
          })
          .catch((err: any) => {
            stdout.write(err);
          });
      }
    });
}

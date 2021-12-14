import { readConfigFile, readDataFile, writeDataFile } from "../config";
import { spinner } from "@secman/spinner";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import { bold, withPrimary } from "../../design/layout";
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
  const hash = CryptoTools.sha256Encrypt(master_password);
  const pswd = hash.toString();

  if (master_password) {
    let data;

    const func = (res: any) => {
      let { access_token, refresh_token, transmission_key, secret } = res.data;

      const master_password_hash = CryptoTools.pbkdf2Encrypt(secret, pswd);

      const refreshSpinner = spinner("🔑 Refreshing token...").start();

      writeDataFile(
        access_token,
        refresh_token,
        transmission_key,
        master_password_hash
      ).then(async () => {
        refreshSpinner.succeed("🔗 Refreshed");

        console.log(bold(`run ${withPrimary("secman " + cmd)} command again`));
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
            master_password: pswd,
          });

          await API.post("/auth/signin", data)
            .then(async (res: any) => {
              func(res);
            })
            .catch((err: any) => {
              console.log(err);
            });
        }
      });
  }
}

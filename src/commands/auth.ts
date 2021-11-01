import { Command, flags } from "@oclif/command";
import chalk from "chalk";
import SendEmail from "../../app/email/send";
import { PRIMARY_COLOR } from "../../constants";
import tokenGenerator from "../../api/generator";
import writeConfigFile, { readConfigFile } from "../../app/config";
import { API } from "../../contract";
import { spinner } from "@secman/spinner";
import { CryptoTools } from "../../tools/crypto";
const prompts = require("prompts");
prompts.override(require("yargs").argv);
import store from "store";

export default class Auth extends Command {
  static description = "Manage secman's authentication state.";

  static flags = {
    help: flags.help({ char: "h" }),
  };

  static aliases = ["login", "signin"];

  async run() {
    const { flags } = this.parse(Auth);
    // let email = readlineSync.questionEMail("Enter your email: ");

    const email =
      readConfigFile("user") ??
      (
        await prompts({
          type: "text",
          name: "e",
          message: "Enter your email: ",
          validate: (value: string) => {
            if (value.length > 0) {
              return true;
            } else {
              return "Please enter your email";
            }
          },
        })
      ).e;

    try {
      let master_password = await prompts({
        type: "password",
        name: "mp",
        message: "Enter your master password: ",
        validate: (value: string) => {
          if (value.length > 0) {
            return true;
          } else {
            return "Please enter your master password";
          }
        },
      }).mp;

      const hash = CryptoTools.sha256Encrypt(master_password);

      const pswd = hash.toString();

      const data = JSON.stringify({
        email: email,
        master_password: pswd,
      });

      await API.post("/auth/signin", data)
        .then(function (res: any) {
          let {
            access_token,
            refresh_token,
            transmission_key,
            email,
            name,
            secret,
          } = res.data;

          master_password = CryptoTools.sha256Encrypt(master_password);

          const master_password_hash = CryptoTools.pbkdf2Encrypt(
            secret,
            master_password
          );

          writeConfigFile(
            name,
            email,
            access_token,
            refresh_token,
            transmission_key,
            master_password_hash,
            secret
          );

          console.log(
            "\n🎉 Welcome",
            chalk.hex(PRIMARY_COLOR).bold(name) + "!"
          );
        })
        .catch(function (err: any) {
          console.log(err);
        });
    } catch (error) {
      console.log(error);
    }
  }
}

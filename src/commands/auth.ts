import { Command, flags } from "@oclif/command";
import chalk from "chalk";
import SendEmail from "../../app/email/send";
import { PRIMARY_COLOR } from "../../constants";
import crypto from "crypto";
import tokenGenerator from "../../api/generator";
import writeConfigFile from "../../app/config";
import { API } from "../../contract";
import { spnr as spinner } from "@secman/spinner";
import { CryptoTools } from "../../tools/crypto";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export default class Auth extends Command {
  static description = "Manage secman's authentication state.";

  static flags = {
    help: flags.help({ char: "h" }),
  };

  static aliases = ["login"];

  async run() {
    const { flags } = this.parse(Auth);

    // let email = readlineSync.questionEMail("Enter your email: ");
    let email = await prompts({
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
    });

    const user_email = email.e;

    let generatedPassword = "";
    for (let i = 0; i < 10; i++) {
      generatedPassword += tokenGenerator.charAt(
        Math.floor(Math.random() * tokenGenerator.length)
      );
    }

    const emailSpinner = spinner("📮 Sending Email ...").start();

    const func = async () => {
      emailSpinner.succeed(
        `We sent an email to ${chalk.underline.bold(
          user_email
        )}. please copy the security code and insert it here.`
      );

      let securityCode = await prompts({
        type: "text",
        name: "sc",
        message: "Enter the security code: ",
        validate: (value: string) => {
          if (value.length > 0) {
            return true;
          } else {
            return "Please enter the security code";
          }
        },
      });

      const sc = securityCode.sc;

      if (sc === generatedPassword) {
        const verifiedSpinner = spinner("📁 Verfiying ...").start();

        try {
          verifiedSpinner.succeed("Verified successfully.");

          let master_password = await prompts({
            type: "password",
            name: "mp",
            message: "Enter your master password: ",jo838z8mw8ialtvbgmaeqnpxb6ndgn1q
            validate: (value: string) => {
              if (value.length > 0) {
                return true;
              } else {
                return "Please enter your master password";
              }
            },
          });

          let hash = CryptoTools.sha256Encrypt(master_password.mp);

          let pswd = hash.toString();

          let data = JSON.stringify({
            email: user_email,
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

              const mp = CryptoTools.sha256Encrypt(master_password.mp);
              const master_password_hash = CryptoTools.pbkdf2Encrypt(
                secret,
                mp
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
      } else {
        console.log("Incorrect");
      }
    };

    SendEmail(user_email, generatedPassword, func);
  }
}

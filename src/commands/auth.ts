import { Command, flags } from "@oclif/command";
import chalk from "chalk";
import { PRIMARY_COLOR } from "../../constants";
import writeConfigFile, { readConfigFile } from "../../app/config";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import { cli as ux } from "cli-ux";
import { readPipe } from "../../tools/readPipe";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export default class Auth extends Command {
  static description = "Manage secman's authentication state.";

  static flags = {
    help: flags.help({ char: "h" }),
    "create-account": flags.boolean({
      char: "c",
      description: "Create a new account.",
      default: false,
    }),
    email: flags.string({
      char: "e",
      description: "Email address of the account to use.",
    }),
    "master-password": flags.string({
      char: "p" || "m",
      description: "Master password of the account to use.",
    }),
    "password-stdin": flags.boolean({
      description: "Take the password from stdin",
      default: false,
    }),
  };

  static aliases = ["login", "signin"];

  async run() {
    const { flags } = this.parse(Auth);
    const configFile = `${process.env.HOME}/.secman/config.json`;

    const _ = async (isNewLogin: boolean) => {
      let email;

      if (flags.email) {
        email = flags.email;
      } else {
        email =
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
      }

      try {
        let password;
        let master_password: any;

        if (flags["master-password"]) {
          password = flags["master-password"];
          master_password = password;
        } else if (flags["password-stdin"]) {
          const stdin = await readPipe();

          if (stdin) {
            password = stdin.replace(/\s/g, "");
            master_password = password;
          }
        } else {
          password = await prompts({
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
          });
          master_password = password.mp;
          console.log("");
        }

        if (master_password) {
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

              const msg = isNewLogin
                ? "🎉 Welcome " + chalk.hex(PRIMARY_COLOR).bold(name) + "!"
                : "Re-authentication successful";

              console.log(msg);
            })
            .catch(function (err: any) {
              if (err.response.status === 401) {
                console.log(
                  chalk.red.bold(
                    `\nInvalid email or master password. if you don't have an account, please create one using the command ${chalk.gray.bold(
                      "`secman auth --create-account`."
                    )}`
                  )
                );
              } else {
                console.log(
                  chalk.red.bold("\nSomething went wrong. Please try again.")
                );
              }
            });
        }
      } catch (error) {
        console.log(error);
      }
    };

    switch (true) {
      case flags["create-account"]:
        ux.open("https://auth.secman.dev");

        break;

      default:
        if (configFile) {
          const user = readConfigFile("name");

          if (user != null) {
            const reauth = await prompts({
              type: "toggle",
              name: "value",
              message: `You are already logged in as ${chalk
                .hex(PRIMARY_COLOR)
                .bold(user)}. Would you like to re-authenticate?`,
              initial: "yes",
              active: "yes",
              inactive: "no",
            });

            console.log("");

            if (reauth.value) {
              _(false);
            } else {
              this.log("Re-authentication cancelled");
            }
          } else {
            _(true);
          }
        } else {
          _(true);
        }
    }
  }
}

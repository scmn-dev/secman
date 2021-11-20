import { Command, flags } from "@oclif/command";
import { PRIMARY_COLOR } from "../../constants";
import writeConfigFile, { readConfigFile } from "../../app/config";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import { cli as ux } from "cli-ux";
import { readPipe } from "../../tools/readPipe";
import { command, withPrimary } from "../../design/layout";
import { AuthExamples } from "../../contents/examples/auth";
import { HOMEDIR } from "../../constants";
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
    user: flags.string({
      char: "u",
      description: "Email address of the account to use.",
    }),
    "master-password": flags.string({
      char: "p",
      description: "Master password of the account to use.",
    }),
    "password-stdin": flags.boolean({
      description: "Take the password from stdin",
      default: false,
    }),
  };

  static aliases = ["login", "signin"];

  static examples = AuthExamples;

  async run() {
    const { flags } = this.parse(Auth);
    const configFile = `${HOMEDIR}/.secman/config.json`;

    const _ = async (isNewLogin: boolean) => {
      let email;

      if (flags.user) {
        email = flags.user;
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
                ? "🎉 Welcome " + withPrimary(name) + "!"
                : "Re-authentication successful";

              console.log(msg);
            })
            .catch(function (err: any) {
              if (err.response.status === 401) {
                console.log(
                  command(
                    `\nInvalid email or master password. if you don't have an account, please create one using the command ${
                      (command("`secman auth --create-account`."), true)
                    }`,
                    true
                  )
                );
              } else {
                console.log(
                  command("\nSomething went wrong. Please try again.", true)
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
              message: `You are already logged in as ${withPrimary(
                user
              )}. Would you like to re-authenticate?`,
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

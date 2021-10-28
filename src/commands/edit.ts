import { Command, flags } from "@oclif/command";
import {
  CCS_ENCRYPTED_FIELDS,
  EMAILS_ENCRYPTED_FIELDS,
  LOGINS_ENCRYPTED_FIELDS,
  NOTES_ENCRYPTED_FIELDS,
  SERVERS_ENCRYPTED_FIELDS,
} from "../../constants";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";
import * as cryptojs from "crypto-js";
import * as chalk from "chalk";
import { spnr as spinner } from "@secman/spinner";
import { readDataFile } from "../../app/config";
import {
  CCFields,
  EmailFields,
  LoginFields,
  NoteFields,
  ServerFields,
} from "../../contents/types";
import { refresh } from "../../app/refresher";
import { EditExamples } from "../../contents/examples/edit";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export default class Edit extends Command {
  static description = "Update or change a value in a password.";

  static aliases = ["show", "print"];

  static flags = {
    help: flags.help({ char: "h" }),
    logins: flags.boolean({
      char: "l",
      description: "edit password from logins type.",
    }),
    "credit-cards": flags.boolean({
      char: "c",
      description: "edit password from credit cards type.",
    }),
    emails: flags.boolean({
      char: "e",
      description: "edit password from emails type.",
    }),
    notes: flags.boolean({
      char: "n",
      description: "edit password from notes type.",
    }),
    servers: flags.boolean({
      char: "s",
      description: "edit password from servers type.",
    }),
    "show-password": flags.boolean({
      char: "p",
      description: "show password.",
    }),
    multi: flags.boolean({
      char: "m",
      description: "edit multiple fields.",
    }),
  };

  static args = [{ name: "PASSWORD_NAME" }];

  static examples = EditExamples;

  async run() {
    const { args, flags } = this.parse(Edit);
    let API_URL = "/api";
    const access_token = readDataFile("access_token");

    if (flags.logins) {
      API_URL += "/logins";
    } else if (flags["credit-cards"]) {
      API_URL += "/credit-cards";
    } else if (flags.emails) {
      API_URL += "/emails";
    } else if (flags.notes) {
      API_URL += "/notes";
    } else if (flags.servers) {
      API_URL += "/servers";
    } else {
      this.error("Incorrect type of entry.");
    }

    const gettingDataSpinner = spinner("📡 Getting data...").start();

    await API.get(API_URL, {
      headers: {
        Authorization: `Bearer ${access_token}`,
        Cookie: `secman_token=${access_token}`,
      },
    })
      .then(async (res: any) => {
        if (res.status === 200 || res.status === 202) {
          gettingDataSpinner.stop();

          const item: any = cryptojs.AES.decrypt(
            res.data.data,
            readDataFile("transmission_key")
          ).toString(cryptojs.enc.Utf8);

          const itemList = JSON.parse(item);

          const ms_hash = readDataFile("master_password_hash");

          const enc_fields = () => {
            if (flags.logins) {
              return LOGINS_ENCRYPTED_FIELDS;
            } else if (flags["credit-cards"]) {
              return CCS_ENCRYPTED_FIELDS;
            } else if (flags.emails) {
              return EMAILS_ENCRYPTED_FIELDS;
            } else if (flags.notes) {
              return NOTES_ENCRYPTED_FIELDS;
            } else if (flags.servers) {
              return SERVERS_ENCRYPTED_FIELDS;
            } else {
              this.error("Incorrect type of entry.");
            }
          };

          itemList.forEach(async (element: any) => {
            if (element.title === args.PASSWORD_NAME) {
              CryptoTools.decryptFields(element, enc_fields(), ms_hash);

              const fields = () => {
                if (flags.logins) {
                  return LoginFields;
                } else if (flags["credit-cards"]) {
                  return CCFields;
                } else if (flags.emails) {
                  return EmailFields;
                } else if (flags.notes) {
                  return NoteFields;
                } else if (flags.servers) {
                  return ServerFields;
                } else {
                  this.error("Incorrect type of entry.");
                }
              };

              // const response = await prompts([
              //   {
              //     type: "select",
              //     name: "value",
              //     message: "Pick a field",
              //     choices: fields(),
              //   },
              // ]);

              // const newValue = await prompts({
              //   type: "text",
              //   name: "value",
              //   message: `Enter the new ${response.value} of ${element.title}`,
              // });

              let response;
              let newValue;

              if (flags.multi) {
                response = await prompts([
                  {
                    type: "multiselect",
                    name: "value",
                    message: "Pick a field",
                    choices: fields(),
                  },
                ]);

                for (const field of response.value) {
                  newValue = await prompts({
                    type: "text",
                    name: "value",
                    message: `Enter the new ${field} of ${element.title}`,
                  });

                  element[field] = newValue.value;
                }

                if (response.value.length === 0) {
                  this.error("You didn't pick any fields.");
                }
              } else {
                response = await prompts([
                  {
                    type: "select",
                    name: "value",
                    message: "Pick a field",
                    choices: fields(),
                  },
                ]);

                newValue = await prompts({
                  type: "text",
                  name: "value",
                  message: `Enter the new ${response.title} of ${element.title}`,
                });

                element[response.value] = newValue.value;
              }

              const payload = CryptoTools.encryptPayload(
                element,
                enc_fields(),
                readDataFile("master_password_hash"),
                readDataFile("transmission_key")
              );

              await API.put(API_URL + `/${element.id}`, payload, {
                headers: {
                  Authorization: `Bearer ${access_token}`,
                  Cookie: `secman_token=${access_token}`,
                },
              }).then(async (res: any) => {
                if (res.status === 200 || res.status === 202) {
                  console.log(chalk.green("Password updated"));
                } else {
                  console.log(chalk.red("Password not updated"));
                }
              });
            }
          });
        }
      })
      .catch(async function (err: any) {
        gettingDataSpinner.stop();
        if (err.response.status === 401) {
          refresh();
        }
      });
  }
}

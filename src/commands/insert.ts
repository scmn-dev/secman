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
import chalk from "chalk";
import { readDataFile } from "../../app/config";
import {
  CCFields,
  EmailFields,
  LoginFields,
  NoteFields,
  ServerFields,
} from "../../contents/types";
import { refresh } from "../../app/refresher";
import { InsertExamples } from "../../contents/examples/insert";
import { Flags } from "../../tools/flags";
const prompts = require("prompts");
prompts.override(require("yargs").argv);

export default class Insert extends Command {
  static description = "Insert a password to your vault.";

  static aliases = ["create", "new", "write"];

  static flags = {
    help: flags.help({ char: "h" }),
    logins: flags.boolean({
      char: "l",
      description: "insert a password from logins type.",
    }),
    "credit-cards": flags.boolean({
      char: "c",
      description: "insert a password from credit cards type.",
    }),
    emails: flags.boolean({
      char: "e",
      description: "insert a password from emails type.",
    }),
    notes: flags.boolean({
      char: "n",
      description: "insert a password from notes type.",
    }),
    servers: flags.boolean({
      char: "s",
      description: "insert a password from servers type.",
    }),
  };

  static examples = InsertExamples;

  async run() {
    const { flags } = this.parse(Insert);
    let API_URL = "/api";
    const access_token = readDataFile("access_token");
    let form: any = {};

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

    if (flags.logins) {
      API_URL += "/logins";

      for (let i = 0; i < LoginFields.length; i++) {
        const loga = await prompts({
          type: LoginFields[i].title === "Password" ? "password" : "text",
          name: "value",
          message: LoginFields[i].title,
        });

        form[LoginFields[i].value] = loga.value;
      }
    } else if (flags["credit-cards"]) {
      API_URL += "/credit-cards";

      for (let i = 0; i < CCFields.length; i++) {
        const cca = await prompts({
          type:
            CCFields[i].title === "Verification Number" ? "password" : "text",
          name: "value",
          message: CCFields[i].title,
        });

        form[CCFields[i].value] = cca.value;
      }
    } else if (flags.emails) {
      API_URL += "/emails";

      for (let i = 0; i < EmailFields.length; i++) {
        const ema = await prompts({
          type: EmailFields[i].title === "Password" ? "password" : "text",
          name: "value",
          message: EmailFields[i].title,
        });

        form[EmailFields[i].value] = ema.value;
      }
    } else if (flags.notes) {
      API_URL += "/notes";

      for (let i = 0; i < NoteFields.length; i++) {
        const not = await prompts({
          type: "text",
          name: "value",
          message: NoteFields[i].title,
        });

        form[NoteFields[i].value] = not.value;
      }
    } else if (flags.servers) {
      API_URL += "/servers";
      for (let i = 0; i < ServerFields.length; i++) {
        const passwords = () => {
          if (
            ServerFields[i].title === "Password" ||
            ServerFields[i].title === "Hosting Password" ||
            ServerFields[i].title === "Admin Password"
          ) {
            return "password";
          } else {
            return "text";
          }
        };

        const ser = await prompts({
          type: passwords(),
          name: "value",
          message: ServerFields[i].title,
        });

        form[ServerFields[i].value] = ser.value;
      }
    } else {
      this.error("Incorrect type of entry.");
    }

    const payload = CryptoTools.encryptPayload(
      form,
      enc_fields(),
      readDataFile("master_password_hash"),
      readDataFile("transmission_key")
    );

    await API.post(API_URL, payload, {
      headers: {
        Authorization: `Bearer ${access_token}`,
        Cookie: `secman_token=${access_token}`,
      },
    })
      .then(async (res: any) => {
        if (res.status === 200 || res.status === 202) {
          console.log(chalk.green("Password created"));
        } else {
          console.log(chalk.red("Password not created"));
        }
      })
      .catch((err: any) => {
        if (err.response.status === 401) {
          refresh(`insert ${Flags(flags)}`);
        }
      });
  }
}

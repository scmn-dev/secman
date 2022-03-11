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
import { readData } from "../../config";
import { InsertExamples } from "../../contents/examples/insert";
import { error, success } from "../../design/layout";

export default class Insert extends Command {
  static description = "Insert a password to your vault.";

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
    title: flags.string({
      char: "t",
      description: "the title of the password.",
    }),
    url: flags.string({
      char: "u",
      description: "the url of the password.",
    }),
    username: flags.string({
      char: "U",
      description: "the username of the password.",
    }),
    password: flags.string({
      char: "p",
      description: "the password of the password.",
    }),
    extra: flags.string({
      char: "x",
      description: "the extra of the password.",
    }),
    cardholder_name: flags.string({
      char: "C",
      description: "the cardholder name of the password.",
    }),
    type: flags.string({
      char: "T",
      description: "the type of the password.",
    }),
    number: flags.string({
      char: "N",
      description: "the number of the password.",
    }),
    expiry_date: flags.string({
      char: "E",
      description: "the expiry date of the password.",
    }),
    verification_number: flags.string({
      char: "V",
      description: "the verification number of the password.",
    }),
    email: flags.string({
      char: "m",
      description: "the email of the password.",
    }),
    note: flags.string({
      char: "o",
      description: "the note of the password.",
    }),
    ip: flags.string({
      char: "i",
      description: "the ip of the password.",
    }),
    hosting_username: flags.string({
      char: "H",
      description: "the hosting username of the password.",
    }),
    hosting_password: flags.string({
      char: "O",
      description: "the hosting password of the password.",
    }),
    admin_username: flags.string({
      char: "a",
      description: "the admin username of the password.",
    }),
    admin_password: flags.string({
      char: "w",
      description: "the admin password of the password.",
    }),
  };

  static examples = InsertExamples;

  async run() {
    const { flags } = this.parse(Insert);
    let API_URL = "/api";
    const access_token = readData("access_token");
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

    await API.get(API_URL, {
      headers: {
        Authorization: `Bearer ${access_token}`,
        Cookie: `secman_token=${access_token}`,
      },
    })
      .then(async (res: any) => {
        if (res.status === 200 || res.status === 202) {
          if (flags.logins) {
            form["title"] = flags.title || "";
            form["url"] = flags.url || "";
            form["username"] = flags.username || "";
            form["password"] = flags.password || "";
            form["extra"] = flags.extra || "";
          } else if (flags["credit-cards"]) {
            form["title"] = flags.title || "";
            form["cardholder_name"] = flags.cardholder_name || "";
            form["type"] = flags.type || "";
            form["number"] = flags.number || "";
            form["expiry_date"] = flags.expiry_date || "";
            form["verification_number"] = flags.verification_number || "";
          } else if (flags.emails) {
            form["title"] = flags.title || "";
            form["email"] = flags.email || "";
            form["password"] = flags.password || "";
          } else if (flags.notes) {
            form["title"] = flags.title || "";
            form["note"] = flags.note || "";
          } else if (flags.servers) {
            form["title"] = flags.title || "";
            form["ip"] = flags.ip || "";
            form["username"] = flags.username || "";
            form["password"] = flags.password || "";
            form["url"] = flags.url || "";
            form["hosting_username"] = flags.hosting_username || "";
            form["hosting_password"] = flags.hosting_password || "";
            form["admin_username"] = flags.admin_username || "";
            form["admin_password"] = flags.admin_password || "";
            form["extra"] = flags.extra || "";
          } else {
            this.error("Incorrect type of entry.");
          }

          const payload = CryptoTools.encryptPayload(
            form,
            enc_fields(),
            readData("master_password_hash"),
            readData("transmission_key")
          );

          await API.post(API_URL, payload, {
            headers: {
              Authorization: `Bearer ${access_token}`,
              Cookie: `secman_token=${access_token}`,
            },
          })
            .then(async (res: any) => {
              if (res.status === 200 || res.status === 202) {
                console.log("200");
              } else {
                console.log("404");
              }
            })
            .catch((err: any) => {
              if (err.response.status === 401) {
                console.log("401");
              }
            });
        }
      })
      .catch((err: any) => {
        console.log(err.response.status);
      });
  }
}

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
import cryptojs from "crypto-js";
import { ReadExamples } from "../../contents/examples/read";
import { spinner } from "@secman/spinner";
import { readData } from "../../config";
const prettier = require("prettier");

export default class Read extends Command {
  static description = "Print the password of a secman entry.";

  static aliases = ["show", "print"];

  static flags = {
    help: flags.help({ char: "h" }),
    logins: flags.boolean({
      char: "l",
      description: "read password from logins type",
    }),
    "credit-cards": flags.boolean({
      char: "c",
      description: "read password from credit cards type",
    }),
    emails: flags.boolean({
      char: "e",
      description: "read password from emails type",
    }),
    notes: flags.boolean({
      char: "n",
      description: "read password from notes type",
    }),
    servers: flags.boolean({
      char: "s",
      description: "read password from servers type",
    }),
    "show-password": flags.boolean({ char: "p", description: "show password" }),
  };

  static args = [{ name: "PASSWORD_NAME" }];

  static examples = ReadExamples;

  async run() {
    const { args, flags } = this.parse(Read);
    let API_URL = "/api";
    const access_token = readData("access_token");

    if (!args.PASSWORD_NAME) {
      this.error("PASSWORD NAME is not found");
    }

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
        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (res.status === 200 || res.status === 202) {
          gettingDataSpinner.stop();

          const ms_hash = readData("master_password_hash");

          itemList.forEach((element: any) => {
            if (element.title === args.PASSWORD_NAME) {
              if (flags.logins) {
                CryptoTools.decryptFields(
                  element,
                  LOGINS_ENCRYPTED_FIELDS,
                  ms_hash
                );

                const checkExtra = element.extra ? element.extra : "No extra";
                const url = element.url.startsWith("http")
                  ? element.url
                  : "https://" + element.url;

                const password = flags["show-password"]
                  ? element.password
                  : "*".repeat(element.password.length);

                const data = {
                  title: element.title,
                  url: url,
                  username: element.username,
                  password: password,
                  extra: checkExtra,
                };

                const formattedData = prettier.format(JSON.stringify(data), {
                  parser: "json",
                  semi: false,
                });

                const fd = formattedData.slice(0, formattedData.length - 1);

                console.log(fd);
              } else if (flags["credit-cards"]) {
                CryptoTools.decryptFields(
                  element,
                  CCS_ENCRYPTED_FIELDS,
                  ms_hash
                );

                const verification_number = flags["show-password"]
                  ? element.verification_number
                  : "*".repeat(element.verification_number.length);

                const data = {
                  card_name: element.title,
                  cardholder_name: element.cardholder_name,
                  type: element.type,
                  number: element.number,
                  expiry_date: element.expiry_date,
                  verification_number: verification_number,
                };

                const formattedData = prettier.format(JSON.stringify(data), {
                  parser: "json",
                  semi: false,
                });

                const fd = formattedData.slice(0, formattedData.length - 1);

                console.log(fd);
              } else if (flags.emails) {
                CryptoTools.decryptFields(
                  element,
                  EMAILS_ENCRYPTED_FIELDS,
                  ms_hash
                );

                const password = flags["show-password"]
                  ? element.password
                  : "*".repeat(element.password.length);

                const data = {
                  email: element.email,
                  password: password,
                };

                const formattedData = prettier.format(JSON.stringify(data), {
                  parser: "json",
                  semi: false,
                });

                const fd = formattedData.slice(0, formattedData.length - 1);

                console.log(fd);
              } else if (flags.notes) {
                CryptoTools.decryptFields(
                  element,
                  NOTES_ENCRYPTED_FIELDS,
                  ms_hash
                );
                const note = flags["show-password"]
                  ? element.note
                  : "*".repeat(element.note.length);

                const data = {
                  title: element.title,
                  note: note,
                };

                const formattedData = prettier.format(JSON.stringify(data), {
                  parser: "json",
                  semi: false,
                });

                const fd = formattedData.slice(0, formattedData.length - 1);

                console.log(fd);
              } else if (flags.servers) {
                CryptoTools.decryptFields(
                  element,
                  SERVERS_ENCRYPTED_FIELDS,
                  ms_hash
                );

                const url = element.url.startsWith("http")
                  ? element.url
                  : "https://" + element.url;
                const password = flags["show-password"]
                  ? element.password
                  : "*".repeat(element.password.length);
                const hosting_password = flags["show-password"]
                  ? element.hosting_password
                  : "*".repeat(element.hosting_password.length);
                const admin_password = flags["show-password"]
                  ? element.admin_password
                  : "*".repeat(element.admin_password.length);
                const checkExtra = element.extra ? element.extra : "No extra";

                const data = {
                  title: element.title,
                  ip: element.ip,
                  url: url,
                  username: element.username,
                  password: password,
                  hosting_username: element.hosting_username,
                  hosting_password: hosting_password,
                  admin_username: element.admin_username,
                  admin_password: admin_password,
                  extra: checkExtra,
                };

                const formattedData = prettier.format(JSON.stringify(data), {
                  parser: "json",
                  semi: false,
                });

                const fd = formattedData.slice(0, formattedData.length - 1);

                console.log(fd);
              }
            }
          });
        }
      })
      .catch(async function (err: any) {
        gettingDataSpinner.stop();

        console.log(err.response.status);
      });
  }
}

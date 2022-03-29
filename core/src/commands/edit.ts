import { Command, flags } from "@oclif/command";
import {
  CCS_ENCRYPTED_FIELDS,
  EMAILS_ENCRYPTED_FIELDS,
  LOGINS_ENCRYPTED_FIELDS,
  NOTES_ENCRYPTED_FIELDS,
  SERVERS_ENCRYPTED_FIELDS,
} from "../../constants";
import { API } from "../../contract";
import { CryptoTools } from "@secman/crypto";
import cryptojs from "crypto-js";
import { spinner } from "@secman/spinner";
import { readData } from "../../config";
import { EditExamples } from "../../contents/examples/edit";

export default class Edit extends Command {
  static description = "Update or change a value in a password.";

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
    field: flags.string({
      char: "f",
      description: "field to edit.",
      default: "",
    }),
    value: flags.string({
      char: "v",
      description: "value to add.",
    }),
  };

  static args = [{ name: "PASSWORD_NAME" }];

  static examples = EditExamples;

  async run() {
    const { args, flags } = this.parse(Edit);
    let API_URL = "/api";
    const access_token = readData("access_token");

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
            readData("transmission_key")
          ).toString(cryptojs.enc.Utf8);

          const itemList = JSON.parse(item);

          const ms_hash = readData("master_password_hash");

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

              let f;

              if (flags.field == "Card Name") {
                f = "title";
              } else if (flags.field == "Card Type") {
                f = "type";
              } else if (flags.field == "Card Number") {
                f = "number";
              } else if (flags.field == "Expiry Date") {
                f = "expiry_date";
              } else if (flags.field == "Verification Number") {
                f = "verification_number";
              } else if (flags.field == "Card Holder Name") {
                f = "cardholder_name";
              } else if (flags.field == "Hosting Username") {
                f = "hosting_username";
              } else if (flags.field == "Hosting Password") {
                f = "hosting_password";
              } else if (flags.field == "Admin Username") {
                f = "admin_username";
              } else if (flags.field == "Admin Password") {
                f = "admin_password";
              } else if (flags.field == "Email Address") {
                f = "email";
              } else if (flags.field == "Ip Address") {
                f = "ip";
              } else {
                f = flags.field;
              }

              element[f] = flags.value;

              const payload = CryptoTools.encryptPayload(
                element,
                enc_fields(),
                readData("master_password_hash"),
                readData("transmission_key")
              );

              await API.put(API_URL + `/${element.id}`, payload, {
                headers: {
                  Authorization: `Bearer ${access_token}`,
                  Cookie: `secman_token=${access_token}`,
                },
              }).then(async (res: any) => {
                if (res.status === 200 || res.status === 202) {
                  console.log("200");
                }
              });
            } else {
              console.log("404");
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

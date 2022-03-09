import { Command, flags } from "@oclif/command";
import { API } from "../../contract";
import * as cryptojs from "crypto-js";
import { spinner } from "@secman/spinner";
import { readData } from "../../config";
import { DeleteExamples } from "../../contents/examples/delete";

export default class Delete extends Command {
  static description = "Delete a password from the vault.";

  static flags = {
    help: flags.help({ char: "h" }),
    logins: flags.boolean({
      char: "l",
      description: "delete password from logins type",
    }),
    "credit-cards": flags.boolean({
      char: "c",
      description: "delete password from credit cards type",
    }),
    emails: flags.boolean({
      char: "e",
      description: "delete password from emails type",
    }),
    notes: flags.boolean({
      char: "n",
      description: "delete password from notes type",
    }),
    servers: flags.boolean({
      char: "s",
      description: "delete password from servers type",
    }),
  };

  static args = [{ name: "PASSWORD_NAME" }];

  static examples = DeleteExamples;

  async run() {
    const { args, flags } = this.parse(Delete);
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
        Cookie: `access_token=${access_token}`,
      },
    })
      .then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        itemList.forEach(async (element: any) => {
          if (element.title === args.PASSWORD_NAME) {
            await API.delete(API_URL + `/${element.id}`, {
              headers: {
                Authorization: `Bearer ${access_token}`,
                Cookie: `access_token=${access_token}`,
              },
            }).then(async (res: any) => {
              console.log(res.status);
            });
          }
        });
      })
      .catch((err: any) => {
        gettingDataSpinner.stop();

        console.log(err.response.status);
      });
  }
}

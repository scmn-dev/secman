import { Command, flags } from "@oclif/command";
import { API } from "../../contract";
import * as cryptojs from "crypto-js";
import { spinner } from "@secman/spinner";
import { readDataFile } from "../../app/config";
import { refresh } from "../../app/refresher";
import { DeleteExamples } from "../../contents/examples/delete";
const prompts = require("prompts");
prompts.override(require("yargs").argv);
import { Types } from "../../tools/flags";
import { error, success } from "../../design/layout";

export default class Delete extends Command {
  static description = "Delete a password from the vault.";

  static aliases = ["remove", "rm", "del"];

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
        Cookie: `access_token=${access_token}`,
      },
    })
      .then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        itemList.forEach(async (element: any) => {
          if (element.title === args.PASSWORD_NAME) {
            const confirm = await prompts({
              type: "toggle",
              name: "yes_or_no",
              message: `Are you sure you want to delete ${args.PASSWORD_NAME}?`,
              initial: "no",
              active: "yes",
              inactive: "no",
            });

            if (confirm.yes_or_no) {
              await API.delete(API_URL + `/${element.id}`, {
                headers: {
                  Authorization: `Bearer ${access_token}`,
                  Cookie: `access_token=${access_token}`,
                },
              }).then(async (res: any) => {
                if (res.status === 200) {
                  console.log(
                    success(`${element.title} was deleted successfully.`)
                  );
                } else {
                  console.log(
                    error(`Failed to delete ${element.title}, try again`)
                  );
                }
              });
            } else {
              console.log("Deletion canceled.");
            }
          }
        });
      })
      .catch((err: any) => {
        gettingDataSpinner.stop();
        if (err.response.status === 401) {
          refresh(`delete -${Types(flags)} ${args.PASSWORD_NAME}`);
        } else if (err.response.status === 404) {
          console.log(error("No data found"));
        } else {
          console.log(error("Something went wrong"));
        }
      });
  }
}

import { Command, flags } from "@oclif/command";
import { API } from "../../contract";
import cryptojs from "crypto-js";
import { spinner } from "@secman/spinner";
import { readDataFile } from "../../app/config";
import { refresh } from "../../app/refresher";
import { ListExamples } from "../../contents/examples/list";
import { Types } from "../../tools/flags";
import { bold, error } from "../../design/layout";

export default class List extends Command {
  static description = "List all passwords.";

  static aliases = [".", "ls"];

  static flags = {
    help: flags.help({ char: "h" }),
    logins: flags.boolean({
      char: "l",
      description: "list passwords from logins type",
    }),
    "credit-cards": flags.boolean({
      char: "c",
      description: "list passwords from credit cards type",
    }),
    emails: flags.boolean({
      char: "e",
      description: "list passwords from emails type",
    }),
    notes: flags.boolean({
      char: "n",
      description: "list passwords from notes type",
    }),
    servers: flags.boolean({
      char: "s",
      description: "list passwords from servers type",
    }),
  };

  static examples = ListExamples;

  async run() {
    const { flags } = this.parse(List);
    const access_token = readDataFile("access_token");

    const gettingDataSpinner = spinner("📡 Getting data...").start();

    const logins = async () => {
      await API.get("/api/logins", {
        headers: {
          Authorization: `Bearer ${access_token}`,
          Cookie: `secman_token=${access_token}`,
        },
      }).then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        console.log(`.
├──Logins`);

        itemList.forEach((element: any) => {
          if (itemList.indexOf(element) === itemList.length - 1) {
            // console.log(`│  └──${element.title}`);
            if (flags.logins) {
              console.log(`└──┴──${element.title}`);
            } else {
              console.log(`│  └──${element.title}`);
            }
          } else {
            console.log(`│  ├──${element.title}`);
          }
        });
      });
    };

    const creditCards = async () => {
      await API.get("/api/credit-cards", {
        headers: {
          Authorization: `Bearer ${access_token}`,
          Cookie: `secman_token=${access_token}`,
        },
      }).then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        //   console.log(`├──Credit Cards`);
        if (flags["credit-cards"]) {
          console.log(`.
├──Credit Cards`);
        } else {
          console.log(`├──Credit Cards`);
        }

        itemList.forEach((element: any) => {
          if (itemList.indexOf(element) === itemList.length - 1) {
            if (flags["credit-cards"]) {
              console.log(`└──┴──${element.title}`);
            } else {
              console.log(`│  └──${element.title}`);
            }
          } else {
            console.log(`│  ├──${element.title}`);
          }
        });
      });
    };

    const emails = async () => {
      await API.get("/api/emails", {
        headers: {
          Authorization: `Bearer ${access_token}`,
          Cookie: `secman_token=${access_token}`,
        },
      }).then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        // console.log(`├──Emails`);
        if (flags.emails) {
          console.log(`.
├──${bold("Emails")}`);
        } else {
          console.log(`├──Emails`);
        }

        itemList.forEach((element: any) => {
          if (itemList.indexOf(element) === itemList.length - 1) {
            // console.log(`│  └──${element.title}`);
            if (flags.emails) {
              console.log(`└──┴──${element.title}`);
            } else {
              console.log(`│  └──${element.title}`);
            }
          } else {
            console.log(`│  ├──${element.title}`);
          }
        });
      });
    };

    const notes = async () => {
      await API.get("/api/notes", {
        headers: {
          Authorization: `Bearer ${access_token}`,
          Cookie: `secman_token=${access_token}`,
        },
      }).then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        // console.log(`├──Notes`);
        if (flags.notes) {
          console.log(`.
├──${bold("Notes")}`);
        } else {
          console.log(`├──Notes`);
        }

        itemList.forEach((element: any) => {
          if (itemList.indexOf(element) === itemList.length - 1) {
            if (flags.notes) {
              console.log(`└──┴──${element.title}`);
            } else {
              console.log(`│  └──${element.title}`);
            }
          } else {
            console.log(`│  ├──${element.title}`);
          }
        });
      });
    };

    const servers = async () => {
      await API.get("/api/servers", {
        headers: {
          Authorization: `Bearer ${access_token}`,
          Cookie: `secman_token=${access_token}`,
        },
      }).then(async (res: any) => {
        gettingDataSpinner.stop();

        const item: any = cryptojs.AES.decrypt(
          res.data.data,
          readDataFile("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        // console.log(`├──Servers`);
        if (flags.servers) {
          console.log(`.
├──${bold("Servers")}`);
        } else {
          console.log(`├──Servers`);
        }

        itemList.forEach((element: any) => {
          if (itemList.indexOf(element) === itemList.length - 1) {
            console.log(`└──┴──${element.title}`);
          } else {
            console.log(`│  ├──${element.title}`);
          }
        });
      });
    };

    const core = async () => {
      logins()
        .then(() => {
          creditCards().then(() => {
            emails().then(() => {
              notes().then(() => {
                servers();
              });
            });
          });
        })
        .catch((err: any) => {
          catcher(err);
        });
    };

    const whatIsCommand = () => {
      if (
        flags.logins ||
        flags["credit-cards"] ||
        flags.emails ||
        flags.notes ||
        flags.servers
      ) {
        return `. -${Types(flags)}`;
      } else {
        return ".";
      }
    };

    const catcher = async (err: any) => {
      gettingDataSpinner.stop();
      if (err.response.status === 401) {
        refresh(whatIsCommand());
      } else if (err.response.status === 404) {
        console.log(error("No data found"));
      } else {
        console.log(error("Something went wrong"));
      }
    };

    if (flags.logins) {
      logins().catch((err: any) => {
        catcher(err);
      });
    } else if (flags["credit-cards"]) {
      creditCards().catch((err: any) => {
        catcher(err);
      });
    } else if (flags.emails) {
      emails().catch((err: any) => {
        catcher(err);
      });
    } else if (flags.notes) {
      notes().catch((err: any) => {
        catcher(err);
      });
    } else if (flags.servers) {
      servers().catch((err: any) => {
        catcher(err);
      });
    } else {
      core();
    }
  }
}

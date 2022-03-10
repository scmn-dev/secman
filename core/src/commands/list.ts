import { Command, flags } from "@oclif/command";
import { API } from "../../contract";
import cryptojs from "crypto-js";
import { spinner } from "@secman/spinner";
import { readData } from "../../config";
import { ListExamples } from "../../contents/examples/list";
import { bold } from "../../design/layout";
import {
  CCS_ENCRYPTED_FIELDS,
  EMAILS_ENCRYPTED_FIELDS,
  LOGINS_ENCRYPTED_FIELDS,
  NOTES_ENCRYPTED_FIELDS,
  SERVERS_ENCRYPTED_FIELDS,
} from "../../constants";
import { CryptoTools } from "../../tools/crypto";

export default class List extends Command {
  static description = "List all passwords.";

  static aliases = ["."];

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
    status: flags.boolean({
      char: "u",
      description: "",
    }),
    tree: flags.boolean({
      char: "t",
      description: "list password in tree view.",
    })
  };

  static examples = ListExamples;

  async run() {
    const { flags } = this.parse(List);
    const access_token = readData("access_token");
    const ms_hash = readData("master_password_hash");

    let login;
    let credit_card;
    let email;
    let note;
    let server;

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
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (flags.status) {
          console.log(res.data.status);
        } else {
          if (flags.tree) {
            console.log(`.
├──Logins`);
            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(
                element,
                LOGINS_ENCRYPTED_FIELDS,
                ms_hash
              );

              if (itemList.indexOf(element) === itemList.length - 1) {
                if (flags.logins) {
                  console.log(`└──┴──${element.title}`);
                } else {
                  console.log(`│  └──${element.title}`);
                }
              } else {
                console.log(`│  ├──${element.title}`);
              }
            });
          } else {
            if (flags.logins) {
              pws();
            }

            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(
                element,
                LOGINS_ENCRYPTED_FIELDS,
                ms_hash
              );

              const pw = "*".repeat(element.password.length);

              login = `  ${element.title} -|- ${element.url} -|- ${element.username} -|- ${pw} -|- ${element.extra}`;

              console.log(login);
            });
          }
        }
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
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (flags.status) {
          console.log(res.data.status);
        } else {
          if (flags.tree) {
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
          } else {
            if (flags["credit-cards"]) {
              pws();
            }

            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(element, CCS_ENCRYPTED_FIELDS, ms_hash);

              const vn = "*".repeat(element.verification_number.length);

              credit_card = `  ${element.title} -|- ${element.cardholder_name} -|- ${element.type} -|- ${element.number} -|- ${element.expiry_date} -|- ${vn}`;

              console.log(credit_card);
            });
          }
        }
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
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (flags.status) {
          console.log(res.data.status);
        } else {
          if (flags.tree) {
            if (flags.emails) {
              console.log(`.
├──${bold("Emails")}`);
            } else {
              console.log(`├──Emails`);
            }

            itemList.forEach((element: any) => {
              if (itemList.indexOf(element) === itemList.length - 1) {
                if (flags.emails) {
                  console.log(`└──┴──${element.title}`);
                } else {
                  console.log(`│  └──${element.title}`);
                }
              } else {
                console.log(`│  ├──${element.title}`);
              }
            });
          } else {
            if (flags.emails) {
              pws();
            }

            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(
                element,
                EMAILS_ENCRYPTED_FIELDS,
                ms_hash
              );

              const pw = "*".repeat(element.password.length);

              email = `  ${element.title} -|- ${element.email} -|- ${pw}`;

              console.log(email);
            });
          }
        }
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
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (flags.status) {
          console.log(res.data.status);
        } else {
          if (flags.tree) {
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
          } else {
            if (flags.notes) {
              pws();
            }

            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(
                element,
                NOTES_ENCRYPTED_FIELDS,
                ms_hash
              );

              const n = "*".repeat(element.note.length);
              note = `  ${element.title} -|- ${n} -|- ${element.note.replace(
                /\n/g,
                " "
              )}`;

              console.log(note);
            });
          }
        }
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
          readData("transmission_key")
        ).toString(cryptojs.enc.Utf8);

        const itemList = JSON.parse(item);

        if (flags.status) {
          console.log(res.data.status);
        } else {
          if (flags.tree) {
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
          } else {
            if (flags.servers) {
              pws();
            }

            itemList.forEach((element: any) => {
              CryptoTools.decryptFields(
                element,
                SERVERS_ENCRYPTED_FIELDS,
                ms_hash
              );

              const pw = "*".repeat(element.password.length);
              const hp = "*".repeat(element.hosting_password.length);
              const ap = "*".repeat(element.admin_password.length);

              server = `  ${element.title} -|- ${element.ip} -|- ${element.url} -|- ${element.username} -|- ${pw} -|- ${element.hosting_username} -|- ${hp} -|- ${element.admin_username} -|- ${ap} -|- ${element.extra}`;

              console.log(server);
            });
          }
        }
      });
    };

    const pws = async () => {
      if (!flags.tree) {
        gettingDataSpinner.stop();
        console.log("passwords: |-");
      }
    };

    const core = async () => {
      pws()
        .then(() => {
          logins().then(() => {
            creditCards().then(() => {
              emails().then(() => {
                notes().then(() => {
                  servers();
                });
              });
            });
          });
        })
        .catch((err: any) => {
          catcher(err);
        });
    };

    const catcher = async (err: any) => {
      gettingDataSpinner.stop();

      console.log(err.response.status);
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

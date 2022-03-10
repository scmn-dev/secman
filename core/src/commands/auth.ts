import { Command, flags } from "@oclif/command";
import writeConfigFile, { readConfig } from "../../config";
import { API } from "../../contract";
import { CryptoTools } from "../../tools/crypto";

export default class Auth extends Command {
  static description = "Manage secman's authentication state.";

  static flags = {
    help: flags.help({ char: "h" }),
    email: flags.string({
      char: "e",
      description: "Email address of the account to use.",
    }),
    masterPassword: flags.string({
      char: "m",
      description: "Master password of the account to use.",
    }),
  };

  async run() {
    const { flags } = this.parse(Auth);

    const _ = async () => {
      const email = flags.email;
      let master_password = flags.masterPassword;

      try {
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

              console.log("200");
            })
            .catch(function (err: any) {
              console.log(err.response.status);
            });
        }
      } catch (error) {
        console.log(error);
      }
    };

    switch (true) {
      default:
        const user = readConfig("name");

        if (user != "") {
          console.log("406");
        }

        _();
    }
  }
}

import sendgrid from "@sendgrid/mail";
import { SENDGRID_API_KEY } from "../../constants";
import { SMEmail } from "./mail";

sendgrid.setApiKey(SENDGRID_API_KEY);

export default function SendEmail(
  email: string,
  generatedPassword: string,
  func: any
) {
  const msg = {
    to: email,
    from: "verify@secman.dev",
    subject: `Secman Login Verification (security code: "${generatedPassword}")`,
    text: "Verify with Secman CLI",
    html: SMEmail(email, generatedPassword),
  };

  try {
    sendgrid.send(msg).then((response) => {
      if (response[0].statusCode === 202 || response[0].statusCode === 200) {
        func();
      }
    });
  } catch (error) {
    console.log(error);
  }
}

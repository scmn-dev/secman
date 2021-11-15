export function Types(flags: any) {
  switch (true) {
    case flags.logins:
      return "l";

    case flags["credit-cards"]:
      return "c";

    case flags.emails:
      return "e";

    case flags.notes:
      return "n";

    case flags.servers:
      return "s";

    default:
      return "";
  }
}

export function Multi(flags: any) {
  switch (true) {
    case flags.multi:
      return "m";

    default:
      return "";
  }
}

export function ShowPassword(flags: any) {
  switch (true) {
    case flags["show-password"]:
      return "p";

    default:
      return "";
  }
}

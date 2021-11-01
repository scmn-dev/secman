export const Flags = (flags: any) => {
  if (flags.logins) {
    return "-l";
  } else if (flags["credit-cards"]) {
    return "-c";
  } else if (flags.emails) {
    return "-e";
  } else if (flags.notes) {
    return "-n";
  } else if (flags.servers) {
    return "-s";
  }
};

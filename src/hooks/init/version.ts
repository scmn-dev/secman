const version = require("../../../package.json").version;

export default async function hook() {
  if (["-v", "-V", "--version", "version"].includes(process.argv[2])) {
    console.log(`secman v${version}
https://github.com/scmn-dev/secman/releases/tag/v${version}`);
    return process.exit(0);
  }
}

const version = require("../../../package.json").version;
const versionDate = require("../../../package.json").versionDate;

export default async function hook() {
  if (["-v", "-V", "--version", "version"].includes(process.argv[2])) {
    console.log(`secman v${version} (${versionDate})
https://github.com/scmn-dev/secman/releases/tag/v${version}`);
    return process.exit(0);
  }
}

import path from "path";

// get "version" and "versionDate" from package.json
const version = require(path.join(process.cwd(), "package.json")).version;
const versionDate = require(path.join(
  process.cwd(),
  "package.json"
)).versionDate;

export default async function hook() {
  if (["-v", "-V", "--version", "version"].includes(process.argv[2])) {
    console.log(`secman v${version} (${versionDate})
https://github.com/scmn-dev/secman/releases/tag/v${version}`);
    return process.exit(0);
  }
}

let version = require("../../../../package.json").version;

export default async function hook() {
  if (["-v", "-V", "--version", "version"].includes(process.argv[2])) {
    console.log("v" + version);

    return process.exit(0);
  }
}

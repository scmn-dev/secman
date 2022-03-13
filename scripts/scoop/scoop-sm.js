const fs = require("fs");
const path = require("path");
const rm = require("rimraf");
const mkdirp = require("mkdirp");
const sh = require("shelljs");
const _crypto = require("crypto");
const { promisify } = require("util");
const { pipeline } = require("stream");

const VERSION_CMD = sh.exec("git describe --abbrev=0 --tags");
const VERSION = VERSION_CMD.replace("\n", "")
  .replace("\r", "")
  .replace("v", "");

const SECMAN_32BIT_URL = `https://github.com/scmn-dev/secman/releases/download/v${VERSION}/secman_windows_v${VERSION}_386.zip`;
const SECMAN_64BIT_URL = `https://github.com/scmn-dev/secman/releases/download/v${VERSION}/secman_windows_v${VERSION}_amd64.zip`;

const ROOT = __dirname;
const DIST_DIR = path.join(ROOT, "..", "..", "dist");
const TEMPLATES = path.join(ROOT, "templates");

async function calculateSHA256(fileName) {
  const hash = _crypto.createHash("sha256");

  hash.setEncoding("hex");
  await promisify(pipeline)(fs.createReadStream(fileName), hash);

  return hash.read();
}

async function updateSecmanScoop(secmanDir) {
  const templatePath = path.join(TEMPLATES, "secman.json");
  const template = fs.readFileSync(templatePath).toString("utf-8");

  const SM_32BIT_FILE = path.join(
    DIST_DIR,
    `secman_windows_v${VERSION}_386.zip`
  );
  const SM_64BIT_FILE = path.join(
    DIST_DIR,
    `secman_windows_v${VERSION}_amd64.zip`
  );

  const SM_32BIT_HASH = await calculateSHA256(SM_32BIT_FILE);
  const SM_64BIT_HASH = await calculateSHA256(SM_64BIT_FILE);

  const templateReplaced = template
    .replace("CLI_VERSION", VERSION)
    .replace("SECMAN_32BIT_URL", SECMAN_32BIT_URL)
    .replace("SECMAN_64BIT_URL", SECMAN_64BIT_URL)
    .replace("32BIT_HASH", SM_32BIT_HASH)
    .replace("64BIT_HASH", SM_64BIT_HASH);

  fs.writeFileSync(path.join(secmanDir, "secman.json"), templateReplaced);
}

async function updateScoop() {
  const tmp = path.join(__dirname, "tmp");
  const scoopDir = path.join(tmp, "scoop");

  mkdirp.sync(tmp);
  rm.sync(scoopDir);

  console.log(`cloning https://github.com/scmn-dev/scoop to ${scoopDir}`);

  sh.exec(`git clone https://github.com/scmn-dev/scoop.git ${scoopDir}`);

  console.log(`done cloning scmn-dev/scoop to ${scoopDir}`);

  await updateSecmanScoop(scoopDir);
}

updateScoop().catch((err) => {
  console.error(`error running scripts/scoop/scoop-sm.js`, err);
  process.exit(1);
});

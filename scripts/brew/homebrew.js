#!/usr/bin/env node
// This code is forked from https://github.com/heroku/cli/blob/master/scripts/release/homebrew.js

const fs = require("fs");
const execa = require("execa");
const https = require("https");
const path = require("path");
const rm = require("rimraf");
const mkdirp = require("mkdirp");
const { promisify } = require("util");
const { pipeline } = require("stream");
const crypto = require("crypto");

const NODE_JS_BASE = "https://nodejs.org/download/release";
const SECMAN_DIR = path.join(__dirname, "..", "..");
const DIST_DIR = path.join(SECMAN_DIR, "dist");
const PJSON = require(path.join(SECMAN_DIR, "package.json"));
const NODE_VERSION = PJSON.oclif.update.node.version;
const VERSION = PJSON.version;

async function getText(url) {
  return new Promise((resolve, reject) => {
    https
      .get(url, (res) => {
        let buffer = [];

        res.on("data", (buf) => {
          buffer.push(buf);
        });

        res.on("close", () => {
          resolve(Buffer.concat(buffer).toString("utf-8"));
        });
      })
      .on("error", reject);
  });
}

async function getDownloadInfoForNodeVersion(version) {
  // https://nodejs.org/download/release/v12.21.0/SHASUMS256.txt
  const url = `${NODE_JS_BASE}/v${version}/SHASUMS256.txt`;
  const shasums = await getText(url);
  const shasumLine = shasums.split("\n").find((line) => {
    return line.includes(`node-v${version}-darwin-x64.tar.xz`);
  });

  if (!shasumLine) {
    throw new Error(`could not find matching shasum for ${version}`);
  }

  const [shasum, filename] = shasumLine.trim().split(/\s+/);
  return {
    url: `${NODE_JS_BASE}/v${version}/${filename}`,
    sha256: shasum,
  };
}

if (!process.env.CIRCLE_TAG) {
  console.log("Not on stable release; skipping releasing homebrew");
  process.exit(0);
}

async function calculateSHA256(fileName) {
  const hash = crypto.createHash("sha256");
  hash.setEncoding("hex");
  await promisify(pipeline)(fs.createReadStream(fileName), hash);
  return hash.read();
}

const ROOT = path.join(__dirname, "homebrew");
const TEMPLATES = path.join(ROOT, "templates");

const CLI_URL = "https://cli-files.secman.dev";

async function updateSecmanFormula(brewDir) {
  const templatePath = path.join(TEMPLATES, "secman.rb");
  const template = fs.readFileSync(templatePath).toString("utf-8");

  const pathToDist = path.join(
    DIST_DIR,
    `secman-v${VERSION}`,
    `secman-v${VERSION}.tar.xz`
  );
  const sha256 = await calculateSHA256(pathToDist);
  const url = `${CLI_URL}/secman-v${VERSION}/secman-v${VERSION}.tar.xz`;

  const templateReplaced = template
    .replace("__CLI_DOWNLOAD_URL__", url)
    .replace("__CLI_SHA256__", sha256)
    .replace("__NODE_VERSION__", NODE_VERSION);

  fs.writeFileSync(
    path.join(brewDir, "Formula", "secman.rb"),
    templateReplaced
  );
}

async function updateSecmanNodeFormula(brewDir) {
  const formulaPath = path.join(brewDir, "Formula", "sm-node.rb");

  console.log(`updating sm-node Formula in ${formulaPath}`);
  console.log(`getting SHA and URL for Node.js version ${NODE_VERSION}`);

  const { url, sha256 } = await getDownloadInfoForNodeVersion(NODE_VERSION);

  console.log(
    `done getting SHA for Node.js version ${NODE_VERSION}: ${sha256}`
  );
  console.log(`done getting URL for Node.js version ${NODE_VERSION}: ${url}`);

  const templatePath = path.join(TEMPLATES, "sm-node.rb");
  const template = fs.readFileSync(templatePath).toString("utf-8");

  const templateReplaced = template
    .replace("__NODE_BIN_URL__", url)
    .replace("__NODE_SHA256__", sha256)
    .replace("__NODE_VERSION__", NODE_VERSION);

  fs.writeFileSync(formulaPath, templateReplaced);
  console.log(`done updating sm-node Formula in ${formulaPath}`);
}

async function updateHomebrew() {
  const tmp = path.join(__dirname, "tmp");
  const homebrewDir = path.join(tmp, "homebrew-brew");
  mkdirp.sync(tmp);
  rm.sync(homebrewDir);

  console.log(
    `cloning https://github.com/scmn-dev/homebrew-secman to ${homebrewDir}`
  );

  await execa("git", [
    "clone",
    "git@github.com:scmn-dev/homebrew-secman.git",
    homebrewDir,
  ]);

  console.log(`done cloning scmn-dev/homebrew-secman to ${homebrewDir}`);

  console.log("updating local git...");
  await updateSecmanNodeFormula(homebrewDir);
  await updateSecmanFormula(homebrewDir);

  const git = async (args, opts = {}) => {
    await execa("git", ["-C", homebrewDir, ...args], opts);
  };

  await git(["add", "Formula"]);
  await git(["config", "--local", "core.pager", "cat"]);
  await git(["diff", "--cached"], { stdio: "inherit" });
  await git(["commit", "-m", `secman v${VERSION}`]);
  if (process.env.SKIP_GIT_PUSH === undefined) {
    await git(["push", "origin", "master"]);
  }
}

updateHomebrew().catch((err) => {
  console.error(`error running scripts/brew/homebrew.js`, err);
  process.exit(1);
});

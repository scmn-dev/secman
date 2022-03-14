import * as path from "path";
import * as fs from "fs";
import { promises as fsPromises } from "fs";
import { isPlainObject as isPlainObj, sortKeys } from "../sort";
import { detectIndent as detect } from "../strings";
const writeFileAtomic = require("write-file-atomic");

const init = (function_: any, filePath: any, data: any, options: any) => {
  if (!filePath) {
    throw new TypeError("Expected a filepath");
  }

  if (data === undefined) {
    throw new TypeError("Expected data to stringify");
  }

  options = {
    indent: "\t",
    sortKeys: false,
    ...options,
  };

  if (options.sortKeys && isPlainObj(data)) {
    data = sortKeys(data, {
      deep: true,
      compare:
        typeof options.sortKeys === "function" ? options.sortKeys : undefined,
    });
  }

  return function_(filePath, data, options);
};

const main = async (filePath: any, data: any, options: any) => {
  let { indent } = options;
  let trailingNewline = "\n";
  try {
    const file = await fsPromises.readFile(filePath, "utf8");
    if (!file.endsWith("\n")) {
      trailingNewline = "";
    }

    if (options.detectIndent) {
      indent = detect(file).indent;
    }
  } catch (error) {
    throw error;
  }

  const json = JSON.stringify(data, options.replacer, indent);

  return writeFileAtomic(filePath, `${json}${trailingNewline}`, {
    mode: options.mode,
    chown: false,
  });
};

const mainSync = (filePath: any, data: any, options: any) => {
  let { indent } = options;
  let trailingNewline = "\n";
  try {
    const file = fs.readFileSync(filePath, "utf8");
    if (!file.endsWith("\n")) {
      trailingNewline = "";
    }

    if (options.detectIndent) {
      indent = detect(file).indent;
    }
  } catch (error) {
    throw error;
  }

  const json = JSON.stringify(data, options.replacer, indent);

  return writeFileAtomic.sync(filePath, `${json}${trailingNewline}`, {
    mode: options.mode,
    chown: false,
  });
};

export async function writeJsonFile(filePath: any, data: any, options: any) {
  await fsPromises.mkdir(path.dirname(filePath), { recursive: true });
  await init(main, filePath, data, options);
}

export function writeJsonFileSync(filePath: any, data: any, options: any) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
  init(mainSync, filePath, data, options);
}

import emojiRegex from "emoji-regex";
import ansiStyles from "ansi-styles";

export function stripAnsi(string: any) {
  if (typeof string !== "string") {
    throw new TypeError(`Expected a \`string\`, got \`${typeof string}\``);
  }

  return string.replace(ansiRegex(), "");
}

export function indentString(string: any, count = 1, options = {}) {
  const { indent = " ", includeEmptyLines = false }: any = options;

  if (typeof string !== "string") {
    throw new TypeError(
      `Expected \`input\` to be a \`string\`, got \`${typeof string}\``
    );
  }

  if (typeof count !== "number") {
    throw new TypeError(
      `Expected \`count\` to be a \`number\`, got \`${typeof count}\``
    );
  }

  if (count < 0) {
    throw new RangeError(
      `Expected \`count\` to be at least 0, got \`${count}\``
    );
  }

  if (typeof indent !== "string") {
    throw new TypeError(
      `Expected \`options.indent\` to be a \`string\`, got \`${typeof indent}\``
    );
  }

  if (count === 0) {
    return string;
  }

  const regex = includeEmptyLines ? /^/gm : /^(?!\s*$)/gm;

  return string.replace(regex, indent.repeat(count));
}

const INDENT_REGEX = /^(?:( )+|\t+)/;
const INDENT_TYPE_SPACE = "space";
const INDENT_TYPE_TAB = "tab";

function makeIndentsMap(string: any, ignoreSingleSpaces: any) {
  const indents = new Map();
  let previousSize = 0;
  let previousIndentType;
  let key;

  for (const line of string.split(/\n/g)) {
    if (!line) {
      continue;
    }

    let indent;
    let indentType;
    let weight;
    let entry;
    const matches = line.match(INDENT_REGEX);

    if (matches === null) {
      previousSize = 0;
      previousIndentType = "";
    } else {
      indent = matches[0].length;
      indentType = matches[1] ? INDENT_TYPE_SPACE : INDENT_TYPE_TAB;

      if (
        ignoreSingleSpaces &&
        indentType === INDENT_TYPE_SPACE &&
        indent === 1
      ) {
        continue;
      }

      if (indentType !== previousIndentType) {
        previousSize = 0;
      }

      previousIndentType = indentType;

      weight = 0;

      const indentDifference = indent - previousSize;
      previousSize = indent;

      if (indentDifference === 0) {
        weight++;
      } else {
        const absoluteIndentDifference =
          indentDifference > 0 ? indentDifference : -indentDifference;
        key = encodeIndentsKey(indentType, absoluteIndentDifference);
      }

      entry = indents.get(key);
      entry = entry === undefined ? [1, 0] : [++entry[0], entry[1] + weight];

      indents.set(key, entry);
    }
  }

  return indents;
}

function encodeIndentsKey(indentType: any, indentAmount: any) {
  const typeCharacter = indentType === INDENT_TYPE_SPACE ? "s" : "t";
  return typeCharacter + String(indentAmount);
}

function decodeIndentsKey(indentsKey: any) {
  const keyHasTypeSpace = indentsKey[0] === "s";
  const type = keyHasTypeSpace ? INDENT_TYPE_SPACE : INDENT_TYPE_TAB;

  const amount = Number(indentsKey.slice(1));

  return { type, amount };
}

function getMostUsedKey(indents: any) {
  let result;
  let maxUsed = 0;
  let maxWeight = 0;

  for (const [key, [usedCount, weight]] of indents) {
    if (usedCount > maxUsed || (usedCount === maxUsed && weight > maxWeight)) {
      maxUsed = usedCount;
      maxWeight = weight;
      result = key;
    }
  }

  return result;
}

function makeIndentString(type: any, amount: any) {
  const indentCharacter = type === INDENT_TYPE_SPACE ? " " : "\t";
  return indentCharacter.repeat(amount);
}

export function detectIndent(string: any) {
  if (typeof string !== "string") {
    throw new TypeError("Expected a string");
  }

  let indents = makeIndentsMap(string, true);
  if (indents.size === 0) {
    indents = makeIndentsMap(string, false);
  }

  const keyOfMostUsedIndent = getMostUsedKey(indents);

  let type;
  let amount = 0;
  let indent = "";

  if (keyOfMostUsedIndent !== undefined) {
    ({ type, amount } = decodeIndentsKey(keyOfMostUsedIndent));
    indent = makeIndentString(type, amount);
  }

  return {
    amount,
    type,
    indent,
  };
}

export function ansiRegex({ onlyFirst = false } = {}) {
  const pattern = [
    "[\\u001B\\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\\u0007)",
    "(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))",
  ].join("|");

  return new RegExp(pattern, onlyFirst ? undefined : "g");
}

const ESCAPES = new Set(["\u001B", "\u009B"]);
const END_CODE = 39;
const ANSI_ESCAPE_BELL = "\u0007";
const ANSI_CSI = "[";
const ANSI_OSC = "]";
const ANSI_SGR_TERMINATOR = "m";
const ANSI_ESCAPE_LINK = `${ANSI_OSC}8;;`;

const wrapAnsiCode = (code: any) =>
  `${ESCAPES.values().next().value}${ANSI_CSI}${code}${ANSI_SGR_TERMINATOR}`;
const wrapAnsiHyperlink = (uri: any) =>
  `${
    ESCAPES.values().next().value
  }${ANSI_ESCAPE_LINK}${uri}${ANSI_ESCAPE_BELL}`;

const wordLengths = (string: any) =>
  string.split(" ").map((character: any) => stringWidth(character));

const wrapWord = (rows: any, word: any, columns: any) => {
  const characters = [...word];

  let isInsideEscape = false;
  let isInsideLinkEscape = false;
  let visible = stringWidth(stripAnsi(rows[rows.length - 1]));

  for (const [index, character] of characters.entries()) {
    const characterLength = stringWidth(character);

    if (visible + characterLength <= columns) {
      rows[rows.length - 1] += character;
    } else {
      rows.push(character);
      visible = 0;
    }

    if (ESCAPES.has(character)) {
      isInsideEscape = true;
      isInsideLinkEscape = characters
        .slice(index + 1)
        .join("")
        .startsWith(ANSI_ESCAPE_LINK);
    }

    if (isInsideEscape) {
      if (isInsideLinkEscape) {
        if (character === ANSI_ESCAPE_BELL) {
          isInsideEscape = false;
          isInsideLinkEscape = false;
        }
      } else if (character === ANSI_SGR_TERMINATOR) {
        isInsideEscape = false;
      }

      continue;
    }

    visible += characterLength;

    if (visible === columns && index < characters.length - 1) {
      rows.push("");
      visible = 0;
    }
  }

  if (!visible && rows[rows.length - 1].length > 0 && rows.length > 1) {
    rows[rows.length - 2] += rows.pop();
  }
};

const stringVisibleTrimSpacesRight = (string: any) => {
  const words = string.split(" ");
  let last = words.length;

  while (last > 0) {
    if (stringWidth(words[last - 1]) > 0) {
      break;
    }

    last--;
  }

  if (last === words.length) {
    return string;
  }

  return words.slice(0, last).join(" ") + words.slice(last).join("");
};

const exec = (string: any, columns: any, options: any = {}) => {
  if (options.trim !== false && string.trim() === "") {
    return "";
  }

  let returnValue = "";
  let escapeCode;
  let escapeUrl;

  const lengths = wordLengths(string);
  let rows = [""];

  for (const [index, word] of string.split(" ").entries()) {
    if (options.trim !== false) {
      rows[rows.length - 1] = rows[rows.length - 1].trimStart();
    }

    let rowLength = stringWidth(rows[rows.length - 1]);

    if (index !== 0) {
      if (
        rowLength >= columns &&
        (options.wordWrap === false || options.trim === false)
      ) {
        rows.push("");
        rowLength = 0;
      }

      if (rowLength > 0 || options.trim === false) {
        rows[rows.length - 1] += " ";
        rowLength++;
      }
    }

    if (options.hard && lengths[index] > columns) {
      const remainingColumns = columns - rowLength;
      const breaksStartingThisLine =
        1 + Math.floor((lengths[index] - remainingColumns - 1) / columns);
      const breaksStartingNextLine = Math.floor((lengths[index] - 1) / columns);
      if (breaksStartingNextLine < breaksStartingThisLine) {
        rows.push("");
      }

      wrapWord(rows, word, columns);
      continue;
    }

    if (
      rowLength + lengths[index] > columns &&
      rowLength > 0 &&
      lengths[index] > 0
    ) {
      if (options.wordWrap === false && rowLength < columns) {
        wrapWord(rows, word, columns);
        continue;
      }

      rows.push("");
    }

    if (rowLength + lengths[index] > columns && options.wordWrap === false) {
      wrapWord(rows, word, columns);
      continue;
    }

    rows[rows.length - 1] += word;
  }

  if (options.trim !== false) {
    rows = rows.map((row) => stringVisibleTrimSpacesRight(row));
  }

  const pre = [...rows.join("\n")];

  for (const [index, character] of pre.entries()) {
    returnValue += character;

    if (ESCAPES.has(character)) {
      const { groups }: any = new RegExp(
        `(?:\\${ANSI_CSI}(?<code>\\d+)m|\\${ANSI_ESCAPE_LINK}(?<uri>.*)${ANSI_ESCAPE_BELL})`
      ).exec(pre.slice(index).join("")) || { groups: {} };
      if (groups.code !== undefined) {
        const code = Number.parseFloat(groups.code);
        escapeCode = code === END_CODE ? undefined : code;
      } else if (groups.uri !== undefined) {
        escapeUrl = groups.uri.length === 0 ? undefined : groups.uri;
      }
    }

    const code = ansiStyles.codes.get(Number(escapeCode));

    if (pre[index + 1] === "\n") {
      if (escapeUrl) {
        returnValue += wrapAnsiHyperlink("");
      }

      if (escapeCode && code) {
        returnValue += wrapAnsiCode(code);
      }
    } else if (character === "\n") {
      if (escapeCode && code) {
        returnValue += wrapAnsiCode(escapeCode);
      }

      if (escapeUrl) {
        returnValue += wrapAnsiHyperlink(escapeUrl);
      }
    }
  }

  return returnValue;
};

export function wrapAnsi(string: any, columns: any, options: any) {
  return String(string)
    .normalize()
    .replace(/\r\n/g, "\n")
    .split("\n")
    .map((line) => exec(line, columns, options))
    .join("\n");
}

export function stringWidth(string: any, options: any = {}) {
  if (typeof string !== "string" || string.length === 0) {
    return 0;
  }

  options = {
    ambiguousIsNarrow: true,
    ...options,
  };

  string = stripAnsi(string);

  if (string.length === 0) {
    return 0;
  }

  string = string.replace(emojiRegex(), "  ");

  const ambiguousCharacterWidth = options.ambiguousIsNarrow ? 1 : 2;
  let width = 0;

  for (const character of string) {
    const codePoint = character.codePointAt(0);

    if (codePoint <= 0x1f || (codePoint >= 0x7f && codePoint <= 0x9f)) {
      continue;
    }

    if (codePoint >= 0x300 && codePoint <= 0x36f) {
      continue;
    }

    const code = eastAsianWidth(character);
    switch (code) {
      case "F":
      case "W":
        width += 2;
        break;
      case "A":
        width += ambiguousCharacterWidth;
        break;
      default:
        width += 1;
    }
  }

  return width;
}

const eastAsianWidth = (character: any) => {
  var x = character.charCodeAt(0);
  var y = character.length == 2 ? character.charCodeAt(1) : 0;
  var codePoint = x;
  if (0xd800 <= x && x <= 0xdbff && 0xdc00 <= y && y <= 0xdfff) {
    x &= 0x3ff;
    y &= 0x3ff;
    codePoint = (x << 10) | y;
    codePoint += 0x10000;
  }

  if (
    0x3000 == codePoint ||
    (0xff01 <= codePoint && codePoint <= 0xff60) ||
    (0xffe0 <= codePoint && codePoint <= 0xffe6)
  ) {
    return "F";
  }
  if (
    0x20a9 == codePoint ||
    (0xff61 <= codePoint && codePoint <= 0xffbe) ||
    (0xffc2 <= codePoint && codePoint <= 0xffc7) ||
    (0xffca <= codePoint && codePoint <= 0xffcf) ||
    (0xffd2 <= codePoint && codePoint <= 0xffd7) ||
    (0xffda <= codePoint && codePoint <= 0xffdc) ||
    (0xffe8 <= codePoint && codePoint <= 0xffee)
  ) {
    return "H";
  }
  if (
    (0x1100 <= codePoint && codePoint <= 0x115f) ||
    (0x11a3 <= codePoint && codePoint <= 0x11a7) ||
    (0x11fa <= codePoint && codePoint <= 0x11ff) ||
    (0x2329 <= codePoint && codePoint <= 0x232a) ||
    (0x2e80 <= codePoint && codePoint <= 0x2e99) ||
    (0x2e9b <= codePoint && codePoint <= 0x2ef3) ||
    (0x2f00 <= codePoint && codePoint <= 0x2fd5) ||
    (0x2ff0 <= codePoint && codePoint <= 0x2ffb) ||
    (0x3001 <= codePoint && codePoint <= 0x303e) ||
    (0x3041 <= codePoint && codePoint <= 0x3096) ||
    (0x3099 <= codePoint && codePoint <= 0x30ff) ||
    (0x3105 <= codePoint && codePoint <= 0x312d) ||
    (0x3131 <= codePoint && codePoint <= 0x318e) ||
    (0x3190 <= codePoint && codePoint <= 0x31ba) ||
    (0x31c0 <= codePoint && codePoint <= 0x31e3) ||
    (0x31f0 <= codePoint && codePoint <= 0x321e) ||
    (0x3220 <= codePoint && codePoint <= 0x3247) ||
    (0x3250 <= codePoint && codePoint <= 0x32fe) ||
    (0x3300 <= codePoint && codePoint <= 0x4dbf) ||
    (0x4e00 <= codePoint && codePoint <= 0xa48c) ||
    (0xa490 <= codePoint && codePoint <= 0xa4c6) ||
    (0xa960 <= codePoint && codePoint <= 0xa97c) ||
    (0xac00 <= codePoint && codePoint <= 0xd7a3) ||
    (0xd7b0 <= codePoint && codePoint <= 0xd7c6) ||
    (0xd7cb <= codePoint && codePoint <= 0xd7fb) ||
    (0xf900 <= codePoint && codePoint <= 0xfaff) ||
    (0xfe10 <= codePoint && codePoint <= 0xfe19) ||
    (0xfe30 <= codePoint && codePoint <= 0xfe52) ||
    (0xfe54 <= codePoint && codePoint <= 0xfe66) ||
    (0xfe68 <= codePoint && codePoint <= 0xfe6b) ||
    (0x1b000 <= codePoint && codePoint <= 0x1b001) ||
    (0x1f200 <= codePoint && codePoint <= 0x1f202) ||
    (0x1f210 <= codePoint && codePoint <= 0x1f23a) ||
    (0x1f240 <= codePoint && codePoint <= 0x1f248) ||
    (0x1f250 <= codePoint && codePoint <= 0x1f251) ||
    (0x20000 <= codePoint && codePoint <= 0x2f73f) ||
    (0x2b740 <= codePoint && codePoint <= 0x2fffd) ||
    (0x30000 <= codePoint && codePoint <= 0x3fffd)
  ) {
    return "W";
  }
  if (
    (0x0020 <= codePoint && codePoint <= 0x007e) ||
    (0x00a2 <= codePoint && codePoint <= 0x00a3) ||
    (0x00a5 <= codePoint && codePoint <= 0x00a6) ||
    0x00ac == codePoint ||
    0x00af == codePoint ||
    (0x27e6 <= codePoint && codePoint <= 0x27ed) ||
    (0x2985 <= codePoint && codePoint <= 0x2986)
  ) {
    return "Na";
  }
  if (
    0x00a1 == codePoint ||
    0x00a4 == codePoint ||
    (0x00a7 <= codePoint && codePoint <= 0x00a8) ||
    0x00aa == codePoint ||
    (0x00ad <= codePoint && codePoint <= 0x00ae) ||
    (0x00b0 <= codePoint && codePoint <= 0x00b4) ||
    (0x00b6 <= codePoint && codePoint <= 0x00ba) ||
    (0x00bc <= codePoint && codePoint <= 0x00bf) ||
    0x00c6 == codePoint ||
    0x00d0 == codePoint ||
    (0x00d7 <= codePoint && codePoint <= 0x00d8) ||
    (0x00de <= codePoint && codePoint <= 0x00e1) ||
    0x00e6 == codePoint ||
    (0x00e8 <= codePoint && codePoint <= 0x00ea) ||
    (0x00ec <= codePoint && codePoint <= 0x00ed) ||
    0x00f0 == codePoint ||
    (0x00f2 <= codePoint && codePoint <= 0x00f3) ||
    (0x00f7 <= codePoint && codePoint <= 0x00fa) ||
    0x00fc == codePoint ||
    0x00fe == codePoint ||
    0x0101 == codePoint ||
    0x0111 == codePoint ||
    0x0113 == codePoint ||
    0x011b == codePoint ||
    (0x0126 <= codePoint && codePoint <= 0x0127) ||
    0x012b == codePoint ||
    (0x0131 <= codePoint && codePoint <= 0x0133) ||
    0x0138 == codePoint ||
    (0x013f <= codePoint && codePoint <= 0x0142) ||
    0x0144 == codePoint ||
    (0x0148 <= codePoint && codePoint <= 0x014b) ||
    0x014d == codePoint ||
    (0x0152 <= codePoint && codePoint <= 0x0153) ||
    (0x0166 <= codePoint && codePoint <= 0x0167) ||
    0x016b == codePoint ||
    0x01ce == codePoint ||
    0x01d0 == codePoint ||
    0x01d2 == codePoint ||
    0x01d4 == codePoint ||
    0x01d6 == codePoint ||
    0x01d8 == codePoint ||
    0x01da == codePoint ||
    0x01dc == codePoint ||
    0x0251 == codePoint ||
    0x0261 == codePoint ||
    0x02c4 == codePoint ||
    0x02c7 == codePoint ||
    (0x02c9 <= codePoint && codePoint <= 0x02cb) ||
    0x02cd == codePoint ||
    0x02d0 == codePoint ||
    (0x02d8 <= codePoint && codePoint <= 0x02db) ||
    0x02dd == codePoint ||
    0x02df == codePoint ||
    (0x0300 <= codePoint && codePoint <= 0x036f) ||
    (0x0391 <= codePoint && codePoint <= 0x03a1) ||
    (0x03a3 <= codePoint && codePoint <= 0x03a9) ||
    (0x03b1 <= codePoint && codePoint <= 0x03c1) ||
    (0x03c3 <= codePoint && codePoint <= 0x03c9) ||
    0x0401 == codePoint ||
    (0x0410 <= codePoint && codePoint <= 0x044f) ||
    0x0451 == codePoint ||
    0x2010 == codePoint ||
    (0x2013 <= codePoint && codePoint <= 0x2016) ||
    (0x2018 <= codePoint && codePoint <= 0x2019) ||
    (0x201c <= codePoint && codePoint <= 0x201d) ||
    (0x2020 <= codePoint && codePoint <= 0x2022) ||
    (0x2024 <= codePoint && codePoint <= 0x2027) ||
    0x2030 == codePoint ||
    (0x2032 <= codePoint && codePoint <= 0x2033) ||
    0x2035 == codePoint ||
    0x203b == codePoint ||
    0x203e == codePoint ||
    0x2074 == codePoint ||
    0x207f == codePoint ||
    (0x2081 <= codePoint && codePoint <= 0x2084) ||
    0x20ac == codePoint ||
    0x2103 == codePoint ||
    0x2105 == codePoint ||
    0x2109 == codePoint ||
    0x2113 == codePoint ||
    0x2116 == codePoint ||
    (0x2121 <= codePoint && codePoint <= 0x2122) ||
    0x2126 == codePoint ||
    0x212b == codePoint ||
    (0x2153 <= codePoint && codePoint <= 0x2154) ||
    (0x215b <= codePoint && codePoint <= 0x215e) ||
    (0x2160 <= codePoint && codePoint <= 0x216b) ||
    (0x2170 <= codePoint && codePoint <= 0x2179) ||
    0x2189 == codePoint ||
    (0x2190 <= codePoint && codePoint <= 0x2199) ||
    (0x21b8 <= codePoint && codePoint <= 0x21b9) ||
    0x21d2 == codePoint ||
    0x21d4 == codePoint ||
    0x21e7 == codePoint ||
    0x2200 == codePoint ||
    (0x2202 <= codePoint && codePoint <= 0x2203) ||
    (0x2207 <= codePoint && codePoint <= 0x2208) ||
    0x220b == codePoint ||
    0x220f == codePoint ||
    0x2211 == codePoint ||
    0x2215 == codePoint ||
    0x221a == codePoint ||
    (0x221d <= codePoint && codePoint <= 0x2220) ||
    0x2223 == codePoint ||
    0x2225 == codePoint ||
    (0x2227 <= codePoint && codePoint <= 0x222c) ||
    0x222e == codePoint ||
    (0x2234 <= codePoint && codePoint <= 0x2237) ||
    (0x223c <= codePoint && codePoint <= 0x223d) ||
    0x2248 == codePoint ||
    0x224c == codePoint ||
    0x2252 == codePoint ||
    (0x2260 <= codePoint && codePoint <= 0x2261) ||
    (0x2264 <= codePoint && codePoint <= 0x2267) ||
    (0x226a <= codePoint && codePoint <= 0x226b) ||
    (0x226e <= codePoint && codePoint <= 0x226f) ||
    (0x2282 <= codePoint && codePoint <= 0x2283) ||
    (0x2286 <= codePoint && codePoint <= 0x2287) ||
    0x2295 == codePoint ||
    0x2299 == codePoint ||
    0x22a5 == codePoint ||
    0x22bf == codePoint ||
    0x2312 == codePoint ||
    (0x2460 <= codePoint && codePoint <= 0x24e9) ||
    (0x24eb <= codePoint && codePoint <= 0x254b) ||
    (0x2550 <= codePoint && codePoint <= 0x2573) ||
    (0x2580 <= codePoint && codePoint <= 0x258f) ||
    (0x2592 <= codePoint && codePoint <= 0x2595) ||
    (0x25a0 <= codePoint && codePoint <= 0x25a1) ||
    (0x25a3 <= codePoint && codePoint <= 0x25a9) ||
    (0x25b2 <= codePoint && codePoint <= 0x25b3) ||
    (0x25b6 <= codePoint && codePoint <= 0x25b7) ||
    (0x25bc <= codePoint && codePoint <= 0x25bd) ||
    (0x25c0 <= codePoint && codePoint <= 0x25c1) ||
    (0x25c6 <= codePoint && codePoint <= 0x25c8) ||
    0x25cb == codePoint ||
    (0x25ce <= codePoint && codePoint <= 0x25d1) ||
    (0x25e2 <= codePoint && codePoint <= 0x25e5) ||
    0x25ef == codePoint ||
    (0x2605 <= codePoint && codePoint <= 0x2606) ||
    0x2609 == codePoint ||
    (0x260e <= codePoint && codePoint <= 0x260f) ||
    (0x2614 <= codePoint && codePoint <= 0x2615) ||
    0x261c == codePoint ||
    0x261e == codePoint ||
    0x2640 == codePoint ||
    0x2642 == codePoint ||
    (0x2660 <= codePoint && codePoint <= 0x2661) ||
    (0x2663 <= codePoint && codePoint <= 0x2665) ||
    (0x2667 <= codePoint && codePoint <= 0x266a) ||
    (0x266c <= codePoint && codePoint <= 0x266d) ||
    0x266f == codePoint ||
    (0x269e <= codePoint && codePoint <= 0x269f) ||
    (0x26be <= codePoint && codePoint <= 0x26bf) ||
    (0x26c4 <= codePoint && codePoint <= 0x26cd) ||
    (0x26cf <= codePoint && codePoint <= 0x26e1) ||
    0x26e3 == codePoint ||
    (0x26e8 <= codePoint && codePoint <= 0x26ff) ||
    0x273d == codePoint ||
    0x2757 == codePoint ||
    (0x2776 <= codePoint && codePoint <= 0x277f) ||
    (0x2b55 <= codePoint && codePoint <= 0x2b59) ||
    (0x3248 <= codePoint && codePoint <= 0x324f) ||
    (0xe000 <= codePoint && codePoint <= 0xf8ff) ||
    (0xfe00 <= codePoint && codePoint <= 0xfe0f) ||
    0xfffd == codePoint ||
    (0x1f100 <= codePoint && codePoint <= 0x1f10a) ||
    (0x1f110 <= codePoint && codePoint <= 0x1f12d) ||
    (0x1f130 <= codePoint && codePoint <= 0x1f169) ||
    (0x1f170 <= codePoint && codePoint <= 0x1f19a) ||
    (0xe0100 <= codePoint && codePoint <= 0xe01ef) ||
    (0xf0000 <= codePoint && codePoint <= 0xffffd) ||
    (0x100000 <= codePoint && codePoint <= 0x10fffd)
  ) {
    return "A";
  }

  return "N";
};

export function widestLine(string: any) {
	let lineWidth = 0;

	for (const line of string.split('\n')) {
		lineWidth = Math.max(lineWidth, stringWidth(line));
	}

	return lineWidth;
}

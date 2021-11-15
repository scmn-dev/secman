// This code is forked from https://github.com/hackmdio/hackmd-cli/blob/develop/src/read-stdin-stream.ts

export const readPipe: () => Promise<string | undefined> = () => {
  return new Promise((resolve) => {
    const stdin = process.openStdin();
    stdin.setEncoding("utf-8");

    let data = "";
    stdin.on("data", (chunk) => {
      data += chunk;
    });

    stdin.on("end", () => {
      resolve(data);
    });

    if (stdin.isTTY) {
      resolve("");
    }
  });
};

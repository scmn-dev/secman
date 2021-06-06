const { expect, test } = require("@oclif/test");

describe("fetch", () => {
  test
    .stdout()
    .command(["fetch"])
    .it("fetch if there's a new release", (ctx) => {
      expect(ctx.stdout).to.contain("");
    });
});

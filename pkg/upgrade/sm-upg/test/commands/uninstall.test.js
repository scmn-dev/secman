const { expect, test } = require("@oclif/test");

describe("fetch", () => {
  test
    .stdout()
    .command(["fetch"])
    .it("uninstall secman", (ctx) => {
      expect(ctx.stdout).to.contain("");
    });
});

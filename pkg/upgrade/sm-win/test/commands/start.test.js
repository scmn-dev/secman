const { expect, test } = require("@oclif/test");

describe("start", () => {
  test
    .stdout()
    .command(["start"])
    .it("start upgrading secman", (ctx) => {
      expect(ctx.stdout).to.contain("");
    });
});

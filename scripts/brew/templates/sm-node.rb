class SmNode < Formula
  desc "node.js dependency for secman"
  homepage "https://secman.dev/docs/cli"
  url "__NODE_BIN_URL__"
  version "__NODE_VERSION__"
  sha256 "__NODE_SHA256__"
  keg_only "sm-node"

  def install
    bin.install buildpath/"bin/node"
  end

  def test
    output = system bin/"node", "version"
    assert output.strip == "v#{version}"
  end
end

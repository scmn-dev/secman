class Secman < Formula
  desc "👊 Human-friendly and amazing secrets manager"
  homepage "https://secman.dev/docs/cli"
  url "__CLI_DOWNLOAD_URL__"
  sha256 "__CLI_SHA256__"
  license "MIT"

  livecheck do
    url :homepage
    regex(/href=.*?aamath[._-]v?(\d+(?:\.\d+)+)\.t/i)
  end

  depends_on "scmn-dev/secman/sm-node"

  def install
    inreplace "bin/secman", /^CLIENT_HOME=/, "export SECMAN_OCLIF_CLIENT_HOME=#{lib/"client"}\nCLIENT_HOME="
    inreplace "bin/secman", "\"$DIR/node\"", Formula["sm-node"].opt_bin/"node"
    libexec.install Dir["*"]
    bin.install_symlink libexec/"bin/secman"
    ENV["SM_PROVIDER"] = "brew"
  end

  test do
    system bin/"secman", "-v"
  end
end

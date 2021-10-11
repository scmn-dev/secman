package authflow

const oauthSuccessPage = `
<!DOCTYPE html>
<meta charset="utf-8" />
<title>Success: Secman With GitHub</title>
<style type="text/css">
  body {
    color: #1b1f23;
    background: #f6f8fa;
    font-size: 14px;
    font-family: -apple-system, "Segoe UI", Helvetica, Arial, sans-serif;
    line-height: 1.5;
    max-width: 620px;
    margin: 28px auto;
    text-align: center;
  }

  h1 {
    font-size: 24px;
    margin-bottom: 0;
  }

  p {
    margin-top: 0;
  }

  .box {
    border: 1px solid #e1e4e8;
    background: white;
    padding: 24px;
    margin: 28px;
  }
</style>
<body>
  <img src="https://assets.secman.dev/logo.svg" />
  <div class="box">
    <h1>Successfully authenticated Secman With Github 🔗</h1>
    <p>You may now close this tab and return to the terminal.</p>
  </div>
</body>
`

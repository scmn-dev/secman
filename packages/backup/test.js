let commandExists = require("command-exists");
const exec = require('child_process').exec;

function execute(command) {
  exec(command, (err, stdout, stderr) => {
    process.stdout.write(stdout);
  })
}

execute('echo "Hello World"');

const SM_GH_UN = () => {
  execute(
    "git config user.name"
  );
};

SM_GH_UN();

commandExists("secman", function (err, commandExists) {
    if (commandExists) {
        execute("echo 'secman is exist'");
    }
});

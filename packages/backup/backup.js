let commandExists = require("command-exists");
const { exec } = require("child_process");

function executeCommand(command) {
  const exec = require('child_process').exec;

  exec(command, (err, stdout, stderr) => {
    process.stdout.write(stdout);
  })
}

const SM_GH_UN = () => {
  executeCommand(
    "git config user.name"
  );
};

const create = () => {
  executeCommand(
    `gh repo create ${SM_GH_UN}/.secman.bk -y --private`
  );
};

// dirs
let SECDIR = "~/.secman.bk";
let SECDIR_primary = "~/.secman";

const cd_SECDIR = () => {
  executeCommand(
    `cd ${SECDIR}`
  );
};

// pkgs
const ghraw_url = "https://raw.githubusercontent.com";

let install_cgit = () => {
    executeCommand(`/bin/bash -c \"$(curl -fsSL ${ghraw_url}/Dev-x-Team/corgit/main/setup)\"`)
}

let install_brew = () => {
    executeCommand(`/bin/bash -c \"$(curl -fsSL ${ghraw_url}/Homebrew/install/HEAD/install.sh)\"`)
}

let install_verx = () => {
    executeCommand(`/bin/bash -c \"$(curl -fsSL ${ghraw_url}/abdfnx/verx/HEAD/install.sh)\"`)
}

const brew_gh = () => {
  executeCommand(
    "brew install gh"
  );
};

function _help() {
    const u =
        "Flags:\n\t-h | --help: help about any command\nCommands:\n\tsy | sync: create private github repo and sync your passwords on it by git\n\tcn | clone: clone .secman repo\n\tph | push: push and commit a new secret\n\tpl | pull: pull secret/s";

    console.log(u);
}

const csi = () => {
    executeCommand("cgit secman-ibk");
}

const rdm = () => {
    executeCommand(`echo \"# My secman backup passwords - ${SM_GH_UN}\" >> ${SECDIR}/README.md`);
}

const mkdir = () => {
    executeCommand(`cp -r ~/.secman ~/.secman.bk`)
}

function repo_work () {
    mkdir();
    rdm();
    create();
    executeCommand(
        `cp -r ~/.secman ~/.secman.bk && gh repo create ${SM_GH_UN}/.secman.bk -y --private && cd ~/.secman.bk && cgit secman-ibk`
    );
}

repo_work();

// function repo() {
//     commandExists("gh", function (err, commandExists) {
//         if (commandExists) {
//             repo_work();
//         } else {
//             commandExists("brew", function (err, commandExists) {
//                 if (commandExists) {
//                     brew_gh();

//                     commandExists("gh", function (err, commandExists) {
//                         if (commandExists) {
//                             repo_work();
//                         }
//                     });
//                 } else {
//                     install_brew();

//                     commandExists("brew", function (err, commandExists) {
//                         if (commandExists) {
//                             brew_gh();

//                             commandExists("gh", function (err, commandExists) {
//                                 if (commandExists) {
//                                     repo_work();
//                                 }
//                             });
//                         }
//                     });
//                 }
//             });
//         }
//     });
// }

// const _ph = () => {
//     const push = () => {
//         executeCommand("cgit ph");
//     }

//     commandExists("cgit", function (err, commandExists) {
//         if (commandExists) {
//             cd_SECDIR();
//             push();
//         } else {
//             install_cgit();

//             if (commandExists) {
//                 cd_SECDIR();
//                 push();
//             }
//         }
//     });
// }

// const _pl = () => {
//     const pull = () => {
//         executeCommand("cgit pl");
//     }

//     commandExists("cgit", function (err, commandExists) {
//         if (commandExists) {
//             cd_SECDIR();
//             pull();
//         } else {
//             install_cgit();

//             if (commandExists) {
//                 cd_SECDIR();
//                 pull();
//             }
//         }
//     });
// }

// function _clone() {
//     const clone = () => {
//         `gh repo clone ${SM_GH_UN}/.secman.bk ${SECDIR}`
//     }

//     clone();
// }

// const version = () => {
//     let v = () => {
//         executeCommand("secman ver");
//     }

//     commandExists("verx", function (err, commandExists) {
//         if (commandExists) {
//             v();
//         } else {
//             install_verx();

//             if (commandExists) {
//                 v();
//             }
//         }
//     });
// }

// var menuHandler;

// // Initialize
// function initialize() {
//     showMenu();

//     process.stdin.setEncoding('utf8');
//     process.stdin.on('readable', checkMenu);

//     function checkMenu() {
//         var input = process.stdin.read();
//         if(input !== null) {
//             menuHandler(input.trim());
//         }
//     }
// }

// // Main
// function showMenu() {
//     console.log(
//         'Init backup (init)' + '\n' +
//         'Clone backup (clone)'  + '\n' +
//         'Push the new passwords in backup (push)'  + '\n' +
//         'Pull (pull)'  + '\n\n' +
//         'Choose operation, then press ENTER: '
//     );

//     menuHandler = function(input) {
//         switch(input) {
//             case "init":
//                 repo();
//                 break;
            
//             case "clone":
//                 _clone();
//                 break;
            
//             case "push":
//                 _ph();
//                 break;
            
//             case "pull":
//                 _pl();
//                 break;
            
//             default:
//                 _help();
//                 break;
//         }
//     };
// }

// initialize();

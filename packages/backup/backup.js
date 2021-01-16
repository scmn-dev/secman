function exec(cmd, handler =
    function(err, stdout, stderr) {
        console.log(stdout);

        if(err !== null) {
            console.log(stderr);
        }
    }) {
    
    const childfork = require("child_process");
    return childfork.exec(cmd, handler);
}

var SM_GH_UN = exec("git config user.name");
const create = `gh repo create ${SM_GH_UN}/.secman.bk -y --private`;
let SECDIR = "~/.secman.bk";
let cd_SECDIR = `cd ${SECDIR}`;

// pkgs
const ghraw_url = "https://raw.githubusercontent.com";
let install_cgit = `/bin/bash -c \"$(curl -fsSL ${ghraw_url}/Dev-x-Team/corgit/main/setup)\"`;
let install_brew = `/bin/bash -c \"$(curl -fsSL ${ghraw_url}/Homebrew/install/HEAD/install.sh)\"`;
let install_verx = `/bin/bash -c \"$(curl -fsSL ${ghraw_url}/abdfnx/verx/HEAD/install.sh)\"`;

function usage() {
    const u=
        "Flags:\n\t-h | --help: help about any command\nCommands:\n\tsy | sync: create private github repo and sync your passwords on it by git\n\tcn | clone: clone .secman repo\n\tph | push: push and commit a new secret\n\tpl | pull: pull secret/s";

    console.log(u);
}

function repo_work() {
    const csi = "cgit secman-i";
    const rdm = `echo \"# My secman backup passwords - ${SM_GH_UN}\" >> $SECDIR/README.md`;

    exec(cd_SECDIR);
    exec(rdm);
    exec(create);
    exec(csi);
}
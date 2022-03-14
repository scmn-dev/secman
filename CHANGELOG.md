# Secman CHANGELOG

> All notable changes to secman will be documented in **CHANGELOG.md**

---

## [[v6.2.0] 2022-03-11](#v620-2022-03-11)

> Secman v3 ✨

this release is the first release of secman v3

it's comes with **new features**, **new design**, **new infrastructure**, **and new secman**.

### Added

- Add `secman auth` command with `create`, `login`, `logout`, and `refresh` sub-commands.
- Add `secman delete` command.
- Create `secman docs` and `secman doctor` commands.
- Build `secman edit` command.
- Create `secman generate` command.
- Add `secman info` and `secman init` commands.
- Build `secman insert` command.
- Create `secman read` command.
- Build `secman whoami` command.
- Initialize **@go-task** config file.
- Add Terminal User Interface (TUI) to secman with **@charmbracelet** libraries.
- Create `secman doctor` command with `secman doctor fix` sub-command.
- Add `secman ui` command.
- Build [**SMUI**](https://github.com/david-tomson/smui) by **@david-tomson**.

### Changed

- Move **secman v2** to secman core cli (**scc**).
- Convert from typescript to golang with typescript.
- Update `README.md`.
- Update **vscode** config.
- Update secman contributing guide.

### Fixed

- Fix messages.
- Fix connections with **secman cloud**.
- Fix Secman Core CLI (scc)

## [[v6.1.3] 2022-01-10](#v613-2022-01-10)

### Added

- Add `settings_editor_theme` prop in `~/.secman/settings.json`

### Changed

- Update pathes
- Update `Secman V1` url to `https://github.com/scmn-dev/secman/tree/v1`

### Fixed

- Fix `secman settings` command

## [[v6.1.2] 2022-01-01](#v612-2022-01-01)

### Added

- add `just-hash` flag to `crypto` command

### Changed

- upgrade `update` command
- remove `base64` hash from `crypto` command

### Fixed

- fix secman core version in `info` command

## [[v6.1.1] 2021-12-27](#v611-2021-12-27)

### Changed

- remove `bcrypt` command

### Fixed

- fix configs in windows

## [[v6.1.0] 2021-12-26](#v610-2021-12-26)

### Changed

- fix all secman help errors and bugs

## [[v6.0.9] 2021-12-26](#v609-2021-12-26)

> a release to fix a bug in the `secman`

## [[v6.0.8] 2021-12-14](#v608-2021-12-14)

### Added

- configure `.devcontainer` for secman

### Changed

nothing

### Fixed

- fix config files warning messages ([#120](https://github.com/scmn-dev/secman/issues/120))
- fix Homedir bug in windows ([#121](https://github.com/scmn-dev/secman/issues/121))

## [[v6.0.71] 2021-11-14](#v6071-2021-11-14)

### Added

nothing

### Changed

nothing

### Fixed

- Fix secman help output.

## [[v6.0.7] 2021-11-13](#v607-2021-11-13)

### Added

- Add `--user` and `--password` flags to `auth` command.
- Create `--password-stdin` flag to read password from stdin to `auth` command.
- Add errors catchers to show error messages.

### Changed

nothing

### Fixed

- Fix exit from secman editor in visual studio code terminal with <kbd>Alt + E</kbd> shortcut.
- Fix missing files error messages.

## [[v6.0.61] 2021-11-10](#v6061-2021-11-10)

### Added

nothing

### Changed

nothing

### Fixed

- Fix version checker in `update` command.

## [[v6.0.6] 2021-11-10](#v606-2021-11-10)

### Added

- Add `modify` & `change` aliases to `secman edit` command.
- Build github api to `api` dir.

### Changed

- Build our own `update` command.

### Fixed

- Fix `update` command.

## [[v6.0.5] 2021-11-08](#v605-2021-11-08)

### Added

nothing

### Changed

nothing

### Fixed

- Fix authentication check for `secman insert` command.

## [[v6.0.4] 2021-11-05](#v604-2021-11-05)

### Added

nothing

### Changed

nothing

## Fixed

- Fix version command.

## [[v6.0.3] 2021-11-05](#v603-2021-11-05)

> Secman V2

this release is the first release of secman v2

it's comes with new features, new design, new infrastructure, and new secman

### Added

- Create `crypto` command.
- Create `bcrypt` command.
- Add `docs` command.
- Build `info` command.
- Add `list` command.
- Create `settings` command.
- Create `update` command.
- `logout` command.
- Build `whoami` command.

### Changed

- All secman.
- Change secman language from `golang` to `typescript`.
- The root command is `secman .` instead of `secman`. if you execute `secman`, it will show the help.
- `auth`, `delete`, `edit`, `generate`, `init`, `insert`, read` commands are totally changed.
- New help design.
- Secman is now using the secman api. it's a new way to use secman. and now users can manage their secrets from everywhere.

### Fixed

- Fix many errors, bugs, and issues.

---

## [[v5.3.8] 2021-09-23](#v538-2021-09-23)

### Added

- Create `--topic` flag to `repo list` command.
- Add new secman docker image: [**smcr/secman-cli**](https://hub.docker.com/r/smcr/secman-cli).
- Add `brews` in `.goreleaser`.
- Add `write` & `new` aliases to `insert` command.
- Build `scmn-dev/browser` package.

### Changed

- Update the installers url
  * unix: [**unix.secman.dev**](https://unix.secman.dev)
  * windows: [**win.secman.dev**](https://win.secman.dev)
- Change the infrastructure of secman docker images
  * `smcr/secman`: secman container image for full experince.
  * `smcr/secman-cli`: lightweight image for demo.
- Move [`scmn-dev/gh-api`](https://github.com/scmn-dev/gh-api) & [`david-tomson/git`](https://github.com/david-tomson/git) repos to secman.
- upgrade secman docker CI.

### Fixed

- Fix `fetch` command in **windows** by adding `FetchClone` in `sync` command.

## [[v5.3.73] 2021-08-09](#v5373-2021-08-09)

### Added

- Add `repo sync` command.
- Add New Features to `gh-api`.

### Changed

- Change Secman Organization to `scmn-dev`.
- Bump Versions of our packages.

### Fixed

- Fix `sm-win` issues in windows.

## [[v5.3.72] 2021-08-04](#v5372-2021-08-04)

### Added

- Add `repo browse` command.

### Changed

- Move `browse` command to `repo browse`.

### Fixed

Nothing.

## [[v5.3.71] 2021-07-03](#v5371-2021-07-03)

### Added

- Create `browse` command, this command opens the repository in the browser.

### Changed

- Change `config` command to `gh-config`.
- Update `gh-api` pkg.
- Upgrade `clean` command, add two flags to it `--git/-g` & `--all/-a`, and add survey.
- Change scoop bucket url from **https://github.com/scmn-dev/sm-scoop** to **https://github.com/scmn-dev/secman** .
- Change **windows release message** to `sm-win start`.
- Update `upgrading` emoji 🚧.

### Fixed

nothing

## [[v5.3.7] 2021-06-20](#v537-2021-06-20)

### Added

- Add Spinner to `upgrade` command.

### Changed

- Change `sm-upg` to `sm-win`.
- Update Release Message.
- Change docs url to **https://docs.secman.dev** .

### Fixed

- Fix [**#57**](https://github.com/scmn-dev/secman/issues/57): **Fetching in windows**.
- Fix upgrading in windows.

## [[v5.3.63] 2021-06-04](#v5363-2021-06-04)

### Added

- Add `read` alias for `show` command.

### Changed

nothing

### Fixed

- Fix [**#52**](https://github.com/scmn-dev/secman/issues/52)

## [[v5.3.62] 2021-06-01](#v5362-2021-06-01)

### Added

nothing

### Changed

nothing

### Fixed

fix [**#49**](https://github.com/scmn-dev/secman/issues/49)

## [[v5.3.61] 2021-05-26](#v5361-2021-05-26)

### Added

- Create `-t/--use-template` flag.
- Add `secman auth get-username` command.
- Add user checker [**#42**](https://github.com/scmn-dev/secman/issues/42).

### Changed

- Change Secman Help Form with
  * **USAGE**
  * **COMMANDS**
  * **FLAGS**
  * **EXAMPLES**
  * **LEARN MORE**
  * **FEEDBACK**
- Change CLI Website from **get.secman.dev** to [**cli.secman.dev**](https://cli.secman.dev).

### Fixed

- Fix [**#43**](https://github.com/scmn-dev/secman/issues/43)
- Fix [**#42**](https://github.com/scmn-dev/secman/issues/42)

## [[v5.3.6] 2021-05-18](#v536-2021-05-18)

### Added

nothing

### Changed

- Complete `Uninstall` command with
  * `-d` & `--delete-data` flags (Just MacOS & Linux)

### Fixed

- Fix `Upgrade` command

## [[v5.3.5] 2021-05-10](#v535-2021-05-10)

### Added

- Add `Open` command, creator: @abdfnx
- Add `Sync`, creator: @abdfnx
  * Start
  * Clone
  * Push
  * Pull
- Add new websites
  * https://secman.dev
  * https://get.secman.dev
  * https://assets.secman.dev
  * https://changelog.secman.dev

### Changed

- ### **Make Secman is self-reliant** [#32](https://github.com/scmn-dev/secman/issues/32) (the most important change), The author of the idea: @abdfnx
- Change Secman Repo Form
- Change Git Config
- Move Our website from **next.js** to [**docusaurus**](https://docusaurus.io) [secman.dev#45](https://github.com/scmn-dev/secman.dev/issues/45), The author of the idea: @Timothee

### Fixed

- Fix version checks in windows [#29](https://github.com/scmn-dev/secman/issues/29)

## [[v5.3.4] 2021-04-19](#v534-2021-04-19)

### Added

- Add Secman Docker Image, creator: @abdfnx
  * in [Docker Hub](https://hub.docker.com/r/smcr/secman)
  * in [Github Packages](https://github.com/orgs/scmn-dev/packages/container/package/secman)
- Add Github Commands, creator: @abdfnx
  * Auth
    - Login
    - Logout
    - Refresh
    - Status
  * Repo
    - Clone
    - Create
    - Fork
    - List
- Configure Gitpod for secman, creator: @abdfnx
- Docker CI, creator: @abdfnx
- Add Contributing Guids at [**./.github/CONTRIBUTING.md**](./.github/CONTRIBUTING.md), creator: @abdfnx
- Create Github API with secman, creator: @abdfnx
- Create [sm-upg](https://www.npmjs.com/package/@secman/sm-upg) package for windows

### Changed

- Remove Macos Job at Secman CI
- Improve **Makefile** and make it more practical

### Fixed

- Fix **error: exit status 1 when type `secman -h`**

## [[v5.3.3] 2021-04-12](#v533-2021-04-12)

### Added

- Create imgs website **https://imgs-secman.web.app**, creator: @abdfnx
- Add [version-checker](https://github.com/scmn-dev/version-checker) package, creator: @abdfnx

### Changed

- Make version is main, ver is an alias
- Make clone is main, cn and / are aliases
- Change clone message, now it's shows the private repo at **https://github.com/:USERNAME/.secman**
- Add usage message in secman-sync
- Also Add usage message in cgit
- Clean up _cgit_ and _verx_
- Move sm folder location from **/home** to **~**
- Remove secman_windows_LATEST_VERSION_x64.msi for security reasons.

### Fixed

- Fix upgrade command

## [[v5.3.2] 2021-04-03](#v532-2021-04-03)

### Added

- Create / command (cn), creator @iMRxM7mD
- Configure upgrade command for windows, creator @abdfnx
- Create deps website **https://scmn-dev.github.io**, creator: @Timothee
- Create uninstall command, creator: @abdfnx
- When install secman in windows, now it's creates env path variable, creator: @abdfnx
- Add special build for scoop, creator: @abdfnx

### Changed

- Make upgrade is main, upg is an alias
- Now secman-un is not a subprogram, it's in sm folder
- When there's a new release, the message was changed
- Improve CircelCI Actions in secman
- Change the url of install secman from script
- Now secman doesn't need bash in windows 👍

### Fixed

- Fix syncing in windows
- Fix installing for windows

## [[v5.3.1] 2021-03-26](#v531-2021-03-26)

### Added

- Show warn message when secman dependencies are not found

### Changed

- Remove Dockerfile

### Fixed

nothing

## [[v5.3.0] 2021-03-10](#v530-2021-03-10)

### Added

- Create `clean` command

- In _sm-win_ folder
  * Add [**vx.ps1**](https://github.com/scmn-dev/sm-win/blob/code/vx.ps1)
  * Add [**ct.ps1**](https://github.com/scmn-dev/sm-win/blob/code/ct.ps1)
  * Add [**secman-sync.ps1**](https://github.com/scmn-dev/sm-win/blob/code/secman-sync.ps1)

### Changed

- Secman Dockerfile
- New version of shell plugin **v3.0.0**
- Remove backup command

### Fixed

- Fix instllation errors in secman_latest_version_x64.rpm
- Fix update checker in windows

## [[v5.2.1] 2021-03-03](#v521-2021-03-03)

### Added

- Add **error** dir, it's contain falied commands
- Create secman schema by graphql
- Create secman formula for homebrew 🍺,
  <br>
  and homebrew-assets repo

### Changed

- Remove **vm** command, because it's not very important thing
- Location of **sm** was changed, from **~** to **/home/**

### Fixed

- Fix the bug in `secman-sync`, the bug is
  when you sync your **~/.secman**
  it's should create a private repo under the username, and push all files to branch
  but in fact, it's create a private repo without push files to branch
  so we fix it...

- Fix installation errors in secman_latest_version.deb
- Fix and Remove virues in **secman-sync.exe**, so **shell** repo was created
- Fix `Docker CI` in github actions
- Fix **fetch** command in windows

## [[v5.2.0] 2021-02-21](#v520-2021-02-21)

### Added

- Add **fetch** command, and it's checking if there's a new password/s in _~/.secman_

### Changed

nothing

### Fixed

nothing

## [[v5.1.21] 2021-02-20](#v5121-2021-02-20)

### Added

- Create **sm** repo and on it all secman deps files

### Changed

- The check update way now is local and faster

### Fixed

- Fix all installations errors
- Fix `secman-sync` URLs
- Fix Github CIs/CDs

## [[v5.1.2] 2021-01-26](#v512-2021-01-26)

### Added

- Update `vm` command _in docker section_

### Changed

- **upd** command now is **upg** and it's alias is `upgrade`

### Fixed

- Fix verx urls in upgrade command

## [[v5.1.1] 2021-01-23](#v511-2021-01-23)

### Added

- Add secman CI
- Update **upd** command

### Changed

- Change `secman-sync` 👇

```code
secman insert hi/code
Enter password for hi/code:

syncing...
[trunk 3b86a48] new secman password
 2 files changed, 40 insertions(+)
 create mode 100644 files/hi/code
```

> When you insert or remove a password, `secman-sync` automatically well sync **.secman**

- Update **backup** command
- **generate** command is changed, it's now **gen** command

### Fixed

- Fix url errors
- Fix CI actions

## [[v5.1.0] 2021-01-22](#v510-2021-01-22)

### Added

- Finish from `vm` command
- Finish from `backup` command
- Add [Dockerfile](https://github.com/scmn-dev/secman/blob/main/api/vm/Dockerfile) for `vm`
- Create _**docs**_ folder in secman repo
- Create **[scmn-dev](https://github.com/scmn-dev)** org
- Add **.secman.yml**
- Create builds files

### Changed

- MV installer files to another repo: [install repo](https://github.com/scmn-dev/install)
- Change secman logo [<img src=".github/assets/icon.svg" align="center" width="15">](#)
- Transfer secman from `abdfnx/secman` to `scmn-dev/secman`
- Update README.md

### Fixed

- Fix a lot of backup errors
- Fix urls errors in `secman` installers & tools, `cgit` & `verx`
- Fix all sync problems & errors

## [[v5.0.2] 2021-01-15](#v502-2021-01-15)

### Added

- Add **upd** command

### Changed

MV to 5v
Dockerfile
M **backup** command

### Fixed

- Fix Installation errors
  - Github url
- Fix **vm** command with `CID`

## [[v5.0.1] 2021-01-15](#v501-2021-01-15)

### Added

- Add version checker
- Add secman docker image to docker hub

### Changed

- Remove install command

### Fixed

- Fix secman version command

## [[v5.0.0] 2021-01-14](#v500-2021-01-14)

### Added

- Start create install command
- Start configure vm command
- Apply new changes in core folder
- Try to add secman docker image

### Changed

- Big Modify for Dockerfile
- Remove **signin** command

### Fixed

Nothing

## [[v4.0.0] 2021-01-12](#v400-2021-01-12)

### Added

- Start create version checker
- Try to end **signin** command

### Changed

- secman-sync in _beta_ mod

### Fixed

- Fix secman installers

## [[v3.0.1] 2021-01-12](#v301-2021-01-12)

Release for tests

## [[v3.0.0] 2021-01-10](#v300-2021-01-10)

### Added

- Add **signin** command
- Add `secman-sync`

### Changed

- Configure corgit for automate git work

### Fixed

Nothing

## [[v2.0.0] 2020-12-27](#v200-2020-12-27)

### Added

- Add secman in github repo

### Changed

Nothing

### Fixed

Nothing

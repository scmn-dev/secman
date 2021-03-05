# Secman CHANGELOG

All notable changes to secman will be documented in **CHANGELOG.md**

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
- Add [Dockerfile](https://github.com/secman-team/secman/blob/main/api/vm/Dockerfile) for `vm`
- Create _**docs**_ folder in secman repo
- Create **[secman-team](https://github.com/secman-team)** org
- Add **.secman.yml**
- Create builds files

### Changed

- MV installer files to another repo: [install repo](https://github.com/secman-team/install)
- Change secman logo [<img src=".github/assets/icon.svg" align="center" width="15">](#)
- Transfer secman from `abdfnx/secman` to `secman-team/secman`
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

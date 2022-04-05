# Contributing

We accept pull requests for bug fixes and features where we've discussed the approach in an issue and given the go-ahead for a community member to work on it.
We'd also love to hear about ideas for new features as issues or discussions.

> To Contribute do:

* Open an issue if you got a problem or an error.
* Open an issue to propose a significant change.
* Open a pull request to fix a bug.
* Open a pull request to fix documentation about any command.
* Open a pull request for any issue labelled [`help wanted`][hw] or [`good first issue`][gfi].

## Build Secman

## Prerequisites:

- `go` version >= `17`.
- `npm` version >= `8.0.0`.
- `yarn` is installed.
- [`task`](https://taskfile.dev) is installed.

## Clone secman repo

```bash
# GitHub CLI
$ gh repo clone scmn-dev/secman

# Git
$ git clone https://github.com/scmn-dev/secman
```

## Change directory to secman repo

```bash
$ cd secman
```

### Build secman

- run **task bfs** to build secman cli.

```bash
$ task bfs
```

### Check secman

```bash
$ secman version
```

## Create or submitting a pull request

1. Create a new branch: `git checkout -b my-new-branch-name`
2. Make sure your changes and new fixes are without errors and bugs.
3. Create pull request at https://github.com/scmn-dev/secman/pulls

[hw]: https://github.com/scmn-dev/secman/labels/help%20wanted
[gfi]: https://github.com/scmn-dev/secman/labels/good%20first%20issue
[code-of-conduct]: ./CODE-OF-CONDUCT.md

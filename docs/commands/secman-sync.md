# secman-sync

by `secman-sync`, you can sync your passwords

## auth

you should authenticate by [`gh cli`](https://cli.github.com) to use **sync** feature

```sh
gh auth login
```

## sync

```sh
secman-sync sync
```

if you sync your passwords for first time, `sync` command will create a private github repo and store the passwords on it

`secman-sync sy` is an alias of `secman-sync sync`

## clone

```sh
secman-sync clone
```

if you lose your passwords, or you use more than device, you can clone your private repo

`secman-sync cn` is an alias of `secman-sync clone`

## push

```sh
secman-sync push
```

if there's a new password/s, it's well push it to the repo, like git

`secman-sync ph` is an alias of `secman-sync push`

## pull

```sh
secman-sync pull
```

we know what `pull` do

alias: `secman-sync pl`

### getting help

```code
secman-sync --help | -h
```

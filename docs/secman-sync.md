# Secman-sync

> by `secman-sync`, you can sync your passwords

## Commands

- **auth**
- **clone**
- **pull**
- **push**
- **sync**

## How it's works ?

1. init **`.git`** dir in ~/.secman

    ```sh
    cd ~/.secman
    git init
    ```

2. create a private repo by [**gh cli**](https://cli.github.com)

    before create repo, `secman-sync` gets git user name

    ```sh
    SM_GH_UN=$(git config user.name)
    ```

    ```sh
    gh repo create $SM_GH_UN/.secman -y --private
    ```

3. push the passwords
    after create the private repo, secman-sync add passwords, and push it

    ```sh
    git add .
    git commit -m "new .secman repo"
    git branch -M trunk
    git remote add origin https://github.com/$SM_GH_UN/.secman
    git push -u origin trunk
    ```

## Getting Help

```code
secman-sync --help | -h
```
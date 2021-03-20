# Introduction

```sm
███████╗╗███████╗ ██████╗███╗   ███╗ █████╗ ███╗    ███╗
██╔════╝║██╔════╝██╔════╝████╗ ████║██╔══██╗█████╗  ███║
███████╗║█████╗  ██║     ██╔████╔██║███████║███║███╗███║
╚════██║║██╔══╝  ██║     ██║╚██╔╝██║██╔══██║███║╚═█████║
███████║║███████╗╚██████╗██║ ╚═╝ ██║██║  ██║███║  ╚═███║
╚══════╝╚═══════╝ ╚═════╝╚═╝     ╚═╝╚═╝  ╚═╝╚══╝    ╚══╝
```

secman is a password manager, can create, edit, generate, and sync passwords.

## Features

- Not GPG cored.
- It uses a master password to securely store your passwords.
- Supports encrypting arbitrary files.
- It syncs your passwords

## Philosophy

> secman aims to create save passwords, and save it

1. You create your password, ex. FIREBASE_TOKEN
2. Secman encrypt it
3. Save it in `~/.secman/sites.json`
4. If you syncing your .secman, **secman-sync** do this mission
5. Push the passwords
6. Check the Updates

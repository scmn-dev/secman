# Initialize the .secman directory, and generate your secret keys.

Usage:
  `secman init [flags]`

Flags:
  **-h**, **--help**   help for init

```sh
secman init
```

Init should only be run one time, before running any other command. It is used for generating your master public private keypair.

By default, secman will create your password vault in the `.secman` directory within your home directory. You can override this location using the `SECDIR` environment variable.

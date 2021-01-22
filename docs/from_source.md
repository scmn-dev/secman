# Installation from source

> `secman` requires [Go](https://golang.org) version 1.11+

If `go` is not installed, follow steps on the [Go website](https://golang.org/doc/install).

1. clone secman repo

    ```sh
    # gh cli
    gh repo clone secman-team/secman
    cd secman
    ```

2. Build and install it

    ```sh
    # by default, it's installs to '/usr/local'; maybe you'll need sudo
    make install

    # if you want to install it to a different location
    make install prefix=/path/to/secman
    ```

3. Run `secman ver` to check if it worked.

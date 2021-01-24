# Installation from source

> `secman` requires [Go](https://golang.org) version 1.11+

If `go` is not installed, follow steps on the [Go website](https://golang.org/doc/install).

1. clone secman repo

    ```sh
    # gh cli
    gh repo clone secman-team/secman

    # git
    git clone https://github.com/secman-team/secman

    # after clone
    cd secman
    ```

2. Build and install it

    ```sh
    # by default, it's installs to '/usr/local'; maybe you'll need sudo
    make
    ```

3. Run `secman ver` to check if it worked.

# Remove a site from the password vault by specifying the entire site-path.

Usage:
  `secman remove [flags]`

Aliases:
  **remove**, **rm**

Examples:
  `secman remove core/docker.com`

Flags:
  **-h**, **--help**   help for remove

```code
secman
├──bb
|  └──ff
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io

secman remove bb/ff

secman
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io
```

remove is used for removing sites from the password vault. `secman rm` is an alias of `secman remove`.

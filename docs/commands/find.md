# Prints all sites that contain the site-path. Used to print just
one group or all sites that contain a certain word in the group or name.

Usage:
  secman find [flags]

Aliases:
  find, ls

Examples:
secman find code.com

Flags:
  -h, --help   help for find

```code
secman find git
└──git
   └──github.com

secman ls dev
└──dev
   └──dev.to
```

`find` and `ls` can both be used to search for all sites that contain a particular substring. It's good for printing out groups of sites as well. `secman ls` is an alias of `secman find`.

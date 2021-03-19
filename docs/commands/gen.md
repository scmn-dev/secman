# Prints a randomly generated password. The length of this password defaults to (24). If a password length is specified as greater than 2048 then generate will fail.

Usage:
  `secman gen [flags]`

Aliases:
  **gen**, **generate**

Examples:
  `secman generate`

Flags:
  -h, --help   help for gen

```code
secman gen
%L4^!s,Rry!}s:U<QwliL{vQKow321-!tr}:232

secman gen 8
#%Xy1t7E
```

secman can also create randomly generated passwords. The default length of secman generated passwords is 24 characters. This length can be changed by passing an optional length to the generate subcommand.

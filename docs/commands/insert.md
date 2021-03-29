# Add a site to your password store. This site can optionally be a part of a group by prepending a group name and slash to the site name. Will prompt for confirmation when a site path is not unique.

Usage:
  `secman insert [flags]`

Examples:
  `secman insert core/docker.com`

Flags:
  **-h**, **--help**   help for insert

## Inserting a password

```code
secman insert accounts/ionic
Enter password for accounts/ionic: 
```

Inserting a password in to your vault is easy. If you wish to group multiple entries together, it can be accomplished by prepending a group name followed by a slash to the pass-name.

Here we are adding ionic to the password store within the accounts group.

## Inserting a file 📝

```sh
secman insert money/budget.csv budget.csv
```

Adding a file works almost the same as insert. Instead it has an extra argument. The file that you want to add to your vault is the final argument.

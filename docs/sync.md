# Sync command

`secman-sync` is a **subprogram**, it's helps you to sync your passwords

## Commands

| Command | Alias | Work |
| :-----: | :---: | :--: |
| sync    |  sy   | start sync your .secman dir |
| clone   |  cn   | clone your .secman from private gh repo |
| push    |  ph   | push the new password |
| pull    |  pl   | pull secret/s |

## Backup

> sometimes, you maybe forget your ~/.secman dir & you delete your USERNAME/.secman

after sync your .secman, you should type `secman-sync backup init`

### Commmands

|  Flag  | Alias | Work |
|  :--:  | :---: | :--: |
| --init |  -i   | init your .secman.bk dir |
| --copy |  -c   | clone your .secman.bk from private gh repo |
| --push |  -p   | push the new backup password |
| --pull |  -l   | pull |
| --Main |  -M   | make .secman.bk is the Main dir |

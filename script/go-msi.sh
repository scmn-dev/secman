#!/bin/bash

smv=$(verx secman-team/secman -l)

# thanks for @mh-cbon 🙏 and his greate repo https://github.com/mh-cbon/go-msi
go-msi make --msi secman_windows_${smv}.msi --version ${smv} -s ./hooks/templates --path ../wix.json

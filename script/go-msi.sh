#!/bin/bash

# thanks for @mh-cbon 🙏 and his greate repo https://github.com/mh-cbon/go-msi
go-msi make --msi secman_windows_latest_version.msi --version latest_version -s ./hooks/templates --path ../wix.json

#!/bin/bash

echo "Enter the new secman name: "
read smn

echo "Enter the version: "
read smv

# thanks for @mh-cbon 🙏 and his greate repo https://github.com/mh-cbon/go-msi
go-msi make --msi $smn.msi --version $smv -s ./hooks/templates --path ../wix.json

#!/bin/bash

### Set git users to the commit we are building from ###
git config user.name "$(git --no-pager log --format=format:'%an' -n 1)"
git config user.email "$(git --no-pager log --format=format:'%ae' -n 1)"

### core ###
cd ~/build
git add .
git commit -m "secman mac build"
git branch -M main
git remote add origin https://github.com/secman-team/macosx
git push -u origin main

#!/bin/bash

### Set git users to the commit we are building from ###
sh -c "git config --global user.name '${GITHUB_ACTOR}' \
    && git config --global user.email '${GITHUB_ACTOR}@users.noreply.github.com' \
    && cd ~/build \
    && git init \
    && git add . \
    && git commit -m \"secman mac build\" \
    && git branch -M main \
    && git remote add origin https://github.com/secman-team/macosx \
    && git push -u origin main"

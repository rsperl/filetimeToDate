#!/usr/bin/env bash

REPOSITORY="$(git remote -v | head -1 | awk '{print $2}')"
BUILD_TIME="$(date --iso-8601=seconds)"
COMMIT_SHA="$(git rev-parse HEAD 2>/dev/null && true)"

# or
# git rev-parse --short=8 HEAD 2>/dev/null && true
COMMIT_SHORT_SHA="${COMMIT_SHA:0:8}"

BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD && true)

COMMIT_TAG="$(git tag --points-at $COMMIT_SHA && true)"

BUILD_HOSTNAME=$(hostname)

BUILD_USERNAME=$USER

# Reference the build-time variables as <package>.<Varname>
go build -o dist/filetime2date \
  -ldflags="-X main.Version=$COMMIT_TAG \
    -X main.Repository=$REPOSITORY \
    -X main.BuildCommit=$COMMIT_SHA \
    -X main.BuildShortCommit=$COMMIT_SHORT_SHA \
    -X main.BranchName=$BRANCH_NAME \
    -X main.BuildTime=$BUILD_TIME \
    -X main.BuildHostname=$BUILD_HOSTNAME \
    -X main.BuildUsername=$BUILD_USERNAME" \
  .

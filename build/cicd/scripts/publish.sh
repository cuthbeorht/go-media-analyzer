#!/bin/sh

CURRENT_VERSION=`git describe --tags`
NEW_VERSION=`cat VERSION`

PROTECTED_BRANCH="main"
CURRENT_BRANCH=`git rev-parse --abbrev-ref HEAD`

if [ $CURRENT_BRANCH != $PROTECTED_BRANCH ]; then
    echo "Must be on protected branch to publish a new version"
    echo "Branch: ${CURRENT_BRANCH}"
    exit -1
fi

if [ $CURRENT_VERSION == $NEW_VERSION ]; then    
    echo "No new version to publish"
fi

git tag $NEW_VERSION
git push origin $NEW_VERSION
GOPROXY=proxy.golang.org go list -m github.com/cuthbeorht/go-media-analyzer@$NEW_VERSION

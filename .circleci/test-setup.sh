#!/bin/bash

# Set up test environment

set -ex

# Install kudzu CMS
go get -u github.com/kudzu-cms/kudzu/...

# test install
kudzu

# create a project and generate code
if [ $CIRCLE_BRANCH = "kudzu-dev" ]; then
        # ensure we have the latest from kudzu-dev branch
        cd /go/src/github.com/kudzu-cms/kudzu
        git checkout kudzu-dev
        git pull origin kudzu-dev

        # create new project using the kudzu-dev branch
        kudzu new --dev github.com/kudzu-cms/ci/test-project
else
        kudzu new github.com/kudzu-cms/ci/test-project
fi

cd /go/src/github.com/kudzu-cms/ci/test-project

kudzu gen content person name:string hashed_secret:string
kudzu gen content message from:@person,hashed_secret to:@person,hashed_secret

# build and run dev http/2 server with TLS
kudzu build


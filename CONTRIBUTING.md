## Contributing

1. Checkout branch kudzu-dev
2. Make code changes
3. Test changes to kudzu-dev branch
    - make a commit to kudzu-dev
    - to manually test, you will need to use a new copy (kudzu new path/to/code), but pass the --dev flag so that kudzu generates a new copy from the kudzu-dev branch, not master by default (i.e. `$kudzu new --dev /path/to/code`)
    - build and run with $ kudzu build and $ kudzu run
4. To add back to master:
    - first push to origin kudzu-dev
    - create a pull request
    - will then be merged into master

_A typical contribution workflow might look like:_
```bash
# clone the repository and checkout kudzu-dev
$ git clone https://github.com/kudzu-cms/kudzu path/to/local/kudzu # (or your fork)
$ git checkout kudzu-dev

# install kudzu with go get or from your own local path
$ go get github.com/kudzu-cms/kudzu/...
# or
$ cd /path/to/local/kudzu
$ go install ./...

# edit files, add features, etc
$ git add -A
$ git commit -m 'edited files, added features, etc'

# now you need to test the feature.. make a new kudzu project, but pass --dev flag
$ kudzu new --dev /path/to/new/project # will create $GOPATH/src/path/to/new/project

# build & run kudzu from the new project directory
$ cd /path/to/new/project
$ kudzu build && kudzu run

# push to your origin:kudzu-dev branch and create a PR at kudzu-cms/kudzu
$ git push origin kudzu-dev
# ... go to https://github.com/kudzu-cms/kudzu and create a PR
```

**Note:** if you intend to work on your own fork and contribute from it, you will
need to also pass `--fork=path/to/your/fork` (using OS-standard filepath structure),
where `path/to/your/fork` _must_ be within `$GOPATH/src`, and you are working from a branch
called `kudzu-dev`.

For example:
```bash
# ($GOPATH/src is implied in the fork path, do not add it yourself)
$ kudzu new --dev --fork=github.com/nilslice/kudzu /path/to/new/project
```

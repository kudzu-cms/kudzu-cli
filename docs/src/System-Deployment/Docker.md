## kudzu Docker build

kudzu is distributed as a [docker image](https://hub.docker.com/r/kudzu/kudzu/),
which aids in kudzu deployment. The Dockerfile in this directory is used by kudzu
to generate the docker image which contains the kudzu executable.

If you are deploying your own kudzu project, you can write a new Dockerfile that
is based from the `kudzu/kudzu` image of your choice. For example:
```docker
FROM kudzu/kudzu:latest

# your project set up ...
# ...
# ...
```

### The following are convenient commands during development of kudzu core:

#### Build the docker image. Run from the root of the project.
```bash
# from the root of kudzu:
docker build -t kudzu-dev
```

#### Start the image, share the local directory and pseudo terminal (tty) into for debugging:
```bash
docker run -v $(pwd):/go/src/github.com/kudzu-cms/kudzu -it kudzu-dev
pwd # will output the go src directory for kudzu
kudzu version # will output the kudzu version
# make an edit on your local and rebuild
go install ./...
```

Special thanks to [**@krismeister**](https://github.com/krismeister) for contributing this!
